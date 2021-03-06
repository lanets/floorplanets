// @flow
import React from 'react';
import styled from 'styled-components';
import { render } from 'react-dom';
import { Provider } from 'react-redux';
import { createStore, applyMiddleware } from 'redux';
import logger from 'redux-logger';

import type { FloorplanConfig } from './types';

import reducers from './reducers';
import Floorplan from './containers/Floorplan';
import FloorplanUI from './containers/FloorplanUI';
import { loadSeats } from './actions/seats';

import { lanETS2016 } from './__mock__/lanets2016';


const Wrapper = styled.div`
 position: relative;
 border: 1px solid black;
`;

export class FloorplanClient {

  config: FloorplanConfig;

  constructor(config: FloorplanConfig) {
    if (!config.div)
      throw new Error("'div' field missing in Floorplan configuration.");

    this.config = config;
  }

  init() {
    const store = createStore(
      reducers,
      // chrome redux dev tool binding
      // https://github.com/zalmoxisus/redux-devtools-extension#11-basic-store
      window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__(),
      applyMiddleware(logger),
    );

    // Simulating the loading of seats from an external
    // source: (file, localstorage, server, etc.) .
    store.dispatch(loadSeats(lanETS2016));

    // Inject floorplan in the div of the HTML host.
    render(
      <Provider store={store}>
        <Wrapper>
            <Floorplan
              onSelectSeat={this.config.onSelectSeat}
              seatColor={this.config.seatColor || (() => null)}
              seatTooltip={this.config.seatTooltip || (() => null)}
              seatText={this.config.seatText || (() => null)}
            />
            <FloorplanUI />
        </Wrapper>
      </Provider>,
      document.getElementById(this.config.div)
    );
  }
}
