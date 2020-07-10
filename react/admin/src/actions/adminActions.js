export const CREATE_ADMIN_SUCCESS = 'CREATE_ADMIN_SUCCESS';
export const UPDATE_ADMIN_SUCCESS = 'UPDATE_ADMIN_SUCCESS';
export const DELETE_ADMIN_SUCCESS = 'DELETE_ADMIN_SUCCESS';
export const GET_ADMIN_LIST_SUCCESS = 'GET_ADMIN_LIST_SUCCESS';

export const createAdminAction = newAdmin => dispatch => {
  dispatch({
    type: CREATE_ADMIN_SUCCESS,
    payload: newAdmin
  });
};

export const updateAdminAction = updateAdmin => dispatch => {
  dispatch({
    type: UPDATE_ADMIN_SUCCESS,
    payload: updateAdmin
  });
};

export const deleteAdminAction = id => dispatch => {
  dispatch({
    type: DELETE_ADMIN_SUCCESS,
    payload: id
  });
};

export const getAdminListAction = adminList => dispatch => {
  dispatch({
    type: GET_ADMIN_LIST_SUCCESS,
    payload: adminList
  });
};
