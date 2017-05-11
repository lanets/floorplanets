// @flow
import React from 'react';
import { paper, Point } from 'paper';

import type { SeatsMap } from '../reducers/types';

import Seat from './Seat';


type Props = {
  seats: SeatsMap,
}

export default class Floorplan extends React.Component {

  props: Props;
  view: paper.view;

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
    this.view.onMouseDrag = (e) => this.translateCamera(e);

    // HARDCODED: center the view for our mocked seats
    this.view.center = new Point(1175, 371);

    // Render the seats based on the loaded data.
    this.seats = [];
    for (const id in this.props.seats) {
      const seatdata = this.props.seats[id];
      const seat = new Seat(seatdata.x, seatdata.y, seatdata.label);
      this.seats.push(seat);
    }

    console.log('Floorplan initialized.');
    this.update();
  }

  translateCamera(event: Object) {
    const delta = event.delta;

    // smoothen transition
    delta.length /= 2;

    // inverts the scroll from the drag direction
    this.view.center = this.view.center.add(new Point(-delta.x, -delta.y));

    event.preventDefault()
  }

  update() {
    // hide seats that are not in the view
    for (let i = 0; i < this.seats.length; i++)
      this.applyOcclusion(this.seats[i]);

    this.view.update();

    requestAnimationFrame(this.update);
  }

  applyOcclusion(seat: Seat) {
    // hide seats that are not in the bounds of the view.
    seat.visible = seat.position.isInside(this.view.bounds);
  }

  render() {
    return (
      <canvas ref="canvas" />
    );
  }
}
