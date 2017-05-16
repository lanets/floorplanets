/* @flow */

import type { SeatsMap } from '../reducers/types';

/**
 *  Defines all the possible actions in the application.
 */
export type Action =
  { type: 'LOAD_SEATS', seats: SeatsMap }

  // Zoom actions
| { type: 'ZOOM_IN', value: number }
| { type: 'ZOOM_OUT', value: number }

  //  Tooltip actions
| { type: 'SHOW_TOOLTIP', text: string, x: number, y: number }
