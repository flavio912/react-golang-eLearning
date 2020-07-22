/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type StructureElement = "lesson" | "module" | "test" | "%future added value";
export type CourseSyllabusCardFrag_course = {
    readonly syllabus: ReadonlyArray<{
        readonly name: string;
        readonly type: StructureElement;
        readonly uuid: string;
        readonly complete: boolean | null;
        readonly syllabus?: ReadonlyArray<{
            readonly name: string;
            readonly type: StructureElement;
            readonly uuid: string;
            readonly complete: boolean | null;
        }> | null;
    }> | null;
    readonly " $refType": "CourseSyllabusCardFrag_course";
};
export type CourseSyllabusCardFrag_course$data = CourseSyllabusCardFrag_course;
export type CourseSyllabusCardFrag_course$key = {
    readonly " $data"?: CourseSyllabusCardFrag_course$data;
    readonly " $fragmentRefs": FragmentRefs<"CourseSyllabusCardFrag_course">;
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
    } as any), v3 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "complete",
        "args": null,
        "storageKey": null
    } as any);
    return {
        "kind": "Fragment",
        "name": "CourseSyllabusCardFrag_course",
        "type": "Course",
        "metadata": null,
        "argumentDefinitions": [],
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
                    (v3 /*: any*/),
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
                                    (v1 /*: any*/),
                                    (v2 /*: any*/),
                                    (v3 /*: any*/)
                                ]
                            }
                        ]
                    }
                ]
            }
        ]
    } as any;
})();
(node as any).hash = '8bdb89b56cd9cec16c31316d894281f7';
export default node;
