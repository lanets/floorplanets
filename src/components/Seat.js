// @flow
import { Point, Shape, Group, PointText } from 'paper';
import { FLAT_COLORS } from '../colors';


export default class Seat {

  id: string;
  item: Group;
  position: Point;
  shape: Shape;
  text: PointText;

  visible: boolean;
  color: string;

  onMouseEnter: () => void;
  onMouseLeave: () => void;
  onSelect: () => void;

  constructor(id: string, x: number, y: number) {

    this.id = id;
    this.position = new Point(x, y);

    this.item = Shape.Circle({
      position: this.position,
      radius: 8,
      strokeColor: 'black',
      fillColor: FLAT_COLORS.POMEGRANATE,
    });

    this.item.onMouseEnter = () => this.onMouseEnter();
    this.item.onMouseLeave = () => this.onMouseLeave();
    this.item.onClick = () => this.onSelect();
  }

  get color(): string {
    return this.item.fillColor;
  }

  set color(val: ?string) {
    this.item.fillColor = val || FLAT_COLORS.POMEGRANATE;
  }

  get visible(): boolean {
    return this.item.visible;
  }

  set visible(val: boolean) {
    this.item.visible = val;
  }
}
