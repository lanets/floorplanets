import React, { Component } from 'react';
import styled from 'styled-components';

import Floorplan from './components/Floorplan';

const Wrapper = styled.div`
  text-align: center;
`;

const Header = styled.div`
  background-color: #222;
  height: 150px;
  padding: 20px;
  color: white;
`;

class App extends Component {
  render() {
    return (
      <Wrapper>
        <Header>
          <h2>Floorplan</h2>
        </Header>
        <Floorplan />
      </Wrapper>
    );
  }
}

export default App;
