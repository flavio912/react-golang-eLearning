/* tslint:disable */
/* eslint-disable */
/* @relayHash 326189572e2d6c1236c92df862c955b3 */

import { ConcreteRequest } from "relay-runtime";
export type OrgOverviewDelegatesQueryVariables = {
    name: string;
};
export type OrgOverviewDelegatesQueryResponse = {
    readonly delegates: {
        readonly edges: ReadonlyArray<{
            readonly uuid: string;
            readonly TTC_ID: string;
            readonly firstName: string;
            readonly lastName: string;
        } | null> | null;
    } | null;
};
export type OrgOverviewDelegatesQuery = {
    readonly response: OrgOverviewDelegatesQueryResponse;
    readonly variables: OrgOverviewDelegatesQueryVariables;
};



/*
query OrgOverviewDelegatesQuery(
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
            "name": "OrgOverviewDelegatesQuery",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "OrgOverviewDelegatesQuery",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "query",
            "name": "OrgOverviewDelegatesQuery",
            "id": null,
            "text": "query OrgOverviewDelegatesQuery(\n  $name: String!\n) {\n  delegates(filter: {name: $name}, page: {limit: 8}) {\n    edges {\n      uuid\n      TTC_ID\n      firstName\n      lastName\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '6ef8d5220f1896262eca40cfc9de110d';
export default node;
