// @flow
import type { SeatData } from './types';
import type { SeatState } from './reducers/types';


/**
 * Transform a SeatState object into a SeatData.
 *
 * This is use to convert the internal structure of a seat into it's
 * public representation, shown to the client.
 */
export function toSeatData(seatState: SeatState) : SeatData {
  return {
    label: seatState.label,
  };
}
