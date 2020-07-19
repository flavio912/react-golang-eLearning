/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type Module_myActiveCourse = {
    readonly course: {
        readonly " $fragmentRefs": FragmentRefs<"CourseSyllabusCardFrag_course">;
    };
    readonly " $refType": "Module_myActiveCourse";
};
export type Module_myActiveCourse$data = Module_myActiveCourse;
export type Module_myActiveCourse$key = {
    readonly " $data"?: Module_myActiveCourse$data;
    readonly " $fragmentRefs": FragmentRefs<"Module_myActiveCourse">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "Module_myActiveCourse",
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
(node as any).hash = 'bf3ef7c5a2886bdf0b29382ecf5f0c74';
export default node;
