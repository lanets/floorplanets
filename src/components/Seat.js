// @flow
import {Point, Shape, Group, PointText} from 'paper';

import { FLAT_COLORS } from '../colors';


export default class Seat {

  item: Group;
  position: Point;
  shape: Shape;
  text: PointText;

  constructor(x: number, y: number, label: string) {
    this.position = new Point(x, y);

    const radius = 9;
    const outCircle = Shape.Circle({
      position: this.position,
      radius,
      fillColor: FLAT_COLORS.POMEGRANATE,
      strokeColor: 'black',
    });

    const inCircle = Shape.Circle({
      position: this.position,
      radius: radius * 0.8,
      fillColor: FLAT_COLORS.ALIZARIN,
    });

    this.text = new PointText({
      position: new Point(this.position.x, this.position.y + 6 * 0.5),
      content: label,
      fillColor: 'black',
      fontFamily: 'Helvetica',
      fontSize: 6,
      justification: 'center',
    });


    this.item = new Group({
      position: this.position,
      children: [
        outCircle,
        inCircle,
        this.text,
      ],
    });

    this.item.onClick = () => console.log(this.item);
  }
}
