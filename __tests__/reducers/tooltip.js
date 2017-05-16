import reducer, { initial } from '../../src/reducers/tooltip';


describe('reducer handles SHOW_TOOLTIP action', () => {
  it('sets the display boolean to true', () => {
    const action = showTooltip('some text goes here');
    const before = initial;
    const after = { ...before, display: true };

    expect(reducer(before, action)).toEqual(after);
  });
});
