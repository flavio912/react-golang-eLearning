/* tslint:disable */
/* eslint-disable */
/* @relayHash 5247563776f66063d1a26672d6c1a71d */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_DelegatesPage_QueryVariables = {};
export type App_DelegatesPage_QueryResponse = {
    readonly delegates: {
        readonly " $fragmentRefs": FragmentRefs<"DelegatesPage_delegates">;
    } | null;
};
export type App_DelegatesPage_Query = {
    readonly response: App_DelegatesPage_QueryResponse;
    readonly variables: App_DelegatesPage_QueryVariables;
};



/*
query App_DelegatesPage_Query {
  delegates {
    ...DelegatesPage_delegates
  }
}

fragment DelegatesPage_delegates on DelegatePage {
  edges {
    uuid
    email
    firstName
    lastName
    lastLogin
    createdAt
  }
  pageInfo {
    total
    offset
    limit
    given
  }
}
*/

const node: ConcreteRequest = ({
    "kind": "Request",
    "fragment": {
        "kind": "Fragment",
        "name": "App_DelegatesPage_Query",
        "type": "Query",
        "metadata": null,
        "argumentDefinitions": [],
        "selections": [
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "delegates",
                "storageKey": null,
                "args": null,
                "concreteType": "DelegatePage",
                "plural": false,
                "selections": [
                    {
                        "kind": "FragmentSpread",
                        "name": "DelegatesPage_delegates",
                        "args": null
                    }
                ]
            }
        ]
    },
    "operation": {
        "kind": "Operation",
        "name": "App_DelegatesPage_Query",
        "argumentDefinitions": [],
        "selections": [
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "delegates",
                "storageKey": null,
                "args": null,
                "concreteType": "DelegatePage",
                "plural": false,
                "selections": [
                    {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "edges",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "Delegate",
                        "plural": true,
                        "selections": [
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "uuid",
                                "args": null,
                                "storageKey": null
                            },
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "email",
                                "args": null,
                                "storageKey": null
                            },
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "firstName",
                                "args": null,
                                "storageKey": null
                            },
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "lastName",
                                "args": null,
                                "storageKey": null
                            },
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "lastLogin",
                                "args": null,
                                "storageKey": null
                            },
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "createdAt",
                                "args": null,
                                "storageKey": null
                            }
                        ]
                    },
                    {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "pageInfo",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "PageInfo",
                        "plural": false,
                        "selections": [
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "total",
                                "args": null,
                                "storageKey": null
                            },
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "offset",
                                "args": null,
                                "storageKey": null
                            },
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "limit",
                                "args": null,
                                "storageKey": null
                            },
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "given",
                                "args": null,
                                "storageKey": null
                            }
                        ]
                    }
                ]
            }
        ]
    },
    "params": {
        "operationKind": "query",
        "name": "App_DelegatesPage_Query",
        "id": null,
        "text": "query App_DelegatesPage_Query {\n  delegates {\n    ...DelegatesPage_delegates\n  }\n}\n\nfragment DelegatesPage_delegates on DelegatePage {\n  edges {\n    uuid\n    email\n    firstName\n    lastName\n    lastLogin\n    createdAt\n  }\n  pageInfo {\n    total\n    offset\n    limit\n    given\n  }\n}\n",
        "metadata": {}
    }
} as any);
(node as any).hash = '78c5bfbba1cc22166fb8e52b95eac228';
export default node;
