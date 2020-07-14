/* tslint:disable */
/* eslint-disable */
/* @relayHash 28367fe1faaad226447e09fe0c514ba1 */

import { ConcreteRequest } from "relay-runtime";
export type UserSearchDelegatesQueryVariables = {
    name?: string | null;
};
export type UserSearchDelegatesQueryResponse = {
    readonly delegates: {
        readonly edges: ReadonlyArray<{
            readonly uuid: string;
            readonly TTC_ID: string;
            readonly firstName: string;
            readonly lastName: string;
        } | null> | null;
        readonly pageInfo: {
            readonly total: number;
        } | null;
    } | null;
};
export type UserSearchDelegatesQuery = {
    readonly response: UserSearchDelegatesQueryResponse;
    readonly variables: UserSearchDelegatesQueryVariables;
};



/*
query UserSearchDelegatesQuery(
  $name: String
) {
  delegates(filter: {name: $name}, page: {limit: 8}) {
    edges {
      uuid
      TTC_ID
      firstName
      lastName
    }
    pageInfo {
      total
    }
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "name",
            "type": "String",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "LinkedField",
            "alias": null,
            "name": "delegates",
            "storageKey": null,
            "args": [
                {
                    "kind": "ObjectValue",
                    "name": "filter",
                    "fields": [
                        {
                            "kind": "Variable",
                            "name": "name",
                            "variableName": "name"
                        }
                    ]
                },
                {
                    "kind": "Literal",
                    "name": "page",
                    "value": {
                        "limit": 8
                    }
                }
            ],
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
                            "name": "TTC_ID",
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
                        }
                    ]
                }
            ]
        } as any)
    ];
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "UserSearchDelegatesQuery",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "UserSearchDelegatesQuery",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "query",
            "name": "UserSearchDelegatesQuery",
            "id": null,
            "text": "query UserSearchDelegatesQuery(\n  $name: String\n) {\n  delegates(filter: {name: $name}, page: {limit: 8}) {\n    edges {\n      uuid\n      TTC_ID\n      firstName\n      lastName\n    }\n    pageInfo {\n      total\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'f74214cc2d32a8111ef6645fee026400';
export default node;
