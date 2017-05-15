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

    case 'ZOOM_OUT': {
      const nextZoom = state.zoom - action.value > 0 ? state.zoom - action.value : state.zoom;
      return { ...state, zoom: nextZoom};
    }

    default:
      return state;
  }
}
