/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CourseStatus = "failed" | "incomplete" | "passed" | "%future added value";
export type OnlineCourse_myActiveCourse = {
    readonly status: CourseStatus;
    readonly enrolledAt: string;
    readonly course: {
        readonly name: string;
        readonly excerpt: string | null;
        readonly introduction: string | null;
        readonly howToComplete: string | null;
        readonly whatYouLearn: ReadonlyArray<string> | null;
        readonly hoursToComplete: number | null;
        readonly " $fragmentRefs": FragmentRefs<"CourseSyllabusCardFrag_course">;
    };
    readonly " $refType": "OnlineCourse_myActiveCourse";
};
export type OnlineCourse_myActiveCourse$data = OnlineCourse_myActiveCourse;
export type OnlineCourse_myActiveCourse$key = {
    readonly " $data"?: OnlineCourse_myActiveCourse$data;
    readonly " $fragmentRefs": FragmentRefs<"OnlineCourse_myActiveCourse">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "OnlineCourse_myActiveCourse",
    "type": "MyCourse",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "status",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "enrolledAt",
            "args": null,
            "storageKey": null
        },
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
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "name",
                    "args": null,
                    "storageKey": null
                },
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "excerpt",
                    "args": null,
                    "storageKey": null
                },
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "introduction",
                    "args": null,
                    "storageKey": null
                },
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "howToComplete",
                    "args": null,
                    "storageKey": null
                },
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "whatYouLearn",
                    "args": null,
                    "storageKey": null
                },
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "hoursToComplete",
                    "args": null,
                    "storageKey": null
                },
                {
                    "kind": "FragmentSpread",
                    "name": "CourseSyllabusCardFrag_course",
                    "args": null
                }
            ]
        }
    ]
} as any);
(node as any).hash = '5836f94c222aec11b34ee4b8596616ee';
export default node;
