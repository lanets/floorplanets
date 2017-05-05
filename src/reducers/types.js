/* @flow */
import type { SeatState } from '../types';


/**
 *  Hashmap of seats.
 *  key: id of the seat
 *  value: SeatState object
 */
export type SeatsMap = {[id: string]: SeatState};

/**
 *  Structure of the Seats store.
 */
export type Seats = {
  seats: SeatsMap,
};
