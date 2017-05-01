/* @flow */
import type { SeatData } from '../types';


/**
 *  Hashmap of seats.
 *  key: id of the seat
 *  value: SeatData object
 */
export type SeatsMap = {[id: number]: SeatData};

/**
 *  Structure of the Seats store.
 */
export type Seats = {
  seats: SeatsMap,
};
