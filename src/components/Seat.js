// @flow
import {Point, Shape, Group, PointText} from 'paper';

import { FLAT_COLORS } from '../colors';


export default class Seat {

  item: Group;
  position: Point;
  shape: Shape;
  text: PointText;

  visible: boolean;

  constructor(x: number, y: number, label: string) {
    this.position = new Point(x, y);

    const radius = 9;
    this.item = Shape.Circle({
      position: this.position,
      radius,
      fillColor: FLAT_COLORS.POMEGRANATE,
      strokeColor: 'black',
    });

    this.item.onClick = () => console.log(this.item);
  }

  get visible(): boolean {
    return this.item.visible;
  }

  set visible(val: boolean) {
    this.item.visible = val;
  }
}
