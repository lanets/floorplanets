// @flow
import React from 'react';
import { paper, Point, Layer } from 'paper';

import type { SeatsMap } from '../reducers/types';
import type { SeatData } from '../types';

import { toSeatData } from '../conversions';
import Seat from './Seat';


type Props = {
  seats: SeatsMap,
  zoom: number,

  // JS api callback
  onSelectSeat: (seat: SeatData) => void,
  seatColor: (seat: SeatData) => string,

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
    this.view.onMouseDrag = (e) => this.translateCamera(e);
    this.view.onMouseUp = (e) => this.handleMouseUp(e);

    this.seatsLayer = paper.project.activeLayer;
    this.rasterLayer = paper.project.addLayer(new Layer());

    // HARDCODED: center the view for our mocked seats
    this.view.center = new Point(1175, 371);

    // Create and render the seats based on the loaded floorplan.
    this.seats = [];
    for (const id in this.props.seats) {
      const seatdata = this.props.seats[id];
      const seat = new Seat(id, seatdata.x, seatdata.y);

      // events binding
      seat.onMouseEnter = () => this.props.showTooltip(seatdata.label);
      seat.onMouseLeave = () => this.props.hideTooltip();
      seat.onSelect = () => this.handleSelectSeat(id);

      this.seats.push(seat);
    }

    this.update();
  }

  /**
   * Called when the stores the floorplan is subscribed to is updated.
   */
  componentDidUpdate() {
    this.view.zoom = this.props.zoom;
    this.update();
  }

  handleMouseUp(e: Object) {
    this.shouldRaster = false;
    this.update();
  }

  translateCamera(event: Object) {
    const delta = event.delta;

    // smoothen transition
    delta.length /= 2;

    // inverts the scroll from the drag direction
    this.view.center = this.view.center.add(new Point(-delta.x, -delta.y));

    // When translating, raster the viewport for insane performances.
    this.shouldRaster = true;

    this.update();
  }

  update() {
    if (this.shouldRaster) {
      // TODO: Change isEmpty call to a call that check if the representation of the seats has changed
      if(this.rasterLayer.isEmpty())
        this.rasterLayer.addChild(this.seatsLayer.rasterize());

      // hide the seats layer and activate the raster layer.
      this.seatsLayer.visible = false;
      this.rasterLayer.visible = true;
    } else {
      // hide the raster layer and activate the seats layer.
      this.rasterLayer.visible = false;
      this.seatsLayer.visible = true;

      this.drawSeats();
    }

    // reset the raster flag and redraw the floorplan.
    this.shouldRaster = false;
    this.view.requestUpdate(this.view.update())
  }

  /**
   * Draw each seats based on the dynamic properties of the configuration provided
   * by the user.
   */
  drawSeats() {
    this.seats.forEach((seat) => {
      // generate the seatData used for callbacks
      const seatData = toSeatData(this.props.seats[seat.id]);

      // seat.visible = seat.position.isInside(this.view.bounds);
      seat.color = this.props.seatColor(seatData);
    });
  }

  handleSelectSeat(id: string) {
    const seatState = this.props.seats[id];
    this.props.onSelectSeat(toSeatData(seatState));
  }

  render() {
    return (
      <canvas ref="canvas" />
    );
  }
}
