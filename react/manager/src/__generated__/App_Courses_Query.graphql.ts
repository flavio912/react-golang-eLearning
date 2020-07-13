/* tslint:disable */
/* eslint-disable */
/* @relayHash 1cd6bfcb53b6e69d68ef8e1b16bc6efd */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type Page = {
    offset?: number | null;
    limit?: number | null;
};
export type App_Courses_QueryVariables = {
    page?: Page | null;
};
export type App_Courses_QueryResponse = {
    readonly courses: {
        readonly " $fragmentRefs": FragmentRefs<"CoursesPage_courses">;
    } | null;
};
export type App_Courses_Query = {
    readonly response: App_Courses_QueryResponse;
    readonly variables: App_Courses_QueryVariables;
};



/*
query App_Courses_Query(
  $page: Page
) {
  courses(page: $page) {
    ...CoursesPage_courses
  }
}

fragment CoursesPage_courses on CoursePage {
  edges {
    ident: id
    name
    color
    excerpt
    price
    category {
      name
      color
    }
  }
  pageInfo {
    total
    offset
    limit
    given
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "page",
            "type": "Page",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "Variable",
            "name": "page",
            "variableName": "page"
        } as any)
    ], v2 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "name",
        "args": null,
        "storageKey": null
    } as any), v3 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "color",
        "args": null,
        "storageKey": null
    } as any);
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "App_Courses_Query",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "courses",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
                    "concreteType": "CoursePage",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "FragmentSpread",
                            "name": "CoursesPage_courses",
                            "args": null
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "App_Courses_Query",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "courses",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
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
                                (v2 /*: any*/),
                                (v3 /*: any*/),
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
                                        (v2 /*: any*/),
                                        (v3 /*: any*/)
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
                }
            ]
        },
        "params": {
            "operationKind": "query",
            "name": "App_Courses_Query",
            "id": null,
            "text": "query App_Courses_Query(\n  $page: Page\n) {\n  courses(page: $page) {\n    ...CoursesPage_courses\n  }\n}\n\nfragment CoursesPage_courses on CoursePage {\n  edges {\n    ident: id\n    name\n    color\n    excerpt\n    price\n    category {\n      name\n      color\n    }\n  }\n  pageInfo {\n    total\n    offset\n    limit\n    given\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '3523ef4afe25f69c3cfc73b1c140d579';
export default node;
