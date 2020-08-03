/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CourseStatus = "complete" | "failed" | "incomplete" | "%future added value";
export type CourseType = "classroom" | "online" | "%future added value";
export type TrainingZone_user = {
    readonly firstName: string;
    readonly myCourses: ReadonlyArray<{
        readonly course: {
            readonly ident: number;
            readonly name: string;
            readonly excerpt: string | null;
            readonly category: {
                readonly color: string;
                readonly name: string;
            } | null;
            readonly type: CourseType;
            readonly bannerImageURL: string | null;
        };
        readonly status: CourseStatus;
        readonly enrolledAt: string;
        readonly progress: {
            readonly percent: number;
        } | null;
    }> | null;
    readonly " $refType": "TrainingZone_user";
};
export type TrainingZone_user$data = TrainingZone_user;
export type TrainingZone_user$key = {
    readonly " $data"?: TrainingZone_user$data;
    readonly " $fragmentRefs": FragmentRefs<"TrainingZone_user">;
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
        "name": "TrainingZone_user",
        "type": "User",
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
                "kind": "LinkedField",
                "alias": null,
                "name": "myCourses",
                "storageKey": null,
                "args": null,
                "concreteType": "MyCourse",
                "plural": true,
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
                                "kind": "LinkedField",
                                "alias": null,
                                "name": "category",
                                "storageKey": null,
                                "args": null,
                                "concreteType": "Category",
                                "plural": false,
                                "selections": [
                                    {
                                        "kind": "ScalarField",
                                        "alias": null,
                                        "name": "color",
                                        "args": null,
                                        "storageKey": null
                                    },
                                    (v0 /*: any*/)
                                ]
                            },
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
                                "name": "bannerImageURL",
                                "args": null,
                                "storageKey": null
                            }
                        ]
                    },
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
                        "name": "progress",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "Progress",
                        "plural": false,
                        "selections": [
                            {
                                "kind": "ScalarField",
                                "alias": null,
                                "name": "percent",
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
(node as any).hash = 'cfb6c5fc9cc21cde24212c8c8c585733';
export default node;
