/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CourseStatus = "complete" | "failed" | "incomplete" | "%future added value";
export type DelegateProfilePage_delegate = {
    readonly firstName: string;
    readonly lastName: string;
    readonly myCourses: ReadonlyArray<{
        readonly status: CourseStatus;
        readonly course: {
            readonly name: string;
            readonly category: {
                readonly name: string;
            } | null;
        };
    }> | null;
    readonly " $refType": "DelegateProfilePage_delegate";
};
export type DelegateProfilePage_delegate$data = DelegateProfilePage_delegate;
export type DelegateProfilePage_delegate$key = {
    readonly " $data"?: DelegateProfilePage_delegate$data;
    readonly " $fragmentRefs": FragmentRefs<"DelegateProfilePage_delegate">;
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
        "name": "DelegateProfilePage_delegate",
        "type": "Delegate",
        "metadata": null,
        "argumentDefinitions": [],
        "selections": [
            {
                "kind": "ScalarField",
                "alias": null,
                "name": "firstName",
                "args": null,
                "storageKey": null
            },
            {
                "kind": "ScalarField",
                "alias": null,
                "name": "lastName",
                "args": null,
                "storageKey": null
            },
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "myCourses",
                "storageKey": null,
                "args": null,
                "concreteType": "MyCourse",
                "plural": true,
                "selections": [
                    {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "status",
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
                            (v0 /*: any*/),
                            {
                                "kind": "LinkedField",
                                "alias": null,
                                "name": "category",
                                "storageKey": null,
                                "args": null,
                                "concreteType": "Category",
                                "plural": false,
                                "selections": [
                                    (v0 /*: any*/)
                                ]
                            }
                        ]
                    }
                ]
            }
        ]
    } as any;
})();
(node as any).hash = 'd85d7e60751bc1864497e64c71994411';
export default node;
