// @flow
import type { Action } from './types';


export function zoomIn(value: number): Action {
  return { type: 'ZOOM_IN', value };
}

export function zoomOut(value: number): Action {
  return { type: 'ZOOM_OUT', value };
}
