import reducer, { initial } from '../../src/reducers/tooltip';
import { showTooltip } from '../../src/actions/tooltip';


describe('Passing an empty state to the tooltip reducer', () => {
  it('uses a default `initial` state', () => {
    const action = showTooltip('some text goes here');
    const before = undefined;
    const after = { ...before, display: true };

    expect(reducer(before, action)).toEqual(after);
  });
});

describe('reducer handles SHOW_TOOLTIP action', () => {
  it('sets the display boolean to true', () => {
    const action = showTooltip(50, 100, 'some text goes here');
    const before = initial;
    const after = { ...before, display: true, x: 50, y: 100 };

    expect(reducer(before, action)).toEqual(after);
  });
});
