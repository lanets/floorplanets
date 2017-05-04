/* @flow */


/*
 * Description of the configuration object used to configure the
 * Floorplan.
 */
export type FloorplanConfig = {

  // The id of the div on your page in which you want to display
  // the floorplan.
  div: string,
};

/*
 * The data binding between the state of a seat and the graphic
 * representation of a seat.
 */
export type SeatData = {
  id: number,
  x: number,
  y: number,
  label: string,  // what's written on the seat
};
