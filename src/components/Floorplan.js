// @flow
import React from 'react';
import paper from 'paper';

import type { SeatsMap } from '../reducers/types';

import Seat from './Seat';

type Props = {
  seats: SeatsMap,
}

export default class Floorplan extends React.Component {

  props: Props;

  componentDidMount() {
    // 4:3 ratio
    this.refs.canvas.width = 960;
    this.refs.canvas.height = 540;

    paper.setup(this.refs.canvas);

    // Simple, hardcoded, seat rendering.
    const seats: Seat[] = [];

    // Render the seats based on the loaded data.
    for (const id in this.props.seats) {
      const seatdata = this.props.seats[id];
      const seat = new Seat(seatdata.x, seatdata.y, seatdata.label);
      seats.push(seat);
    }

    console.log('Floorplan initialized.');
  }

  render() {
    return (
      <canvas ref="canvas" />
    );
  }
}
