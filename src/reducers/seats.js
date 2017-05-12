/* @flow */

import type { Action } from '../actions/types';
import type { Seats } from './types';

export const initial: Seats = {
  seats: {},
};

export default function reducer(state: Seats = initial, action: Action): Seats {
  switch(action.type) {

    case 'LOAD_SEATS':
      return { ...state, seats: action.seats };

    default:
      return state;
  }
};
