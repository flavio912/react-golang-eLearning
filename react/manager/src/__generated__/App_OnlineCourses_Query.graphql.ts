/* tslint:disable */
/* eslint-disable */
/* @relayHash 48776a4bb16e047e0f8c53836aa46ffa */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type Page = {
    offset?: number | null;
    limit?: number | null;
};
export type App_OnlineCourses_QueryVariables = {
    page?: Page | null;
};
export type App_OnlineCourses_QueryResponse = {
    readonly onlineCourses: {
        readonly " $fragmentRefs": FragmentRefs<"CoursesPage_onlineCourses">;
    } | null;
};
export type App_OnlineCourses_Query = {
    readonly response: App_OnlineCourses_QueryResponse;
    readonly variables: App_OnlineCourses_QueryVariables;
};



/*
query App_OnlineCourses_Query(
  $page: Page
) {
  onlineCourses(page: $page) {
    ...CoursesPage_onlineCourses
  }
}

fragment CoursesPage_onlineCourses on OnlineCoursePage {
  edges {
    uuid
    info {
      name
      color
      excerpt
      price
      category {
        name
        color
      }
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
            "name": "App_OnlineCourses_Query",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "onlineCourses",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
                    "concreteType": "OnlineCoursePage",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "FragmentSpread",
                            "name": "CoursesPage_onlineCourses",
                            "args": null
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "App_OnlineCourses_Query",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "onlineCourses",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
                    "concreteType": "OnlineCoursePage",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "edges",
                            "storageKey": null,
                            "args": null,
                            "concreteType": "OnlineCourse",
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
                                    "kind": "LinkedField",
                                    "alias": null,
                                    "name": "info",
                                    "storageKey": null,
                                    "args": null,
                                    "concreteType": "CourseInfo",
                                    "plural": false,
                                    "selections": [
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
            "name": "App_OnlineCourses_Query",
            "id": null,
            "text": "query App_OnlineCourses_Query(\n  $page: Page\n) {\n  onlineCourses(page: $page) {\n    ...CoursesPage_onlineCourses\n  }\n}\n\nfragment CoursesPage_onlineCourses on OnlineCoursePage {\n  edges {\n    uuid\n    info {\n      name\n      color\n      excerpt\n      price\n      category {\n        name\n        color\n      }\n    }\n  }\n  pageInfo {\n    total\n    offset\n    limit\n    given\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'da85bef9d2e027aeb69035eb08fd76c1';
export default node;
