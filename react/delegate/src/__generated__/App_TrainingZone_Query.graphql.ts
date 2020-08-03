/* tslint:disable */
/* eslint-disable */
/* @relayHash 3fc7c6bf05200779697dc0026ec6db0c */

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
    status
    enrolledAt
  }
}
*/

const node: ConcreteRequest = ({
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
        "text": "query App_TrainingZone_Query {\n  user {\n    ...TrainingZone_user\n  }\n}\n\nfragment TrainingZone_user on User {\n  firstName\n  myCourses {\n    status\n    enrolledAt\n  }\n}\n",
        "metadata": {}
    }
} as any);
(node as any).hash = '452e5cc310cfa57fc278c5ba3c97d7a3';
export default node;
