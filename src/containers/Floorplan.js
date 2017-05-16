// @flow
import { connect } from 'react-redux';

import Floorplan from '../components/Floorplan';
import { showTooltip, hideTooltip } from '../actions/tooltip';


function mapStateToProps(state: Object, props: Object) {
  return {
    seats: state.seats.seats,
    zoom: state.zoom.zoom,
  }
}

function mapDispatchToProps(dispatch: Function) {
  return {
    showTooltip: (x: number, y:number, text: string ) => dispatch(showTooltip(x, y, text)),
    hideTooltip: () => dispatch(hideTooltip()),
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(Floorplan);
