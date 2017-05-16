// @flow
import type { Action } from '../actions/types';
import type { Tooltip } from './types';


export const initial: Tooltip = {
  display: false,
  x: 0,
  y: 0,
  text: '',
}

export default function reducer(state: Tooltip, action: Action): Tooltip {
  switch(action.type) {

    case 'SHOW_TOOLTIP':
      return {
        ...state,
        display: true,
        x: action.x,
        y: action.y,
        text: action.text,
      };

    default:
      return state;
  }
}
