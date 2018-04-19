import { combineReducers } from 'redux'
import { routerReducer } from 'react-router-redux'

import schemes from './schemes';

export default combineReducers({
  routing: routerReducer,
  schemes
})
