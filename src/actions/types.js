/* @flow */

import type { SeatsMap } from '../reducers/types';

/**
 *  Defines all the possible actions in the application.
 */
export type Action =
   { type: 'LOAD_SEATS', seats: SeatsMap }
