/* tslint:disable */
/* eslint-disable */
/* @relayHash 4a6f3586f28b44b76e35ded9b0277453 */

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
  }
}

fragment TrainingProgress_activity on ActivityPage {
  ...ActivityTable_activity
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
    ];
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
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "type",
                                            "args": null,
                                            "storageKey": null
                                        },
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
                                                {
                                                    "kind": "ScalarField",
                                                    "alias": null,
                                                    "name": "name",
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
                        }
                    ]
                }
            ]
        },
        "params": {
            "operationKind": "query",
            "name": "App_Progress_Query",
            "id": null,
            "text": "query App_Progress_Query(\n  $offset: Int\n  $limit: Int\n) {\n  user {\n    activity(page: {offset: $offset, limit: $limit}) {\n      ...TrainingProgress_activity\n    }\n  }\n}\n\nfragment ActivityTable_activity on ActivityPage {\n  edges {\n    type\n    createdAt\n    course {\n      ident: id\n      name\n    }\n  }\n  pageInfo {\n    total\n  }\n}\n\nfragment TrainingProgress_activity on ActivityPage {\n  ...ActivityTable_activity\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '5791ba32cd9ca0d2e45a9a13abfed66b';
export default node;
