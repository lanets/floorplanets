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
    paper.setup(canvas);

    // view bindings
    this.view = paper.view;
    this.view.autoUpdate = false;

    this.view.zoom = this.props.zoom;

    this.view.onMouseDrag = (e) => this.translateCamera(e);

    // HARDCODED: center the view for our mocked seats
    this.view.center = new Point(1175, 371);

    // Create and render the seats based on the loaded data.
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

  componentDidUpdate() {
    this.view.zoom = this.props.zoom;
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

    event.preventDefault()
    this.update();
  }

  update() {
    requestAnimationFrame(() => {

      if (this.shouldRaster) {

        if(paper.project.layers.length < 2) {
          const rasterLayer = paper.project.addLayer(new Layer());
          // Freeze the canvas and render a raster on the raster layer
          rasterLayer.addChild(paper.project.layers[0].rasterize());

          // hide the floorplan layer
          paper.project.layers[0].visible = false;
        }

      } else {
        // Draw each seats based on the dynamic properties of the configuration provided
        // by the user.
        this.seats.forEach((seat) => {

          // generate the seatData used for callbacks
          const seatData = toSeatData(this.props.seats[seat.id]);

          // seat.visible = seat.position.isInside(this.view.bounds);
          seat.color = this.props.seatColor(seatData);

        });
      }


      // reset the raster flag
      this.shouldRaster = false;

      // redraw the whole floorplan
      this.view.update();
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
