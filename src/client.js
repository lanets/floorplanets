// @flow
import React from 'react';
import { render } from 'react-dom';
import { Provider } from 'react-redux';
import { createStore, applyMiddleware } from 'redux';
import logger from 'redux-logger';

import type { FloorplanConfig } from './types';

import reducers from './reducers';
import Floorplan from './containers/Floorplan';
import { seatsData } from './__mock__/seats';
import { loadSeats } from './actions/seats';


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
    store.dispatch(loadSeats(seatsData));

    // Inject floorplan in the div of the HTML host.
    render(
      <Provider store={store}>
          <Floorplan />
      </Provider>,
      document.getElementById(this.config.div)
    );
  }
}