/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type QuestionType = "SINGLE_CHOICE" | "%future added value";
export type Test_test = {
    readonly name: string;
    readonly uuid: string;
    readonly questions: ReadonlyArray<{
        readonly uuid: string;
        readonly text: string;
        readonly questionType: QuestionType;
        readonly answers: ReadonlyArray<{
            readonly uuid: string;
            readonly text: string | null;
            readonly imageURL: string | null;
        }> | null;
    }> | null;
    readonly " $refType": "Test_test";
};
export type Test_test$data = Test_test;
export type Test_test$key = {
    readonly " $data"?: Test_test$data;
    readonly " $fragmentRefs": FragmentRefs<"Test_test">;
};



const node: ReaderFragment = (function () {
    var v0 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "uuid",
        "args": null,
        "storageKey": null
    } as any), v1 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "text",
        "args": null,
        "storageKey": null
    } as any);
    return {
        "kind": "Fragment",
        "name": "Test_test",
        "type": "Test",
        "metadata": null,
        "argumentDefinitions": [],
        "selections": [
            {
                "kind": "ScalarField",
                "alias": null,
                "name": "name",
                "args": null,
                "storageKey": null
            },
            (v0 /*: any*/),
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "questions",
                "storageKey": null,
                "args": null,
                "concreteType": "Question",
                "plural": true,
                "selections": [
                    (v0 /*: any*/),
                    (v1 /*: any*/),
                    {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "questionType",
                        "args": null,
                        "storageKey": null
                    },
                    {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "answers",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "Answer",
                        "plural": true,
                        "selections": [
                            (v0 /*: any*/),
                            (v1 /*: any*/),
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "imageURL",
                                "args": null,
                                "storageKey": null
                            }
                        ]
                    }
                ]
            }
        ]
    } as any;
})();
(node as any).hash = 'bb19957fbe11770eb893646548e89b4d';
export default node;
