import { FloorplanClient } from '../src/client';


describe('FloorplanClient', () => {
  it('throws an error if div is not specified in the configuration', () => {
    expect(() => {
      const floorplan = new FloorplanClient({});
    }).toThrowError();
  });
});
