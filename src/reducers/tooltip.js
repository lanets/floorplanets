// @flow
import type { Action } from '../actions/types';
import type { Tooltip } from './types';


export const initial: Tooltip = {
  display: false,
}

export default function reducer(state: Tooltip, action: Action): Tooltip {
  switch(action.type) {

    case 'SHOW_TOOLTIP':
      return { ...state, display: true };

    default:
      return state;
  }
}
