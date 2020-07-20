/* tslint:disable */
/* eslint-disable */
/* @relayHash 679154bccc5b30192721b786023f53d4 */

import { ConcreteRequest } from "relay-runtime";
export type CourseStatus = "failed" | "incomplete" | "passed" | "%future added value";
export type QuestionAnswer = {
    questionUUID: string;
    answerUUID: string;
};
export type Test_SubmitAnswersMutationVariables = {
    courseID: number;
    testUUID: string;
    answers: Array<QuestionAnswer>;
};
export type Test_SubmitAnswersMutationResponse = {
    readonly submitTest: {
        readonly courseStatus: CourseStatus;
        readonly passed: boolean;
    } | null;
};
export type Test_SubmitAnswersMutation = {
    readonly response: Test_SubmitAnswersMutationResponse;
    readonly variables: Test_SubmitAnswersMutationVariables;
};



/*
mutation Test_SubmitAnswersMutation(
  $courseID: Int!
  $testUUID: UUID!
  $answers: [QuestionAnswer!]!
) {
  submitTest(input: {courseID: $courseID, testUUID: $testUUID, answers: $answers}) {
    courseStatus
    passed
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "courseID",
            "type": "Int!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "testUUID",
            "type": "UUID!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "answers",
            "type": "[QuestionAnswer!]!",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "LinkedField",
            "alias": null,
            "name": "submitTest",
            "storageKey": null,
            "args": [
                {
                    "kind": "ObjectValue",
                    "name": "input",
                    "fields": [
                        {
                            "kind": "Variable",
                            "name": "answers",
                            "variableName": "answers"
                        },
                        {
                            "kind": "Variable",
                            "name": "courseID",
                            "variableName": "courseID"
                        },
                        {
                            "kind": "Variable",
                            "name": "testUUID",
                            "variableName": "testUUID"
                        }
                    ]
                }
            ],
            "concreteType": "SubmitTestPayload",
            "plural": false,
            "selections": [
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "courseStatus",
                    "args": null,
                    "storageKey": null
                },
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "passed",
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
            "name": "Test_SubmitAnswersMutation",
            "type": "Mutation",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "Test_SubmitAnswersMutation",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "mutation",
            "name": "Test_SubmitAnswersMutation",
            "id": null,
            "text": "mutation Test_SubmitAnswersMutation(\n  $courseID: Int!\n  $testUUID: UUID!\n  $answers: [QuestionAnswer!]!\n) {\n  submitTest(input: {courseID: $courseID, testUUID: $testUUID, answers: $answers}) {\n    courseStatus\n    passed\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'bca4eefc5841a77ba5dbbd2960f4d7ad';
export default node;
