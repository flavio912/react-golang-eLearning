/**
 * @flow
 * @relayHash 03f2f0a8a086063281c7cd54451542a2
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
type ExamplePage_propName$ref = any;
export type ExamplePage_QueryVariables = {||};
export type ExamplePage_QueryResponse = {|
  +user: ?{|
    +$fragmentRefs: ExamplePage_propName$ref
  |}
|};
export type ExamplePage_Query = {|
  variables: ExamplePage_QueryVariables,
  response: ExamplePage_QueryResponse,
|};
*/


/*
query ExamplePage_Query {
  user {
    ...ExamplePage_propName
    id
  }
}

fragment ExampleComponent_propName on Todo {
  complete
  text
}

fragment ExamplePage_propName on User {
  todos(first: 2147483647) {
    edges {
      node {
        id
        ...ExampleComponent_propName
      }
    }
  }
  id
  totalCount
  completedCount
}
*/

const node/*: ConcreteRequest*/ = (function(){
var v0 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "id",
  "args": null,
  "storageKey": null
};
return {
  "kind": "Request",
  "fragment": {
    "kind": "Fragment",
    "name": "ExamplePage_Query",
    "type": "Query",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
      {
        "kind": "LinkedField",
        "alias": null,
        "name": "user",
        "storageKey": null,
        "args": null,
        "concreteType": "User",
        "plural": false,
        "selections": [
          {
            "kind": "FragmentSpread",
            "name": "ExamplePage_propName",
            "args": null
          }
        ]
      }
    ]
  },
  "operation": {
    "kind": "Operation",
    "name": "ExamplePage_Query",
    "argumentDefinitions": [],
    "selections": [
      {
        "kind": "LinkedField",
        "alias": null,
        "name": "user",
        "storageKey": null,
        "args": null,
        "concreteType": "User",
        "plural": false,
        "selections": [
          {
            "kind": "LinkedField",
            "alias": null,
            "name": "todos",
            "storageKey": "todos(first:2147483647)",
            "args": [
              {
                "kind": "Literal",
                "name": "first",
                "value": 2147483647
              }
            ],
            "concreteType": "TodoConnection",
            "plural": false,
            "selections": [
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "edges",
                "storageKey": null,
                "args": null,
                "concreteType": "TodoEdge",
                "plural": true,
                "selections": [
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "node",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Todo",
                    "plural": false,
                    "selections": [
                      (v0/*: any*/),
                      {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "complete",
                        "args": null,
                        "storageKey": null
                      },
                      {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "text",
                        "args": null,
                        "storageKey": null
                      }
                    ]
                  }
                ]
              }
            ]
          },
          (v0/*: any*/),
          {
            "kind": "ScalarField",
            "alias": null,
            "name": "totalCount",
            "args": null,
            "storageKey": null
          },
          {
            "kind": "ScalarField",
            "alias": null,
            "name": "completedCount",
            "args": null,
            "storageKey": null
          }
        ]
      }
    ]
  },
  "params": {
    "operationKind": "query",
    "name": "ExamplePage_Query",
    "id": null,
    "text": "query ExamplePage_Query {\n  user {\n    ...ExamplePage_propName\n    id\n  }\n}\n\nfragment ExampleComponent_propName on Todo {\n  complete\n  text\n}\n\nfragment ExamplePage_propName on User {\n  todos(first: 2147483647) {\n    edges {\n      node {\n        id\n        ...ExampleComponent_propName\n      }\n    }\n  }\n  id\n  totalCount\n  completedCount\n}\n",
    "metadata": {}
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = '4f7e0fa9afb53d559f1722e2aa0388a2';

module.exports = node;
