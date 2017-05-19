// @flow
import type { Action } from '../actions/types';
import type { Tooltip } from './types';


export const initial: Tooltip = {
  display: false,
  text: '',
}

export default function reducer(state: Tooltip = initial, action: Action): Tooltip {
  switch(action.type) {

    case 'SHOW_TOOLTIP':
      return { ...state, display: true, text: action.text };

    case 'HIDE_TOOLTIP':
      return { ...state, display: false };

    default:
      return state;
  }
}
