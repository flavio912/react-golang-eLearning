/* tslint:disable */
/* eslint-disable */
/* @relayHash 72e6b7c6de203f4008dbbb4a27fd82e5 */

import { ConcreteRequest } from "relay-runtime";
export type DelegatesPageQueryVariables = {
    name: string;
};
export type DelegatesPageQueryResponse = {
    readonly delegates: {
        readonly edges: ReadonlyArray<{
            readonly uuid: string;
            readonly TTC_ID: string;
            readonly firstName: string;
            readonly lastName: string;
        } | null> | null;
    } | null;
};
export type DelegatesPageQuery = {
    readonly response: DelegatesPageQueryResponse;
    readonly variables: DelegatesPageQueryVariables;
};



/*
query DelegatesPageQuery(
  $name: String!
) {
  delegates(filter: {name: $name}, page: {limit: 8}) {
    edges {
      uuid
      TTC_ID
      firstName
      lastName
    }
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "name",
            "type": "String!",
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
                }
            ]
        } as any)
    ];
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "DelegatesPageQuery",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "DelegatesPageQuery",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "query",
            "name": "DelegatesPageQuery",
            "id": null,
            "text": "query DelegatesPageQuery(\n  $name: String!\n) {\n  delegates(filter: {name: $name}, page: {limit: 8}) {\n    edges {\n      uuid\n      TTC_ID\n      firstName\n      lastName\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '98b7a56254e9f75b28181c08c5c7aede';
export default node;
