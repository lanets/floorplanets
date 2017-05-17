// @flow
import React from 'react';
import styled from 'styled-components';

import type { Tooltip } from '../reducers/types';


const ZoomButtons = styled.div`
  position: absolute;
  display: flex;
  flex-direction: column;
  justify-content: center;
  top: 0;
  left: 0;
  padding: 8px;
`;

const ZoomButton = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: #ccc;
  border: 1px solid #000000;
  cursor: pointer;
  text-align: center;
  margin: 0 0 8px 0;
`;

const TooltipWrap = styled.div`
  position: fixed;
  display: flex;
  justify-content: center;
  left: ${props => props.x}px;
  top: ${props => props.y}px;
  background-color: #000000;
  color: #ffffff;
  padding: 8px;
  border-radius: 4px;
  opacity: 0.8;
`;


type Props = {
  tooltip: Tooltip,

  zoomIn: (val: number) => void,
  zoomOut: (val: number) => void,
};

type State = {
  mouseX: number,
  mouseY: number,
};

export default class FloorplanUI extends React.Component {

  props: Props;
  state: State;

  constructor(props: Props) {
    super(props);

    this.state = {
      mouseX: 0,
      mouseY: 0,
    };

    window.onmousemove = e => this.setState({ mouseX: e.x, mouseY: e.y });
  }

  renderTooltip() {
    return (
      <TooltipWrap x={this.state.mouseX} y={this.state.mouseY}>
        { this.props.tooltip.text }
      </TooltipWrap>
    );
  }

  render() {
    return (
      <div>
        <ZoomButtons>
          <ZoomButton onClick={() => this.props.zoomIn(0.5)}>+</ZoomButton>
          <ZoomButton onClick={() => this.props.zoomOut(0.5)}>-</ZoomButton>
        </ZoomButtons>
        { this.props.tooltip.display ? this.renderTooltip() : null }
      </div>
    );
  }
}
