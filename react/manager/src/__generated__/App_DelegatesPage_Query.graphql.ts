/* tslint:disable */
/* eslint-disable */
/* @relayHash 0c3a9fc492212be9ba91b648a129884b */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_DelegatesPage_QueryVariables = {
    offset?: number | null;
    limit?: number | null;
};
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
query App_DelegatesPage_Query(
  $offset: Int
  $limit: Int
) {
  delegates(page: {offset: $offset, limit: $limit}) {
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
    myCourses {
      status
    }
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

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "offset",
            "type": "Int",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "limit",
            "type": "Int",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "ObjectValue",
            "name": "page",
            "fields": [
                {
                    "kind": "Variable",
                    "name": "limit",
                    "variableName": "limit"
                },
                {
                    "kind": "Variable",
                    "name": "offset",
                    "variableName": "offset"
                }
            ]
        } as any)
    ];
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "App_DelegatesPage_Query",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "delegates",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
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
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "delegates",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
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
                                },
                                {
                                    "kind": "LinkedField",
                                    "alias": null,
                                    "name": "myCourses",
                                    "storageKey": null,
                                    "args": null,
                                    "concreteType": "MyCourse",
                                    "plural": true,
                                    "selections": [
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "status",
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
            "text": "query App_DelegatesPage_Query(\n  $offset: Int\n  $limit: Int\n) {\n  delegates(page: {offset: $offset, limit: $limit}) {\n    ...DelegatesPage_delegates\n  }\n  manager {\n    ...DelegatesPage_manager\n  }\n}\n\nfragment DelegatesPage_delegates on DelegatePage {\n  edges {\n    uuid\n    email\n    firstName\n    lastName\n    lastLogin\n    createdAt\n    myCourses {\n      status\n    }\n  }\n  pageInfo {\n    total\n    offset\n    limit\n    given\n  }\n}\n\nfragment DelegatesPage_manager on Manager {\n  company {\n    name\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '93b668abafcf6598c7f307aa7c28d715';
export default node;
