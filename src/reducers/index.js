// @flow
import { combineReducers } from 'redux';

import seats from './seats';
import zoom from './zoom';

export default combineReducers({
  seats,
  zoom,
});
