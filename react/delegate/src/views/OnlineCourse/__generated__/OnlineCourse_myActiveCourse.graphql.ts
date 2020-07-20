/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CourseStatus = "failed" | "incomplete" | "passed" | "%future added value";
export type StructureElement = "lesson" | "module" | "test" | "%future added value";
export type OnlineCourse_myActiveCourse = {
    readonly status: CourseStatus;
    readonly enrolledAt: string;
    readonly upTo: string | null;
    readonly course: {
        readonly ident: number;
        readonly name: string;
        readonly excerpt: string | null;
        readonly introduction: string | null;
        readonly howToComplete: string | null;
        readonly whatYouLearn: ReadonlyArray<string> | null;
        readonly hoursToComplete: number | null;
        readonly syllabus: ReadonlyArray<{
            readonly name: string;
            readonly type: StructureElement;
            readonly uuid: string;
        }> | null;
        readonly " $fragmentRefs": FragmentRefs<"CourseSyllabusCardFrag_course">;
    };
    readonly " $refType": "OnlineCourse_myActiveCourse";
};
export type OnlineCourse_myActiveCourse$data = OnlineCourse_myActiveCourse;
export type OnlineCourse_myActiveCourse$key = {
    readonly " $data"?: OnlineCourse_myActiveCourse$data;
    readonly " $fragmentRefs": FragmentRefs<"OnlineCourse_myActiveCourse">;
};



const node: ReaderFragment = (function () {
    var v0 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "name",
        "args": null,
        "storageKey": null
    } as any);
    return {
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
                "kind": "ScalarField",
                "alias": null,
                "name": "upTo",
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
                        "alias": "ident",
                        "name": "id",
                        "args": null,
                        "storageKey": null
                    },
                    (v0 /*: any*/),
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
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "syllabus",
                        "storageKey": null,
                        "args": null,
                        "concreteType": null,
                        "plural": true,
                        "selections": [
                            (v0 /*: any*/),
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "type",
                                "args": null,
                                "storageKey": null
                            },
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "uuid",
                                "args": null,
                                "storageKey": null
                            }
                        ]
                    },
                    {
                        "kind": "FragmentSpread",
                        "name": "CourseSyllabusCardFrag_course",
                        "args": null
                    }
                ]
            }
        ]
    } as any;
})();
(node as any).hash = '301698d3cb50c59c5742acc4c4ba50a2';
export default node;
