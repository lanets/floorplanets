/* @flow */


/**
 *  Structure shown to the client to describe a seat object.
 */
export type SeatData = {
  label: string,
  data: Object,
};

/*
 * Description of the configuration object used to configure the
 * Floorplan.
 */
export type FloorplanConfig = {

  // The id of the div on the HTML page in which the floorplan is going to
  // be displayed.
  div: string,

  // Fired when the user selects a seat.
  onSelectSeat: (seat: SeatData) => void,

  // callback defining the color a seat should have
  seatColor: (seat: SeatData) => ?string,

  // callback defining the text to render in the tooltip of a seat
  seatTooltip: (seat: SeatData) => ?string,

  // callback defining the text to render inside the seat.
  seatText: (seat: SeatData) => ?string,
};

