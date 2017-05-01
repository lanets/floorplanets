// @flow
import { connect } from 'react-redux';

import Floorplan from '../components/Floorplan';


function mapStateToProps(state: Object, props: Object) {
  return {
    seats: state.seats.seats,
  }
}

export default connect(
  mapStateToProps,
)(Floorplan);
