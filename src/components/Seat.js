// @flow
import { Point, Shape, Group, PointText } from 'paper';
import { FLAT_COLORS } from '../colors';


const FONTSIZE = 4;

export default class Seat {

  id: string;
  item: Group;
  position: Point;
  seatShape: Shape;
  seatText: PointText;

  visible: boolean;
  color: string;
  label: string;

  onMouseEnter: () => void;
  onMouseLeave: () => void;
  onSelect: () => void;

  constructor(id: string, x: number, y: number) {

    this.id = id;
    this.position = new Point(x, y);

    // the circle representing the seat.
    this.seatShape = Shape.Circle({
      position: this.position,
      radius: 8,
      strokeColor: 'black',
      fillColor: FLAT_COLORS.POMEGRANATE,
    });

    // the text label inside the seat.
    this.seatText = new PointText({
      point: new Point(x, y + FONTSIZE * 0.5),
      fontSize: FONTSIZE,
      justification: 'center',
      fillColor: 'black',
      content: '',
    });

    this.item = new Group([
      this.seatShape,
      this.seatText,
    ])

    this.item.onMouseEnter = () => this.onMouseEnter();
    this.item.onMouseLeave = () => this.onMouseLeave();
    this.item.onClick = () => this.onSelect();
  }

  get color(): string {
    return this.seatShape.fillColor;
  }

  set color(val: string) {
    this.seatShape.fillColor = val;
  }

  get visible(): boolean {
    return this.item.visible;
  }

  set visible(val: boolean) {
    this.item.visible = val;
  }

  get text(): string {
    return this.seatText.content;
  }
  set text(val: string) {
    this.seatText.content = val;
  }
}
