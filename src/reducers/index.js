// @flow
import { combineReducers } from 'redux';

import seats from './seats';
import zoom from './zoom';
import tooltip from './tooltip';

export default combineReducers({
  seats,
  zoom,
  tooltip,
});
