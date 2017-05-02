// @flow
import {Point, Path, Group, PointText} from 'paper';

import { FLAT_COLORS } from '../colors';


export default class Seat {

  item: Group;
  position: Point;
  shape: Path;
  text: PointText;

  constructor(x: number, y: number, label: string) {
    this.position = new Point(x, y);

    const radius = 40;
    const outCircle = Path.Circle({
      position: this.position,
      radius,
      fillColor: FLAT_COLORS.POMEGRANATE,
      strokeColor: 'black',
    });

    const inCircle = Path.Circle({
      position: this.position,
      radius: radius * 0.8,
      fillColor: FLAT_COLORS.ALIZARIN,
    });

    this.text = new PointText({
      position: new Point(this.position.x, this.position.y + 12 * 0.5),
      content: label,
      fillColor: 'black',
      fontFamily: 'Helvetica',
      fontSize: 12,
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

    this.item.onClick = () => console.log(`Seat ${label} clicked !`);
  }
}
