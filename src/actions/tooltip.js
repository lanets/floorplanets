// @flow

import type { Action } from './types';


export function showTooltip(text: string): Action {
  return { type: 'SHOW_TOOLTIP', text };
}

export function hideTooltip(): Action {
  return { type: 'HIDE_TOOLTIP' };
}
