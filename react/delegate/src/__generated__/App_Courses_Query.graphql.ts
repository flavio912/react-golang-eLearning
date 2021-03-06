/* tslint:disable */
/* eslint-disable */
/* @relayHash 2a740ce9f18f5e1f86021ab0c593d7b9 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_Courses_QueryVariables = {};
export type App_Courses_QueryResponse = {
    readonly user: {
        readonly " $fragmentRefs": FragmentRefs<"OnlineCourses_user">;
    } | null;
};
export type App_Courses_Query = {
    readonly response: App_Courses_QueryResponse;
    readonly variables: App_Courses_QueryVariables;
};



/*
query App_Courses_Query {
  user {
    ...OnlineCourses_user
  }
}

fragment OnlineCourses_user on User {
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
            "name": "App_Courses_Query",
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
                            "name": "OnlineCourses_user",
                            "args": null
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "App_Courses_Query",
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
            "name": "App_Courses_Query",
            "id": null,
            "text": "query App_Courses_Query {\n  user {\n    ...OnlineCourses_user\n  }\n}\n\nfragment OnlineCourses_user on User {\n  firstName\n  myCourses {\n    course {\n      ident: id\n      name\n      excerpt\n      category {\n        color\n        name\n      }\n      type\n      bannerImageURL\n    }\n    status\n    progress {\n      percent\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'ea0d86ae1d92dbad653e52da079f95d5';
export default node;
