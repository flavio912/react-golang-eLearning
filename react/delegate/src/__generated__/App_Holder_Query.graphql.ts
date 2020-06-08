/* tslint:disable */
/* eslint-disable */
/* @relayHash 55e76af96306363a100a293bc8208c56 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_Holder_QueryVariables = {};
export type App_Holder_QueryResponse = {
    readonly user: {
        readonly " $fragmentRefs": FragmentRefs<"AppHolder_user">;
    } | null;
};
export type App_Holder_Query = {
    readonly response: App_Holder_QueryResponse;
    readonly variables: App_Holder_QueryVariables;
};



/*
query App_Holder_Query {
  user {
    ...AppHolder_user
  }
}

fragment AppHolder_user on User {
  firstName
  lastName
  profileImageUrl
}
*/

const node: ConcreteRequest = ({
    "kind": "Request",
    "fragment": {
        "kind": "Fragment",
        "name": "App_Holder_Query",
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
                        "name": "AppHolder_user",
                        "args": null
                    }
                ]
            }
        ]
    },
    "operation": {
        "kind": "Operation",
        "name": "App_Holder_Query",
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
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "lastName",
                        "args": null,
                        "storageKey": null
                    },
                    {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "profileImageUrl",
                        "args": null,
                        "storageKey": null
                    }
                ]
            }
        ]
    },
    "params": {
        "operationKind": "query",
        "name": "App_Holder_Query",
        "id": null,
        "text": "query App_Holder_Query {\n  user {\n    ...AppHolder_user\n  }\n}\n\nfragment AppHolder_user on User {\n  firstName\n  lastName\n  profileImageUrl\n}\n",
        "metadata": {}
    }
} as any);
(node as any).hash = 'cf97afc4f114dab9b297244184c98029';
export default node;
