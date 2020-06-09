import fetchData from './index';

const endpoints = {
  login: '/auth/login',
  getUser: '/auth/getUser'
};

export const login = async (email: string, password: string) => {
  const resp = await fetchData(endpoints.login, { email, password }, 'post');
  if (resp && resp.success) {
    window.localStorage.setItem('token', resp.token);
    window.localStorage.setItem('organisationUuid', resp.organisationUuid);
    return true;
  } else {
    return false;
  }
};

export const getCurrentUser = async () => {
  const resp = await fetchData(endpoints.getUser, {}, 'get');
  if (resp && resp.success) {
    return resp.data;
  } else {
    return false;
  }
};
