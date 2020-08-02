/* tslint:disable */
/* eslint-disable */
/* @relayHash a27599d94f3afd8d92f8b081069672ce */

import { ConcreteRequest } from "relay-runtime";
export type mutations_UpdateDelegateMutationVariables = {
    uuid: string;
    firstName?: string | null;
    lastName?: string | null;
    jobTitle?: string | null;
    email?: string | null;
    phone?: string | null;
};
export type mutations_UpdateDelegateMutationResponse = {
    readonly updateDelegate: {
        readonly firstName: string;
        readonly lastName: string;
        readonly jobTitle: string;
        readonly TTC_ID: string;
        readonly email: string | null;
        readonly telephone: string | null;
    } | null;
};
export type mutations_UpdateDelegateMutation = {
    readonly response: mutations_UpdateDelegateMutationResponse;
    readonly variables: mutations_UpdateDelegateMutationVariables;
};



/*
mutation mutations_UpdateDelegateMutation(
  $uuid: UUID!
  $firstName: String
  $lastName: String
  $jobTitle: String
  $email: String
  $phone: String
) {
  updateDelegate(input: {uuid: $uuid, firstName: $firstName, lastName: $lastName, email: $email, jobTitle: $jobTitle, telephone: $phone}) {
    firstName
    lastName
    jobTitle
    TTC_ID
    email
    telephone
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "uuid",
            "type": "UUID!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "firstName",
            "type": "String",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "lastName",
            "type": "String",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "jobTitle",
            "type": "String",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "email",
            "type": "String",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "phone",
            "type": "String",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "LinkedField",
            "alias": null,
            "name": "updateDelegate",
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
                        },
                        {
                            "kind": "Variable",
                            "name": "uuid",
                            "variableName": "uuid"
                        }
                    ]
                }
            ],
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
                    "name": "jobTitle",
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
        } as any)
    ];
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "mutations_UpdateDelegateMutation",
            "type": "Mutation",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "mutations_UpdateDelegateMutation",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "mutation",
            "name": "mutations_UpdateDelegateMutation",
            "id": null,
            "text": "mutation mutations_UpdateDelegateMutation(\n  $uuid: UUID!\n  $firstName: String\n  $lastName: String\n  $jobTitle: String\n  $email: String\n  $phone: String\n) {\n  updateDelegate(input: {uuid: $uuid, firstName: $firstName, lastName: $lastName, email: $email, jobTitle: $jobTitle, telephone: $phone}) {\n    firstName\n    lastName\n    jobTitle\n    TTC_ID\n    email\n    telephone\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '2ed62300b6c041409a74e81d8c4d6f97';
export default node;
