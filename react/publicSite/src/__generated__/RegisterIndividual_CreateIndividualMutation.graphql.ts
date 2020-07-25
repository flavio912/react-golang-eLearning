/* tslint:disable */
/* eslint-disable */
/* @relayHash 2aad064b97e92d217e24685df3c07ba5 */

import { ConcreteRequest } from "relay-runtime";
export type RegisterIndividual_CreateIndividualMutationVariables = {
    firstName: string;
    lastName: string;
    email: string;
    password: string;
};
export type RegisterIndividual_CreateIndividualMutationResponse = {
    readonly createIndividual: {
        readonly user: {
            readonly email: string;
        } | null;
    } | null;
};
export type RegisterIndividual_CreateIndividualMutation = {
    readonly response: RegisterIndividual_CreateIndividualMutationResponse;
    readonly variables: RegisterIndividual_CreateIndividualMutationVariables;
};



/*
mutation RegisterIndividual_CreateIndividualMutation(
  $firstName: String!
  $lastName: String!
  $email: String!
  $password: String!
) {
  createIndividual(input: {firstName: $firstName, lastName: $lastName, email: $email, password: $password}) {
    user {
      __typename
      email
    }
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
                    "name": "lastName",
                    "variableName": "lastName"
                },
                {
                    "kind": "Variable",
                    "name": "password",
                    "variableName": "password"
                }
            ]
        } as any)
    ], v2 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "email",
        "args": null,
        "storageKey": null
    } as any);
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "RegisterIndividual_CreateIndividualMutation",
            "type": "Mutation",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "createIndividual",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
                    "concreteType": "CreateIndividualResponse",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "user",
                            "storageKey": null,
                            "args": null,
                            "concreteType": null,
                            "plural": false,
                            "selections": [
                                (v2 /*: any*/)
                            ]
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "RegisterIndividual_CreateIndividualMutation",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "createIndividual",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
                    "concreteType": "CreateIndividualResponse",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "user",
                            "storageKey": null,
                            "args": null,
                            "concreteType": null,
                            "plural": false,
                            "selections": [
                                {
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "__typename",
                                    "args": null,
                                    "storageKey": null
                                },
                                (v2 /*: any*/)
                            ]
                        }
                    ]
                }
            ]
        },
        "params": {
            "operationKind": "mutation",
            "name": "RegisterIndividual_CreateIndividualMutation",
            "id": null,
            "text": "mutation RegisterIndividual_CreateIndividualMutation(\n  $firstName: String!\n  $lastName: String!\n  $email: String!\n  $password: String!\n) {\n  createIndividual(input: {firstName: $firstName, lastName: $lastName, email: $email, password: $password}) {\n    user {\n      __typename\n      email\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '855235d6ac4704d489b0b759ee470d7b';
export default node;
