import { Environment, Network, RecordSource, Store } from 'relay-runtime';

function readCookie(name: string) {
  var nameEQ = name + '=';
  var ca = document.cookie.split(';');
  for (var i = 0; i < ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == ' ') c = c.substring(1, c.length);
    if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length, c.length);
  }
  return null;
}

function fetchQuery(operation: any, variables: any) {
  const CSRF_TOKEN = readCookie('csrf');

  const headers: any = {
    'content-type': 'application/json',
    'X-CSRF-TOKEN': CSRF_TOKEN
  };

  return fetch('http://localhost:8080/graphql', {
    method: 'POST',
    headers,
    body: JSON.stringify({
      query: operation.text,
      variables
    }),
    credentials: 'include'
  }).then((response) => {
    return response.json();
  });
}

const environment = new Environment({
  network: Network.create(fetchQuery),
  store: new Store(new RecordSource())
});

export default environment;
