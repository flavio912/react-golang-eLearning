/* tslint:disable */
/* eslint-disable */
/* @relayHash 78abe42ef5295920bd980e6eab5f3858 */

import { ConcreteRequest } from "relay-runtime";
export type FinaliseLogin_FinaliseMutationVariables = {
    token: string;
    password: string;
};
export type FinaliseLogin_FinaliseMutationResponse = {
    readonly finaliseDelegate: {
        readonly token: string;
    } | null;
};
export type FinaliseLogin_FinaliseMutation = {
    readonly response: FinaliseLogin_FinaliseMutationResponse;
    readonly variables: FinaliseLogin_FinaliseMutationVariables;
};



/*
mutation FinaliseLogin_FinaliseMutation(
  $token: String!
  $password: String!
) {
  finaliseDelegate(input: {token: $token, password: $password}) {
    token
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "token",
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
            "name": "finaliseDelegate",
            "storageKey": null,
            "args": [
                {
                    "kind": "ObjectValue",
                    "name": "input",
                    "fields": [
                        {
                            "kind": "Variable",
                            "name": "password",
                            "variableName": "password"
                        },
                        {
                            "kind": "Variable",
                            "name": "token",
                            "variableName": "token"
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
            "name": "FinaliseLogin_FinaliseMutation",
            "type": "Mutation",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "FinaliseLogin_FinaliseMutation",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "mutation",
            "name": "FinaliseLogin_FinaliseMutation",
            "id": null,
            "text": "mutation FinaliseLogin_FinaliseMutation(\n  $token: String!\n  $password: String!\n) {\n  finaliseDelegate(input: {token: $token, password: $password}) {\n    token\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '58db059572fca8844aa9c7f34daf8f0e';
export default node;
