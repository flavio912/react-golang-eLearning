/* tslint:disable */
/* eslint-disable */
/* @relayHash d3cf219b70a67bbce98253b005fdabc1 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type AccessType = "open" | "restricted" | "%future added value";
export type Page = {
    offset?: number | null;
    limit?: number | null;
};
export type CoursesFilter = {
    name?: string | null;
    accessType?: AccessType | null;
    backgroundCheck?: boolean | null;
    price?: number | null;
    allowedToBuy?: boolean | null;
    categoryUUID?: string | null;
};
export type App_Courses_QueryVariables = {
    page?: Page | null;
    filter?: CoursesFilter | null;
    categoryUUID: string;
};
export type App_Courses_QueryResponse = {
    readonly courses: {
        readonly " $fragmentRefs": FragmentRefs<"Courses_courses">;
    } | null;
    readonly category: {
        readonly " $fragmentRefs": FragmentRefs<"Courses_category">;
    } | null;
};
export type App_Courses_Query = {
    readonly response: App_Courses_QueryResponse;
    readonly variables: App_Courses_QueryVariables;
};



/*
query App_Courses_Query(
  $page: Page
  $filter: CoursesFilter
  $categoryUUID: UUID!
) {
  courses(page: $page, filter: $filter) {
    ...Courses_courses
  }
  category(uuid: $categoryUUID) {
    ...Courses_category
  }
}

fragment Courses_category on Category {
  uuid
  name
  color
}

fragment Courses_courses on CoursePage {
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
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "page",
            "type": "Page",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "filter",
            "type": "CoursesFilter",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "categoryUUID",
            "type": "UUID!",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "Variable",
            "name": "filter",
            "variableName": "filter"
        } as any),
        ({
            "kind": "Variable",
            "name": "page",
            "variableName": "page"
        } as any)
    ], v2 = [
        ({
            "kind": "Variable",
            "name": "uuid",
            "variableName": "categoryUUID"
        } as any)
    ], v3 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "name",
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
                            "name": "Courses_courses",
                            "args": null
                        }
                    ]
                },
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "category",
                    "storageKey": null,
                    "args": (v2 /*: any*/),
                    "concreteType": "Category",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "FragmentSpread",
                            "name": "Courses_category",
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
                                (v3 /*: any*/),
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
                },
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "category",
                    "storageKey": null,
                    "args": (v2 /*: any*/),
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
                        (v3 /*: any*/),
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
        },
        "params": {
            "operationKind": "query",
            "name": "App_Courses_Query",
            "id": null,
            "text": "query App_Courses_Query(\n  $page: Page\n  $filter: CoursesFilter\n  $categoryUUID: UUID!\n) {\n  courses(page: $page, filter: $filter) {\n    ...Courses_courses\n  }\n  category(uuid: $categoryUUID) {\n    ...Courses_category\n  }\n}\n\nfragment Courses_category on Category {\n  uuid\n  name\n  color\n}\n\nfragment Courses_courses on CoursePage {\n  edges {\n    ident: id\n    name\n    price\n    excerpt\n    introduction\n    bannerImageURL\n  }\n  pageInfo {\n    total\n    offset\n    limit\n    given\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'fc95ad93c6f680a588152ecc45308a49';
export default node;
