/* tslint:disable */
/* eslint-disable */
/* @relayHash f536606fb54cd93aa2df4ac00fac653f */

import { ConcreteRequest } from "relay-runtime";
export type LoginPage_LoginMutationVariables = {
    ttcId: string;
    password: string;
};
export type LoginPage_LoginMutationResponse = {
    readonly delegateLogin: {
        readonly token: string;
    } | null;
};
export type LoginPage_LoginMutation = {
    readonly response: LoginPage_LoginMutationResponse;
    readonly variables: LoginPage_LoginMutationVariables;
};



/*
mutation LoginPage_LoginMutation(
  $ttcId: String!
  $password: String!
) {
  delegateLogin(input: {TTC_ID: $ttcId, password: $password}) {
    token
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "ttcId",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "password",
            "type": "String!",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "LinkedField",
            "alias": null,
            "name": "delegateLogin",
            "storageKey": null,
            "args": [
                {
                    "kind": "ObjectValue",
                    "name": "input",
                    "fields": [
                        {
                            "kind": "Variable",
                            "name": "TTC_ID",
                            "variableName": "ttcId"
                        },
                        {
                            "kind": "Variable",
                            "name": "password",
                            "variableName": "password"
                        }
                    ]
                }
            ],
            "concreteType": "AuthToken",
            "plural": false,
            "selections": [
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "token",
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
            "name": "LoginPage_LoginMutation",
            "type": "Mutation",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "LoginPage_LoginMutation",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "mutation",
            "name": "LoginPage_LoginMutation",
            "id": null,
            "text": "mutation LoginPage_LoginMutation(\n  $ttcId: String!\n  $password: String!\n) {\n  delegateLogin(input: {TTC_ID: $ttcId, password: $password}) {\n    token\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '82eaa1d56fabdd9a6ab00ac3760a962d';
export default node;
