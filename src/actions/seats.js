/* @flow */
import type { Action } from './types';
import type { SeatsMap } from '../reducers/types';


export function loadSeats(seats: SeatsMap ): Action {
  return { type: 'LOAD_SEATS', seats };
}
