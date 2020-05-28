/* tslint:disable */
/* eslint-disable */
/* @relayHash 920b1935a145a37f17f83ed24de70569 */

import { ConcreteRequest } from "relay-runtime";
export type mutations_CreateDelegateMutationVariables = {
    firstName: string;
    lastName: string;
    jobTitle: string;
    email: string;
    phone: string;
};
export type mutations_CreateDelegateMutationResponse = {
    readonly createDelegate: {
        readonly firstName: string;
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
) {
  createDelegate(input: {firstName: $firstName, lastName: $lastName, email: $email, jobTitle: $jobTitle, telephone: $phone}) {
    firstName
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
            "concreteType": "Delegate",
            "plural": false,
            "selections": [
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "firstName",
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
            "text": "mutation mutations_CreateDelegateMutation(\n  $firstName: String!\n  $lastName: String!\n  $jobTitle: String!\n  $email: String!\n  $phone: String!\n) {\n  createDelegate(input: {firstName: $firstName, lastName: $lastName, email: $email, jobTitle: $jobTitle, telephone: $phone}) {\n    firstName\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'c161828a23fe94bbe703a124cc89ffae';
export default node;
