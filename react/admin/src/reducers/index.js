import { combineReducers } from 'redux';
import sessionReducer from './sessionReducer';
import adminReducer from './adminReducer';

const rootReducer = combineReducers({
  session: sessionReducer,
  admin: adminReducer
});

export default rootReducer;
