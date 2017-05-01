/* @flow */

import type { SeatData } from '../types';

/**
 *  Defines all the possible actions in the application.
 */
export type Action =
   { type: 'LOAD_SEATS', seats: { [id: number]: SeatData } }
