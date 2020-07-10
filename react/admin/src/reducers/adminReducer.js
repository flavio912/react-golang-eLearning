import * as actionTypes from 'src/actions';
import initialState from './initialState';

const adminReducer = (state = initialState.admin, action) => {
  switch (action.type) {
    case actionTypes.GET_ADMIN_LIST_SUCCESS:
      return {
        ...state,
        list: [...action.payload],
        listLoaded: true
      };
    case actionTypes.CREATE_ADMIN_SUCCESS:
      return {
        ...state,
        list: [action.payload, ...state.list]
      };
    case actionTypes.DELETE_ADMIN_SUCCESS:
      return {
        ...state,
        list: [...state.list.filter(item => item.uuid !== action.payload)]
      };
    case actionTypes.UPDATE_ADMIN_SUCCESS:
      return {
        ...state,
        list: [
          ...state.list.map(item => {
            if (item.uuid === action.payload.uuid) return action.payload;
            return item;
          })
        ]
      };
    default: {
      return state;
    }
  }
};

export default adminReducer;
