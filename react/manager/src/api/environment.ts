import { Environment, Network, RecordSource, Store } from 'relay-runtime';
//@ts-ignore
import jwtDecode from 'jwt-decode';

type GraphError = {
  message: string;
  type?: string;
};

export class FetchError extends Error {
  public type: string | boolean = false;
  constructor(error: GraphError) {
    super(error.message); // (1)
    this.type = error?.type || false; // (2)
  }
}

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

async function fetchQuery(operation: any, variables: any) {
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
  }).then(async (response) => {
    const json = await response.json();
    if (json && 'errors' in json) {
      throw new FetchError({
        message:
          json.errors[0]?.extensions?.message ||
          json.errors[0]?.extensions?.message,
        type: json.errors[0]?.extensions?.type
      });
    }
    return json;
  });
}

const environment = new Environment({
  network: Network.create(fetchQuery),
  store: new Store(new RecordSource())
});

export default environment;
