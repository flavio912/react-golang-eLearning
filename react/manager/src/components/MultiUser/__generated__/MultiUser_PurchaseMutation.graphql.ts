/* tslint:disable */
/* eslint-disable */
/* @relayHash c17a3059c211995effe7e9330e077c19 */

import { ConcreteRequest } from "relay-runtime";
export type MultiUser_PurchaseMutationVariables = {
    courses: Array<number>;
    users: Array<string>;
    extraEmail?: string | null;
};
export type MultiUser_PurchaseMutationResponse = {
    readonly purchaseCourses: {
        readonly transactionComplete: boolean;
        readonly stripeClientSecret: string | null;
    } | null;
};
export type MultiUser_PurchaseMutation = {
    readonly response: MultiUser_PurchaseMutationResponse;
    readonly variables: MultiUser_PurchaseMutationVariables;
};



/*
mutation MultiUser_PurchaseMutation(
  $courses: [Int!]!
  $users: [UUID!]!
  $extraEmail: String
) {
  purchaseCourses(input: {courses: $courses, users: $users, extraInvoiceEmail: $extraEmail, acceptedTerms: true, backgroundCheckConfirm: true}) {
    transactionComplete
    stripeClientSecret
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "courses",
            "type": "[Int!]!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "users",
            "type": "[UUID!]!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "extraEmail",
            "type": "String",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "LinkedField",
            "alias": null,
            "name": "purchaseCourses",
            "storageKey": null,
            "args": [
                {
                    "kind": "ObjectValue",
                    "name": "input",
                    "fields": [
                        {
                            "kind": "Literal",
                            "name": "acceptedTerms",
                            "value": true
                        },
                        {
                            "kind": "Literal",
                            "name": "backgroundCheckConfirm",
                            "value": true
                        },
                        {
                            "kind": "Variable",
                            "name": "courses",
                            "variableName": "courses"
                        },
                        {
                            "kind": "Variable",
                            "name": "extraInvoiceEmail",
                            "variableName": "extraEmail"
                        },
                        {
                            "kind": "Variable",
                            "name": "users",
                            "variableName": "users"
                        }
                    ]
                }
            ],
            "concreteType": "PurchaseCoursesResponse",
            "plural": false,
            "selections": [
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "transactionComplete",
                    "args": null,
                    "storageKey": null
                },
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "stripeClientSecret",
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
            "name": "MultiUser_PurchaseMutation",
            "type": "Mutation",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "MultiUser_PurchaseMutation",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "mutation",
            "name": "MultiUser_PurchaseMutation",
            "id": null,
            "text": "mutation MultiUser_PurchaseMutation(\n  $courses: [Int!]!\n  $users: [UUID!]!\n  $extraEmail: String\n) {\n  purchaseCourses(input: {courses: $courses, users: $users, extraInvoiceEmail: $extraEmail, acceptedTerms: true, backgroundCheckConfirm: true}) {\n    transactionComplete\n    stripeClientSecret\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'c676cad525bd94e422aa6e68221916ea';
export default node;
