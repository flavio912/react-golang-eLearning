/* tslint:disable */
/* eslint-disable */
/* @relayHash 24c249e5803a4069cfb8cbc5ba0eeec9 */

import { ConcreteRequest } from "relay-runtime";
export type AppHolderQueryVariables = {
    name: string;
    offset: number;
};
export type AppHolderQueryResponse = {
    readonly courses: {
        readonly edges: ReadonlyArray<{
            readonly notId: number;
            readonly name: string;
            readonly bannerImageURL: string | null;
            readonly introduction: string | null;
        } | null> | null;
        readonly pageInfo: {
            readonly total: number;
            readonly limit: number;
            readonly offset: number;
            readonly given: number;
        } | null;
    } | null;
};
export type AppHolderQuery = {
    readonly response: AppHolderQueryResponse;
    readonly variables: AppHolderQueryVariables;
};



/*
query AppHolderQuery(
  $name: String!
  $offset: Int!
) {
  courses(filter: {name: $name}, page: {limit: 4, offset: $offset}) {
    edges {
      notId: id
      name
      bannerImageURL
      introduction
    }
    pageInfo {
      total
      limit
      offset
      given
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
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "offset",
            "type": "Int!",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "LinkedField",
            "alias": null,
            "name": "courses",
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
                    "kind": "ObjectValue",
                    "name": "page",
                    "fields": [
                        {
                            "kind": "Literal",
                            "name": "limit",
                            "value": 4
                        },
                        {
                            "kind": "Variable",
                            "name": "offset",
                            "variableName": "offset"
                        }
                    ]
                }
            ],
            "concreteType": "CoursePage",
            "plural": false,
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "edges",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Course",
                    "plural": true,
                    "selections": [
                        {
                            "kind": "ScalarField",
                            "alias": "notId",
                            "name": "id",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "name",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "bannerImageURL",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "introduction",
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
                            "name": "limit",
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
                            "name": "given",
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
            "name": "AppHolderQuery",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "AppHolderQuery",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "query",
            "name": "AppHolderQuery",
            "id": null,
            "text": "query AppHolderQuery(\n  $name: String!\n  $offset: Int!\n) {\n  courses(filter: {name: $name}, page: {limit: 4, offset: $offset}) {\n    edges {\n      notId: id\n      name\n      bannerImageURL\n      introduction\n    }\n    pageInfo {\n      total\n      limit\n      offset\n      given\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '94543c2038b4dbed1f61e898a775f6a5';
export default node;
