/* tslint:disable */
/* eslint-disable */
/* @relayHash d41b4f0c9173800a20fb659634baae48 */

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
  activeCourses {
    course {
      name
      excerpt
      color
      type
      ident: id
    }
    currentAttempt
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
                        "name": "activeCourses",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "ActiveCourse",
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
                                        "kind": "ScalarField",
                                        "alias": null,
                                        "name": "color",
                                        "args": null,
                                        "storageKey": null
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
                                        "alias": "ident",
                                        "name": "id",
                                        "args": null,
                                        "storageKey": null
                                    }
                                ]
                            },
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "currentAttempt",
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
        "text": "query App_Courses_Query {\n  user {\n    ...OnlineCourses_user\n  }\n}\n\nfragment OnlineCourses_user on User {\n  firstName\n  activeCourses {\n    course {\n      name\n      excerpt\n      color\n      type\n      ident: id\n    }\n    currentAttempt\n  }\n}\n",
        "metadata": {}
    }
} as any);
(node as any).hash = 'ea0d86ae1d92dbad653e52da079f95d5';
export default node;
