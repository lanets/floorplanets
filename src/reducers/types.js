/* @flow */


/*
 * The data binding between the state of a seat and the graphic
 * representation of a seat.
 */
export type SeatState = {
  id: number,
  x: number,
  y: number,
  label: string,  // what's written on the seat
};

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
