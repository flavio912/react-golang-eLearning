/* tslint:disable */
/* eslint-disable */
/* @relayHash 666fe375cd10da4c0717984fa0f8f487 */

import { ConcreteRequest } from "relay-runtime";
export type Payment_PurchaseMutationVariables = {
    courses: Array<number>;
    users?: Array<string> | null;
    extraEmail?: string | null;
};
export type Payment_PurchaseMutationResponse = {
    readonly purchaseCourses: {
        readonly transactionComplete: boolean;
        readonly stripeClientSecret: string | null;
    } | null;
};
export type Payment_PurchaseMutation = {
    readonly response: Payment_PurchaseMutationResponse;
    readonly variables: Payment_PurchaseMutationVariables;
};



/*
mutation Payment_PurchaseMutation(
  $courses: [Int!]!
  $users: [UUID!]
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
            "type": "[UUID!]",
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
            "name": "Payment_PurchaseMutation",
            "type": "Mutation",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "Payment_PurchaseMutation",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "mutation",
            "name": "Payment_PurchaseMutation",
            "id": null,
            "text": "mutation Payment_PurchaseMutation(\n  $courses: [Int!]!\n  $users: [UUID!]\n  $extraEmail: String\n) {\n  purchaseCourses(input: {courses: $courses, users: $users, extraInvoiceEmail: $extraEmail, acceptedTerms: true, backgroundCheckConfirm: true}) {\n    transactionComplete\n    stripeClientSecret\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '0c99830638316b851b6eb5f59217824c';
export default node;
