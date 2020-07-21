/* tslint:disable */
/* eslint-disable */
/* @relayHash 7c88f5bb451a133bbe2000bfae76a464 */

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
      }
      type
      bannerImageURL
    }
    status
  }
}
*/

const node: ConcreteRequest = ({
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
                                            }
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
        "text": "query App_Courses_Query {\n  user {\n    ...OnlineCourses_user\n  }\n}\n\nfragment OnlineCourses_user on User {\n  firstName\n  myCourses {\n    course {\n      ident: id\n      name\n      excerpt\n      category {\n        color\n      }\n      type\n      bannerImageURL\n    }\n    status\n  }\n}\n",
        "metadata": {}
    }
} as any);
(node as any).hash = 'ea0d86ae1d92dbad653e52da079f95d5';
export default node;
