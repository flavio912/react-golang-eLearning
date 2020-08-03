/* tslint:disable */
/* eslint-disable */
/* @relayHash e2408f5be56c74f8a0aebac9136ae2f5 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_TrainingZone_QueryVariables = {};
export type App_TrainingZone_QueryResponse = {
    readonly user: {
        readonly " $fragmentRefs": FragmentRefs<"TrainingZone_user">;
    } | null;
};
export type App_TrainingZone_Query = {
    readonly response: App_TrainingZone_QueryResponse;
    readonly variables: App_TrainingZone_QueryVariables;
};



/*
query App_TrainingZone_Query {
  user {
    ...TrainingZone_user
  }
}

fragment TrainingZone_user on User {
  firstName
  myCourses {
    course {
      ident: id
      name
      excerpt
      category {
        color
        name
      }
      type
      bannerImageURL
    }
    status
    enrolledAt
    progress {
      percent
    }
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = ({
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
            "name": "App_TrainingZone_Query",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": [],
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
                            "kind": "FragmentSpread",
                            "name": "TrainingZone_user",
                            "args": null
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "App_TrainingZone_Query",
            "argumentDefinitions": [],
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
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "firstName",
                            "args": null,
                            "storageKey": null
                        },
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
                                        (v0 /*: any*/),
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "excerpt",
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
                                                    "name": "color",
                                                    "args": null,
                                                    "storageKey": null
                                                },
                                                (v0 /*: any*/)
                                            ]
                                        },
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
                                            "name": "bannerImageURL",
                                            "args": null,
                                            "storageKey": null
                                        }
                                    ]
                                },
                                {
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "status",
                                    "args": null,
                                    "storageKey": null
                                },
                                {
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "enrolledAt",
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
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "percent",
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
            "name": "App_TrainingZone_Query",
            "id": null,
            "text": "query App_TrainingZone_Query {\n  user {\n    ...TrainingZone_user\n  }\n}\n\nfragment TrainingZone_user on User {\n  firstName\n  myCourses {\n    course {\n      ident: id\n      name\n      excerpt\n      category {\n        color\n        name\n      }\n      type\n      bannerImageURL\n    }\n    status\n    enrolledAt\n    progress {\n      percent\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '452e5cc310cfa57fc278c5ba3c97d7a3';
export default node;
