/* @flow */
import type { Action } from './types';
import type { SeatData } from '../types';


export function loadSeats(seats: { [id: number]: SeatData }) {
  return { type: 'LOAD_SEATS', seats };
}
