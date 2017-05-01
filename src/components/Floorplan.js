// @flow
import React from 'react';
import paper from 'paper';


export default class Floorplan extends React.Component {

  componentDidMount() {
    // 4:3 ratio
    this.refs.canvas.width = 960;
    this.refs.canvas.height = 540;

    paper.setup(this.refs.canvas);

    console.log('Floorplan initialized.');
  }

  render() {
    return (
      <canvas ref="canvas" />
    );
  }
}
