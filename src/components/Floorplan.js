// @flow
import React from 'react';
import paper from 'paper';

import type { SeatsMap } from '../reducers/types';

import Seat from './Seat';


type Props = {
  seats: SeatsMap,

  onObjectSelected: (obj: Object) => void,
}

export default class Floorplan extends React.Component {

  props: Props;

  componentDidMount() {
    // 4:3 ratio
    this.refs.canvas.width = 960;
    this.refs.canvas.height = 540;

    paper.setup(this.refs.canvas);

    this.generateSeats();
  }

  generateSeats() {
    // Render the seats based on the loaded data.
    for (const id in this.props.seats) {
      const seatdata = this.props.seats[id];
      const seat = new Seat(seatdata.x, seatdata.y, seatdata.label);

      seat.onSelect = () => this.handleSelectSeat(id);
    }
  }

  handleSelectSeat(id: string) {
    const seatState = this.props.seats[id];
    //TODO: Transform seatState into a SeatData.
    this.props.onObjectSelected(seatState);
  }

  render() {
    return (
      <canvas ref="canvas" />
    );
  }
}
