// @flow
import type { Action } from '../actions/types';
import type { Zoom } from './types';

export const initial: Zoom = {
  zoom: 1,
};

export default function reducer(state: Zoom = initial, action: Action): Zoom {
  switch(action.type) {

    case 'ZOOM_IN':
      return { ...state, zoom: state.zoom + action.value};

    case 'ZOOM_OUT':
      return { ...state, zoom: state.zoom - action.value};

    default:
      return state;
  }
}
