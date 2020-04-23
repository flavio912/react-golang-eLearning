/* tslint:disable */
/* eslint-disable */
/* @relayHash 9a07bc52fb48f0e89e8c9ad4af2ec57f */

import { ConcreteRequest } from "relay-runtime";
export type LoginPage_LoginMutationVariables = {
    email: string;
    password: string;
};
export type LoginPage_LoginMutationResponse = {
    readonly managerLogin: {
        readonly token: string;
    } | null;
};
export type LoginPage_LoginMutation = {
    readonly response: LoginPage_LoginMutationResponse;
    readonly variables: LoginPage_LoginMutationVariables;
};



/*
mutation LoginPage_LoginMutation(
  $email: String!
  $password: String!
) {
  managerLogin(input: {email: $email, password: $password}) {
    token
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "email",
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
            "name": "managerLogin",
            "storageKey": null,
            "args": [
                {
                    "kind": "ObjectValue",
                    "name": "input",
                    "fields": [
                        {
                            "kind": "Variable",
                            "name": "email",
                            "variableName": "email"
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
            "text": "mutation LoginPage_LoginMutation(\n  $email: String!\n  $password: String!\n) {\n  managerLogin(input: {email: $email, password: $password}) {\n    token\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'c21c415fb7cc373d41068136c9e932df';
export default node;
