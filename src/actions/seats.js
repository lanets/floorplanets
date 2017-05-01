/* @flow */
import type { Action } from './types';
import type { SeatData } from '../types';


export function loadSeats(seats: { [id: string]: SeatData }): Action {
  return { type: 'LOAD_SEATS', seats };
}
