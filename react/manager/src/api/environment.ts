import { Environment, Network, RecordSource, Store } from 'relay-runtime';

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

function fetchQuery(operation: any, variables: any) {
  const headers: any = {
    'content-type': 'application/json'
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
