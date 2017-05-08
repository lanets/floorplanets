// @flow
import React from 'react';
import paper from 'paper';

import Seat from './Seat';


export default class Floorplan extends React.Component {

  componentDidMount() {
    // 4:3 ratio
    this.refs.canvas.width = 960;
    this.refs.canvas.height = 540;

    paper.setup(this.refs.canvas);

    // Simple, hardcoded, seat rendering.
    const seats: Seat[] = [];
    seats.push(new Seat(60, 60, 'G-13'));
    seats.push(new Seat(150, 60, 'G-14'));
    seats.push(new Seat(240, 60, 'G-15'));

    console.log('Floorplan initialized.');
  }

  render() {
    return (
      <canvas ref="canvas" />
    );
  }
}
