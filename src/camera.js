// @flow
import { Point } from 'paper';

export function translationCenter(oldCenter: Point, delta: Point) {
  delta.length /= 2;
  return oldCenter.add(new Point(-delta.x, -delta.y));
}
