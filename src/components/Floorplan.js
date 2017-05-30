// @flow
import React from 'react';
import { paper, Point, Layer } from 'paper';

import type { SeatsMap } from '../reducers/types';
import type { SeatData } from '../types';

import { toSeatData } from '../conversions';
import Seat from './Seat';
import { FLAT_COLORS } from '../colors';


type Props = {
  seats: SeatsMap,
  zoom: number,

  // JS api callback
  onSelectSeat: (seat: SeatData) => void,
  seatColor: (seat: SeatData) => ?string,
  seatTooltip: (seat: SeatData) => ?string,
  seatLabel: (seat: SeatData) => ?string,

  showTooltip: (text: string ) => void,
  hideTooltip: () => void,
}

export default class Floorplan extends React.Component {

  props: Props;
  view: paper.view;
  rasterLayer: Layer;
  seatsLayer: Layer;

  shouldRaster: bool;
  seats: Seat[];
  update: () => void;

  constructor(props: Props) {
    super(props);
    this.update = this.update.bind(this);
  }

  componentDidMount() {
    const canvas = this.refs.canvas;

    // Setup canvas
    canvas.width = 960;
    canvas.height = 540;

    // Configure the floorplan
    paper.setup(canvas);
    this.view = paper.view;
    this.view.autoUpdate = false;
    this.view.zoom = this.props.zoom;
    this.view.onMouseDrag = (e) => this.onMouseDrag(e);
    this.view.onMouseUp = (e) => this.onMouseUp(e);

    this.seatsLayer = paper.project.activeLayer;
    this.rasterLayer = paper.project.addLayer(new Layer());

    // HARDCODED: center the view for our mocked seats
    this.view.center = new Point(1175, 371);

    // Create and render the seats based on the loaded floorplan.
    this.seats = [];
    for (const id in this.props.seats) {
      const seatstate = this.props.seats[id];
      const seat = new Seat(id, seatstate.x, seatstate.y);

      // events binding
      seat.onMouseEnter = () => {
        const tooltipText = this.props.seatTooltip(toSeatData(seatstate));
        this.props.showTooltip(tooltipText || seatstate.label);
      };
      seat.onMouseLeave = () => this.props.hideTooltip();
      seat.onSelect = () => this.props.onSelectSeat(toSeatData(seatstate));

      this.seats.push(seat);
    }

    this.update();
  }

  /**
   *
   * Called when the stores the floorplan is subscribed to is updated.
   */
  componentDidUpdate() {
    this.view.zoom = this.props.zoom;
    this.update();
  }

  /**
   * Called when the user drags the floorplan around.
   */
  onMouseDrag(e: Object) {
    const delta = e.delta;

    // reduce drag length
    delta.length /= 2;

    // inverts the scroll from the drag direction
    this.view.center = this.view.center.add(new Point(-delta.x, -delta.y));

    // When dragging the floorplan, raster the viewport for insane performances.
    this.shouldRaster = true;

    this.update();
  }

  /**
   * Called when the user is done dragging the floorplan
   */
  onMouseUp(e: Object) {
    this.shouldRaster = false;
    this.update();
  }

  update() {
    if (this.shouldRaster) {
      if(this.rasterLayer.isEmpty())
        this.rasterLayer.addChild(this.seatsLayer.rasterize());

      // hide the seats layer and activate the raster layer.
      this.seatsLayer.visible = false;
      this.rasterLayer.visible = true;
    } else {
      // hide the raster layer and activate the seats layer.
      this.rasterLayer.visible = false;
      this.seatsLayer.visible = true;

      this.updateSeats();
    }

    // reset the raster flag and redraw the floorplan.
    this.shouldRaster = false;
    this.view.requestUpdate(this.view.update())
  }

  /**
   * Update each seats based on the callback properties of the configuration provided
   * by the user.
   */
  updateSeats() {
    this.seats.forEach((seat) => {
      const seatData = toSeatData(this.props.seats[seat.id]);

      seat.color = this.props.seatColor(seatData) || FLAT_COLORS.POMEGRANATE;
      seat.label = this.props.seatLabel(seatData) || '';
    });
  }

  render() {
    return <canvas ref="canvas" />;
  }
}
