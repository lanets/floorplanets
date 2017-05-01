/* @flow */
import reducer from '../../src/reducers/seats';
import { loadSeats, lockSeat } from '../../src/actions/seats';
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


describe('reducer handles LOCK_SEAT action', () => {
  it('it makes a seat unavailable', () => {
    const action = lockSeat(2);
    const before = reducer(initial, loadSeats(seatsData));

    const after = reducer(before, action);
    expect(after.seats[2].available).toEqual(false);
  });
});
