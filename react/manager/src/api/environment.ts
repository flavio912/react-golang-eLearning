import { Environment, Network, RecordSource, Store } from "relay-runtime";

function fetchQuery(operation: any, variables: any) {
  const headers: any = {
    "content-type": "application/json",
  };

  const token = localStorage.getItem("auth");
  if (token) {
    headers.Authorization = `Bearer ${token}`;
  }
  return fetch("http://localhost:8080/graphql", {
    method: "POST",
    headers,
    body: JSON.stringify({
      query: operation.text,
      variables,
    }),
  }).then((response) => {
    return response.json();
  });
}

const environment = new Environment({
  network: Network.create(fetchQuery),
  store: new Store(new RecordSource()),
});

export default environment;
