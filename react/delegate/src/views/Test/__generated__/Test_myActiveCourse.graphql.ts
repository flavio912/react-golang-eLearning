/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type Test_myActiveCourse = {
    readonly course: {
        readonly " $fragmentRefs": FragmentRefs<"CourseSyllabusCardFrag_course">;
    };
    readonly " $refType": "Test_myActiveCourse";
};
export type Test_myActiveCourse$data = Test_myActiveCourse;
export type Test_myActiveCourse$key = {
    readonly " $data"?: Test_myActiveCourse$data;
    readonly " $fragmentRefs": FragmentRefs<"Test_myActiveCourse">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "Test_myActiveCourse",
    "type": "MyCourse",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
        {
            "kind": "LinkedField",
            "alias": null,
            "name": "course",
            "storageKey": null,
            "args": null,
            "concreteType": "Course",
            "plural": false,
            "selections": [
                {
                    "kind": "FragmentSpread",
                    "name": "CourseSyllabusCardFrag_course",
                    "args": null
                }
            ]
        }
    ]
} as any);
(node as any).hash = '9a8fdbd131ed0c89a4c503c84792a85c';
export default node;
