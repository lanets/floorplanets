import { toSeatData } from '../src/conversions';

describe('toSeatData', () => {
  let seatState, converted;

  beforeEach(() => {
    seatState = { x: 10, y: 40, id: 1, label: 'G-13' };
    converted = toSeatData(seatState);
  });

  it('removes the position of the seat in the output', () => {
    expect(converted.x).toBeUndefined();
    expect(converted.y).toBeUndefined();
  });

  it('removes the id of the seat in the output', () => {
    expect(converted.id).toBeUndefined();
  });

  it('keeps the label of the seat in the output', () => {
    expect(converted.label).toEqual('G-13');
  });

  it('keeps the data object in the output', () => {
    const data = { username: 'timinou', type: 'VIP' };
    seatState['data'] = data;
    converted = toSeatData(seatState);

    expect(converted.data).toEqual(data);
  });

  it('if data is not set on the seatState, an empty dict is returned', () => {
    expect(converted.data).toEqual({});
  });
});
