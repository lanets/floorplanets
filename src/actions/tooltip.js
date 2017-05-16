// @flow

import type { Action } from './types';


export function showTooltip(text: string): Action {
  return { type: 'SHOW_TOOLTIP', text };
}
