import reducer, { initial } from '../../src/reducers/tooltip';
import { showTooltip } from '../../src/actions/tooltip';


describe('reducer handles SHOW_TOOLTIP action', () => {
  it('sets the display boolean to true', () => {
    const action = showTooltip(50, 100, 'some text goes here');
    const before = initial;
    const after = reducer(before, action);

    expect(after.display).toEqual(true);
  });

  it('sets the x and y values to the action values', () => {
    const action = showTooltip(50, 100, 'some text goes here');
    const before = initial;
    const after = reducer(before, action);

    expect(after.x).toEqual(50);
    expect(after.y).toEqual(100);
  });
});
