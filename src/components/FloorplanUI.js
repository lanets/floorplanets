// @flow
import React from 'react';
import styled from 'styled-components';


const Wrapper = styled.div`
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

export default class FloorplanUI extends React.Component {

  render() {
    return (
      <Wrapper>
        <ZoomButton>+</ZoomButton>
        <ZoomButton>-</ZoomButton>
      </Wrapper>
    );
  }
}
