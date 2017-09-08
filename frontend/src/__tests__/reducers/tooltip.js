import reducer, { initial } from '../../reducers/tooltip';
import { showTooltip, hideTooltip } from '../../actions/tooltip';


describe('reducer handles SHOW_TOOLTIP action', () => {
  it('sets the display boolean to true', () => {
    const action = showTooltip('some text goes here');
    const before = initial;
    const after = reducer(before, action);

    expect(after.display).toEqual(true);
  });

  it('sets the text values to the action value', () => {
    const action = showTooltip('some text goes here');
    const before = initial;
    const after = reducer(before, action);

    expect(after.text).toEqual('some text goes here');
  });
});

describe('reducer handles HIDE_TOOLTIP action', () => {
  it('sets the display boolean to false', () => {
    const action = hideTooltip();
    const before = { ...initial, display: true };
    const after = reducer(before, action);

    expect(after.display).toEqual(false);
  });
});
