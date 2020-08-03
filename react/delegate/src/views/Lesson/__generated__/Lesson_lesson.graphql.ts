/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type Lesson_lesson = {
    readonly name: string;
    readonly uuid: string;
    readonly voiceoverURL: string | null;
    readonly description: string;
    readonly transcript: string | null;
    readonly " $refType": "Lesson_lesson";
};
export type Lesson_lesson$data = Lesson_lesson;
export type Lesson_lesson$key = {
    readonly " $data"?: Lesson_lesson$data;
    readonly " $fragmentRefs": FragmentRefs<"Lesson_lesson">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "Lesson_lesson",
    "type": "Lesson",
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
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "uuid",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "voiceoverURL",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "description",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "transcript",
            "args": null,
            "storageKey": null
        }
    ]
} as any);
(node as any).hash = '158f6e8025d7b9df4176d8eb9e3bc368';
export default node;
