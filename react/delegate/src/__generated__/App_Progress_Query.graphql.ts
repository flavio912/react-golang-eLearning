/* tslint:disable */
/* eslint-disable */
/* @relayHash c3b02856fda26bc31f15043de7825461 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_Progress_QueryVariables = {
    offset?: number | null;
    limit?: number | null;
};
export type App_Progress_QueryResponse = {
    readonly user: {
        readonly activity: {
            readonly " $fragmentRefs": FragmentRefs<"TrainingProgress_activity">;
        } | null;
        readonly " $fragmentRefs": FragmentRefs<"TrainingProgress_user">;
    } | null;
};
export type App_Progress_Query = {
    readonly response: App_Progress_QueryResponse;
    readonly variables: App_Progress_QueryVariables;
};



/*
query App_Progress_Query(
  $offset: Int
  $limit: Int
) {
  user {
    activity(page: {offset: $offset, limit: $limit}) {
      ...TrainingProgress_activity
    }
    ...TrainingProgress_user
  }
}

fragment ActivityTable_activity on ActivityPage {
  edges {
    type
    createdAt
    course {
      ident: id
      name
    }
  }
  pageInfo {
    total
    limit
    offset
  }
}

fragment TrainingProgress_activity on ActivityPage {
  ...ActivityTable_activity
}

fragment TrainingProgress_user on User {
  firstName
  type
  myCourses {
    status
    progress {
      total
      completed
    }
    course {
      name
      category {
        name
      }
    }
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
    ], v2 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "type",
        "args": null,
        "storageKey": null
    } as any), v3 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "name",
        "args": null,
        "storageKey": null
    } as any), v4 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "total",
        "args": null,
        "storageKey": null
    } as any);
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "App_Progress_Query",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "user",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "User",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "activity",
                            "storageKey": null,
                            "args": (v1 /*: any*/),
                            "concreteType": "ActivityPage",
                            "plural": false,
                            "selections": [
                                {
                                    "kind": "FragmentSpread",
                                    "name": "TrainingProgress_activity",
                                    "args": null
                                }
                            ]
                        },
                        {
                            "kind": "FragmentSpread",
                            "name": "TrainingProgress_user",
                            "args": null
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "App_Progress_Query",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "user",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "User",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "activity",
                            "storageKey": null,
                            "args": (v1 /*: any*/),
                            "concreteType": "ActivityPage",
                            "plural": false,
                            "selections": [
                                {
                                    "kind": "LinkedField",
                                    "alias": null,
                                    "name": "edges",
                                    "storageKey": null,
                                    "args": null,
                                    "concreteType": "Activity",
                                    "plural": true,
                                    "selections": [
                                        (v2 /*: any*/),
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
                                            "name": "course",
                                            "storageKey": null,
                                            "args": null,
                                            "concreteType": "Course",
                                            "plural": false,
                                            "selections": [
                                                {
                                                    "kind": "ScalarField",
                                                    "alias": "ident",
                                                    "name": "id",
                                                    "args": null,
                                                    "storageKey": null
                                                },
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
                                        (v4 /*: any*/),
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
                                        }
                                    ]
                                }
                            ]
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "firstName",
                            "args": null,
                            "storageKey": null
                        },
                        (v2 /*: any*/),
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
                                },
                                {
                                    "kind": "LinkedField",
                                    "alias": null,
                                    "name": "progress",
                                    "storageKey": null,
                                    "args": null,
                                    "concreteType": "Progress",
                                    "plural": false,
                                    "selections": [
                                        (v4 /*: any*/),
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "completed",
                                            "args": null,
                                            "storageKey": null
                                        }
                                    ]
                                },
                                {
                                    "kind": "LinkedField",
                                    "alias": null,
                                    "name": "course",
                                    "storageKey": null,
                                    "args": null,
                                    "concreteType": "Course",
                                    "plural": false,
                                    "selections": [
                                        (v3 /*: any*/),
                                        {
                                            "kind": "LinkedField",
                                            "alias": null,
                                            "name": "category",
                                            "storageKey": null,
                                            "args": null,
                                            "concreteType": "Category",
                                            "plural": false,
                                            "selections": [
                                                (v3 /*: any*/)
                                            ]
                                        }
                                    ]
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        "params": {
            "operationKind": "query",
            "name": "App_Progress_Query",
            "id": null,
            "text": "query App_Progress_Query(\n  $offset: Int\n  $limit: Int\n) {\n  user {\n    activity(page: {offset: $offset, limit: $limit}) {\n      ...TrainingProgress_activity\n    }\n    ...TrainingProgress_user\n  }\n}\n\nfragment ActivityTable_activity on ActivityPage {\n  edges {\n    type\n    createdAt\n    course {\n      ident: id\n      name\n    }\n  }\n  pageInfo {\n    total\n    limit\n    offset\n  }\n}\n\nfragment TrainingProgress_activity on ActivityPage {\n  ...ActivityTable_activity\n}\n\nfragment TrainingProgress_user on User {\n  firstName\n  type\n  myCourses {\n    status\n    progress {\n      total\n      completed\n    }\n    course {\n      name\n      category {\n        name\n      }\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '77975e85570e873e96330c7108f6a0ca';
export default node;
