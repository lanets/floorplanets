/* @flow */
import type { SeatState } from '../types';


/**
 *  Hashmap of seats.
 *  key: id of the seat
 *  value: SeatState object
 */
export type SeatsMap = {[id: string]: SeatState};

/**
 *  Structure of the `seats` store.
 */
export type Seats = {
  seats: SeatsMap,
};

/**
 *  Structure of the `zoom` store.
 */
export type Zoom = {
  zoom: number,
};

/**
 *  Structure of the `tooltip` store.
 */
export type Tooltip = {
  display: boolean,
  x: number,
  y: number,
  text: string,
};
