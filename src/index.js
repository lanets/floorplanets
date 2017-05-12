// @flow
import React from 'react';
import { render } from 'react-dom';
import { Provider } from 'react-redux';
import { createStore, applyMiddleware } from 'redux';
import logger from 'redux-logger';

import App from './App';
import reducers from './reducers';

import { lanETS2016 } from './__mock__/lanets2016';
import { loadSeats } from './actions/seats';


let store = createStore(
  reducers,
  // chrome redux dev tool binding
  // https://github.com/zalmoxisus/redux-devtools-extension#11-basic-store
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__(),
  applyMiddleware(logger),
);

// Simulating the loading of seats from an external
// source: (file, localstorage, server, etc.) .
store.dispatch(loadSeats(lanETS2016));

render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);
