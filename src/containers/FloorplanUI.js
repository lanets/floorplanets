// @flow
import { connect } from 'react-redux';
import * as zoomActions from '../actions/zoom';

import FloorplanUI from '../components/FloorplanUI';


function mapDispatchToProps(dispatch: Function) {
  return {
    zoomIn: (value) => dispatch(zoomActions.zoomIn(value)),
    zoomOut: (value) => dispatch(zoomActions.zoomOut(value)),
  };
}

export default connect(
  null,
  mapDispatchToProps,
)(FloorplanUI);
