/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CoursesPage_onlineCourses = {
    readonly edges: ReadonlyArray<{
        readonly uuid: unknown;
        readonly info: {
            readonly name: string | null;
            readonly color: string | null;
            readonly excerpt: string | null;
            readonly price: number | null;
            readonly category: {
                readonly name: string;
                readonly color: string;
            } | null;
        };
    } | null>;
    readonly pageInfo: {
        readonly total: number;
        readonly offset: number;
        readonly limit: number;
        readonly given: number;
    } | null;
    readonly " $refType": "CoursesPage_onlineCourses";
};
export type CoursesPage_onlineCourses$data = CoursesPage_onlineCourses;
export type CoursesPage_onlineCourses$key = {
    readonly " $data"?: CoursesPage_onlineCourses$data;
    readonly " $fragmentRefs": FragmentRefs<"CoursesPage_onlineCourses">;
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
        "name": "color",
        "args": null,
        "storageKey": null
    } as any);
    return {
        "kind": "Fragment",
        "name": "CoursesPage_onlineCourses",
        "type": "OnlineCoursePage",
        "metadata": null,
        "argumentDefinitions": [],
        "selections": [
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "edges",
                "storageKey": null,
                "args": null,
                "concreteType": "OnlineCourse",
                "plural": true,
                "selections": [
                    {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "uuid",
                        "args": null,
                        "storageKey": null
                    },
                    {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "info",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "CourseInfo",
                        "plural": false,
                        "selections": [
                            (v0 /*: any*/),
                            (v1 /*: any*/),
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
                                "name": "price",
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
                                    (v0 /*: any*/),
                                    (v1 /*: any*/)
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "pageInfo",
                "storageKey": null,
                "args": null,
                "concreteType": "PageInfo",
                "plural": false,
                "selections": [
                    {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "total",
                        "args": null,
                        "storageKey": null
                    },
                    {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "offset",
                        "args": null,
                        "storageKey": null
                    },
                    {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "limit",
                        "args": null,
                        "storageKey": null
                    },
                    {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "given",
                        "args": null,
                        "storageKey": null
                    }
                ]
            }
        ]
    } as any;
})();
(node as any).hash = 'e364f655046380720f4f8f5e971363a8';
export default node;
