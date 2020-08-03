/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type StructureElement = "lesson" | "module" | "test" | "%future added value";
export type Lesson_myActiveCourse = {
    readonly course: {
        readonly syllabus: ReadonlyArray<{
            readonly name: string;
            readonly type: StructureElement;
            readonly uuid: string;
            readonly syllabus?: ReadonlyArray<{
                readonly name: string;
                readonly uuid: string;
                readonly type: StructureElement;
            }> | null;
        }> | null;
        readonly " $fragmentRefs": FragmentRefs<"CourseSyllabusCardFrag_course">;
    };
    readonly upTo: string | null;
    readonly " $refType": "Lesson_myActiveCourse";
};
export type Lesson_myActiveCourse$data = Lesson_myActiveCourse;
export type Lesson_myActiveCourse$key = {
    readonly " $data"?: Lesson_myActiveCourse$data;
    readonly " $fragmentRefs": FragmentRefs<"Lesson_myActiveCourse">;
};



const node: ReaderFragment = (function () {
    var v0 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "name",
        "args": null,
        "storageKey": null
    } as any), v1 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "type",
        "args": null,
        "storageKey": null
    } as any), v2 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "uuid",
        "args": null,
        "storageKey": null
    } as any);
    return {
        "kind": "Fragment",
        "name": "Lesson_myActiveCourse",
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
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "syllabus",
                        "storageKey": null,
                        "args": null,
                        "concreteType": null,
                        "plural": true,
                        "selections": [
                            (v0 /*: any*/),
                            (v1 /*: any*/),
                            (v2 /*: any*/),
                            {
                                "kind": "InlineFragment",
                                "type": "Module",
                                "selections": [
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
                                            (v2 /*: any*/),
                                            (v1 /*: any*/)
                                        ]
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "kind": "FragmentSpread",
                        "name": "CourseSyllabusCardFrag_course",
                        "args": null
                    }
                ]
            },
            {
                "kind": "ScalarField",
                "alias": null,
                "name": "upTo",
                "args": null,
                "storageKey": null
            }
        ]
    } as any;
})();
(node as any).hash = '95454b3bea1ea117782531c05485051f';
export default node;