import { toSeatData } from '../src/conversions';

describe('toSeatData', () => {
  let seatState, converted;

  beforeEach(() => {
    seatState = { x: 10, y: 40, id: 1, label: 'G-13' };
    converted = toSeatData(seatState);
  });

  it('conversions removes the position of the seat in the output', () => {
    expect(converted.x).toBeUndefined();
    expect(converted.y).toBeUndefined();
  });

  it('conversions removes the id of the seat in the output', () => {
    expect(converted.id).toBeUndefined();
  });

  it('conversions keeps the label of the seat in the output', () => {
    expect(converted.label).toEqual('G-13');
  });
});
