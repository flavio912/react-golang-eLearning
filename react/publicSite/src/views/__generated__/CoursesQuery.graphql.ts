/* tslint:disable */
/* eslint-disable */
/* @relayHash 1771290d2b48e99b2bc3ba4bc53014db */

import { ConcreteRequest } from "relay-runtime";
export type CoursesQueryVariables = {
    categoryUUID?: string | null;
    offset: number;
};
export type CoursesQueryResponse = {
    readonly courses: {
        readonly edges: ReadonlyArray<{
            readonly ident: number;
            readonly name: string;
            readonly price: number;
            readonly excerpt: string | null;
            readonly introduction: string | null;
            readonly bannerImageURL: string | null;
        } | null> | null;
        readonly pageInfo: {
            readonly total: number;
            readonly offset: number;
            readonly limit: number;
            readonly given: number;
        } | null;
    } | null;
};
export type CoursesQuery = {
    readonly response: CoursesQueryResponse;
    readonly variables: CoursesQueryVariables;
};



/*
query CoursesQuery(
  $categoryUUID: UUID
  $offset: Int!
) {
  courses(filter: {categoryUUID: $categoryUUID}, page: {limit: 5, offset: $offset}) {
    edges {
      ident: id
      name
      price
      excerpt
      introduction
      bannerImageURL
    }
    pageInfo {
      total
      offset
      limit
      given
    }
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "categoryUUID",
            "type": "UUID",
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
                            "name": "categoryUUID",
                            "variableName": "categoryUUID"
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
                            "value": 5
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
                            "alias": "ident",
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
                            "name": "price",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "excerpt",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "introduction",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "bannerImageURL",
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
        } as any)
    ];
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "CoursesQuery",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "CoursesQuery",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "query",
            "name": "CoursesQuery",
            "id": null,
            "text": "query CoursesQuery(\n  $categoryUUID: UUID\n  $offset: Int!\n) {\n  courses(filter: {categoryUUID: $categoryUUID}, page: {limit: 5, offset: $offset}) {\n    edges {\n      ident: id\n      name\n      price\n      excerpt\n      introduction\n      bannerImageURL\n    }\n    pageInfo {\n      total\n      offset\n      limit\n      given\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '757133afea85cbd6ebd73c3f26aa4874';
export default node;
