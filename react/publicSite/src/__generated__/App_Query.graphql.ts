/* tslint:disable */
/* eslint-disable */
/* @relayHash 8e201af2e86a3fbe58ac2b84dae85a32 */

import { ConcreteRequest } from "relay-runtime";
export type App_QueryVariables = {};
export type App_QueryResponse = {
    readonly manager: {
        readonly uuid: string;
        readonly firstName: string;
        readonly lastName: string;
    } | null;
};
export type App_Query = {
    readonly response: App_QueryResponse;
    readonly variables: App_QueryVariables;
};



/*
query App_Query {
  manager {
    uuid
    firstName
    lastName
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LinkedField",
            "alias": null,
            "name": "manager",
            "storageKey": null,
            "args": null,
            "concreteType": "Manager",
            "plural": false,
            "selections": [
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "uuid",
                    "args": null,
                    "storageKey": null
                },
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
                }
            ]
        } as any)
    ];
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "App_Query",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": [],
            "selections": (v0 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "App_Query",
            "argumentDefinitions": [],
            "selections": (v0 /*: any*/)
        },
        "params": {
            "operationKind": "query",
            "name": "App_Query",
            "id": null,
            "text": "query App_Query {\n  manager {\n    uuid\n    firstName\n    lastName\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '437b84b52656bf9735572b91f8dffb87';
export default node;
