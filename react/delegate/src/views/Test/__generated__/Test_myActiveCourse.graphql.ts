/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type StructureElement = "lesson" | "module" | "test" | "%future added value";
export type Test_myActiveCourse = {
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
    readonly " $refType": "Test_myActiveCourse";
};
export type Test_myActiveCourse$data = Test_myActiveCourse;
export type Test_myActiveCourse$key = {
    readonly " $data"?: Test_myActiveCourse$data;
    readonly " $fragmentRefs": FragmentRefs<"Test_myActiveCourse">;
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
            }
        ]
    } as any;
})();
(node as any).hash = '75c4e7c76f0659ec495519ba6f3bd1b9';
export default node;
