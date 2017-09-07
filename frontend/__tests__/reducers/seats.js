import reducer from '../../src/reducers/seats';
import { loadSeats } from '../../src/actions/seats';
import { seatsData } from '../../src/__mock__/seats';

import { initial } from '../../src/reducers/seats';

describe('reducer handles LOAD_SEATS action', () => {
  it('It loads the seats contained in the action payload', () => {
    const action = loadSeats(seatsData);
    const before = initial;
    const after = { ...before, seats: seatsData };

    expect(reducer(before, action)).toEqual(after);
  });
});
