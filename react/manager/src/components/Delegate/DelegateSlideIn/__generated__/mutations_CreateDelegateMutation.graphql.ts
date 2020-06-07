/* tslint:disable */
/* eslint-disable */
/* @relayHash d13caf27023a74b544d37a5b521affe4 */

import { ConcreteRequest } from "relay-runtime";
export type mutations_CreateDelegateMutationVariables = {
    firstName: string;
    lastName: string;
    jobTitle: string;
    email: string;
    phone: string;
    generatePassword?: boolean | null;
};
export type mutations_CreateDelegateMutationResponse = {
    readonly createDelegate: {
        readonly delegate: {
            readonly firstName: string;
            readonly lastName: string;
            readonly TTC_ID: string;
            readonly email: string | null;
            readonly telephone: string | null;
        };
        readonly generatedPassword: string | null;
    } | null;
};
export type mutations_CreateDelegateMutation = {
    readonly response: mutations_CreateDelegateMutationResponse;
    readonly variables: mutations_CreateDelegateMutationVariables;
};



/*
mutation mutations_CreateDelegateMutation(
  $firstName: String!
  $lastName: String!
  $jobTitle: String!
  $email: String!
  $phone: String!
  $generatePassword: Boolean
) {
  createDelegate(input: {firstName: $firstName, lastName: $lastName, email: $email, jobTitle: $jobTitle, telephone: $phone, generatePassword: $generatePassword}) {
    delegate {
      firstName
      lastName
      TTC_ID
      email
      telephone
    }
    generatedPassword
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "firstName",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "lastName",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "jobTitle",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "email",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "phone",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "generatePassword",
            "type": "Boolean",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "LinkedField",
            "alias": null,
            "name": "createDelegate",
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
                            "name": "firstName",
                            "variableName": "firstName"
                        },
                        {
                            "kind": "Variable",
                            "name": "generatePassword",
                            "variableName": "generatePassword"
                        },
                        {
                            "kind": "Variable",
                            "name": "jobTitle",
                            "variableName": "jobTitle"
                        },
                        {
                            "kind": "Variable",
                            "name": "lastName",
                            "variableName": "lastName"
                        },
                        {
                            "kind": "Variable",
                            "name": "telephone",
                            "variableName": "phone"
                        }
                    ]
                }
            ],
            "concreteType": "CreateDelegateResponse",
            "plural": false,
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "delegate",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Delegate",
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
                            "name": "TTC_ID",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "email",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "telephone",
                            "args": null,
                            "storageKey": null
                        }
                    ]
                },
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "generatedPassword",
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
            "name": "mutations_CreateDelegateMutation",
            "type": "Mutation",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "mutations_CreateDelegateMutation",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "mutation",
            "name": "mutations_CreateDelegateMutation",
            "id": null,
            "text": "mutation mutations_CreateDelegateMutation(\n  $firstName: String!\n  $lastName: String!\n  $jobTitle: String!\n  $email: String!\n  $phone: String!\n  $generatePassword: Boolean\n) {\n  createDelegate(input: {firstName: $firstName, lastName: $lastName, email: $email, jobTitle: $jobTitle, telephone: $phone, generatePassword: $generatePassword}) {\n    delegate {\n      firstName\n      lastName\n      TTC_ID\n      email\n      telephone\n    }\n    generatedPassword\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '95feec1cdac1ecb372568c00f434df92';
export default node;
