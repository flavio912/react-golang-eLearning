const BASE_URL = 'https://ostento.io';

export default function fetchData(
  endpoint: string,
  payload = {},
  method = 'get',
  headers = {}
) {
  const token = window.localStorage.getItem('token');
  return fetch(BASE_URL + endpoint, {
    method,
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
      ...headers
    },
    body: method != 'get' ? JSON.stringify(payload) : undefined
  })
    .then(resp => {
      if (resp.status == 500) {
        console.error('500 Error', resp);
        throw '500 Error from server';
      }
      if (resp.status == 422) {
        console.error('Param Error', resp);

        return resp;
      }
      return resp;
    })
    .catch(err => {
      console.error(err);
      return { json: () => ({ success: false, message: 'Connection Error' }) };
    })
    .then(resp => resp.json());
}
