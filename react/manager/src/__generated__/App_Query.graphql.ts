/* tslint:disable */
/* eslint-disable */
/* @relayHash eff14ac7c99523c5fdee1358f5029c27 */

import { ConcreteRequest } from "relay-runtime";
export type App_QueryVariables = {};
export type App_QueryResponse = {
    readonly info: string;
};
export type App_Query = {
    readonly response: App_QueryResponse;
    readonly variables: App_QueryVariables;
};



/*
query App_Query {
  info
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "ScalarField",
            "alias": null,
            "name": "info",
            "args": null,
            "storageKey": null
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
            "text": "query App_Query {\n  info\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'cb71d961476f1308026f001d073d9121';
export default node;
