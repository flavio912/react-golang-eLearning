/* tslint:disable */
/* eslint-disable */
/* @relayHash 7fd9f0134a2fc6e4e6322d94090a33fc */

import { ConcreteRequest } from "relay-runtime";
export type TabsCoursesQueryVariables = {
    name: string;
};
export type TabsCoursesQueryResponse = {
    readonly courses: {
        readonly edges: ReadonlyArray<{
            readonly ident: number;
            readonly name: string;
            readonly price: number;
            readonly category: {
                readonly uuid: string | null;
                readonly name: string;
                readonly color: string;
            } | null;
        } | null> | null;
    } | null;
};
export type TabsCoursesQuery = {
    readonly response: TabsCoursesQueryResponse;
    readonly variables: TabsCoursesQueryVariables;
};



/*
query TabsCoursesQuery(
  $name: String!
) {
  courses(filter: {name: $name}, page: {limit: 60}) {
    edges {
      ident: id
      name
      price
      category {
        uuid
        name
        color
      }
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
    ], v1 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "name",
        "args": null,
        "storageKey": null
    } as any), v2 = [
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
                    "kind": "Literal",
                    "name": "page",
                    "value": {
                        "limit": 60
                    }
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
                        (v1 /*: any*/),
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "price",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "category",
                            "storageKey": null,
                            "args": null,
                            "concreteType": "Category",
                            "plural": false,
                            "selections": [
                                {
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "uuid",
                                    "args": null,
                                    "storageKey": null
                                },
                                (v1 /*: any*/),
                                {
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "color",
                                    "args": null,
                                    "storageKey": null
                                }
                            ]
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
            "name": "TabsCoursesQuery",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v2 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "TabsCoursesQuery",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v2 /*: any*/)
        },
        "params": {
            "operationKind": "query",
            "name": "TabsCoursesQuery",
            "id": null,
            "text": "query TabsCoursesQuery(\n  $name: String!\n) {\n  courses(filter: {name: $name}, page: {limit: 60}) {\n    edges {\n      ident: id\n      name\n      price\n      category {\n        uuid\n        name\n        color\n      }\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '19c1f47d9a8730a8d7bae0124be094ed';
export default node;
