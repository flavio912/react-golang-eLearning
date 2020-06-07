/* tslint:disable */
/* eslint-disable */
/* @relayHash 42929ac79f4f0e5a67b9a75ef187d825 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_DelegatesPage_QueryVariables = {};
export type App_DelegatesPage_QueryResponse = {
    readonly delegates: {
        readonly " $fragmentRefs": FragmentRefs<"DelegatesPage_delegates">;
    } | null;
    readonly manager: {
        readonly " $fragmentRefs": FragmentRefs<"DelegatesPage_manager">;
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
  manager {
    ...DelegatesPage_manager
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

fragment DelegatesPage_manager on Manager {
  company {
    name
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
            },
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "manager",
                "storageKey": null,
                "args": null,
                "concreteType": "Manager",
                "plural": false,
                "selections": [
                    {
                        "kind": "FragmentSpread",
                        "name": "DelegatesPage_manager",
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
            },
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "manager",
                "storageKey": null,
                "args": null,
                "concreteType": "Manager",
                "plural": false,
                "selections": [
                    {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "company",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "Company",
                        "plural": false,
                        "selections": [
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "name",
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
        "text": "query App_DelegatesPage_Query {\n  delegates {\n    ...DelegatesPage_delegates\n  }\n  manager {\n    ...DelegatesPage_manager\n  }\n}\n\nfragment DelegatesPage_delegates on DelegatePage {\n  edges {\n    uuid\n    email\n    firstName\n    lastName\n    lastLogin\n    createdAt\n  }\n  pageInfo {\n    total\n    offset\n    limit\n    given\n  }\n}\n\nfragment DelegatesPage_manager on Manager {\n  company {\n    name\n  }\n}\n",
        "metadata": {}
    }
} as any);
(node as any).hash = '1cfd762ff257bf53ed4aa0afcfa62c34';
export default node;
