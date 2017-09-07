import reducer, { initial } from '../../src/reducers/zoom';
import { zoomIn, zoomOut } from '../../src/actions/zoom';


describe('reducer handles ZOOM_IN action', () => {
  it('increment the zoom setting by the value provided in the action payload', () => {
    const action = zoomIn(0.5);
    const before = initial;
    const after = { ...before, zoom: 1.5 };

    expect(reducer(before, action)).toEqual(after);
  });
});

describe('reducer handles ZOOM_OUT action', () => {
  it('increment the zoom setting by the value provided in the action payload', () => {
    const action = zoomOut(0.5);
    const before = initial;
    const after = { ...before, zoom: 0.5 };

    expect(reducer(before, action)).toEqual(after);
  });

  it('prevents the zoom value to go below 0', () => {
    const action = zoomOut(0.5);

    const before = initial;
    before.zoom = 0.5; // initial zoom is now set a 0.5

    const after = { ...before, zoom: 0.5 }; // the zoom should not go below or equal to 0, therefore should stay the same.

    expect(reducer(before, action)).toEqual(after);
  });
});
