// @flow
import React from 'react';
import { paper } from 'paper';

import type { SeatsMap } from '../reducers/types';
import type { SeatData } from '../types';

import Seat from './Seat';
import Viewport from '../viewport';
import { toSeatData } from '../conversions';
import { FLAT_COLORS } from '../colors';


type Props = {
  seats: SeatsMap,
  zoom: number,

  // JS api callback
  onSelectSeat: (seat: SeatData) => void,
  seatColor: (seat: SeatData) => ?string,
  seatTooltip: (seat: SeatData) => ?string,
  seatText: (seat: SeatData) => ?string,

  // Redux props
  showTooltip: (text: string ) => void,
  hideTooltip: () => void,
}

/**
 * Component rendering a HTML5 Canvas Element containing a Floorplan.
 */
export default class Floorplan extends React.Component {

  props: Props;
  viewport: Viewport;

  seats: Seat[];

  componentDidMount() {
    const canvas = this.refs.canvas;

    // Setup canvas
    canvas.width = 960;
    canvas.height = 540;

    // Configure the floorplan
    paper.setup(canvas);
    this.viewport = new Viewport(paper.view, paper.project);
    this.viewport.setZoom(this.props.zoom);

    // HARDCODED: center the view for our mocked seats
    this.viewport.setCenter(1175, 371);

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

  componentDidUpdate() {
    this.update();
  }

  /**
   * Update the floorplan's seat and redraw the viewport.
   */
  update() {
    this.viewport.setZoom(this.props.zoom);

    // Use the user defined callback to modify how each seat's should look like.
    this.seats.forEach((seat) => {
      const seatData = toSeatData(this.props.seats[seat.id]);
      seat.color = this.props.seatColor(seatData) || FLAT_COLORS.POMEGRANATE;
      seat.text = this.props.seatText(seatData) || '';
    });

    // finally, redraw the viewport
    this.viewport.redraw();
  }

  render() {
    return <canvas ref="canvas" />;
  }
}
