/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CoursesPage_courses = {
    readonly edges: ReadonlyArray<{
        readonly ident: number;
        readonly name: string;
        readonly color: string | null;
        readonly excerpt: string | null;
        readonly price: number;
        readonly category: {
            readonly name: string;
            readonly color: string;
        } | null;
    } | null> | null;
    readonly pageInfo: {
        readonly total: number;
        readonly offset: number;
        readonly limit: number;
        readonly given: number;
    } | null;
    readonly " $refType": "CoursesPage_courses";
};
export type CoursesPage_courses$data = CoursesPage_courses;
export type CoursesPage_courses$key = {
    readonly " $data"?: CoursesPage_courses$data;
    readonly " $fragmentRefs": FragmentRefs<"CoursesPage_courses">;
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
        "name": "CoursesPage_courses",
        "type": "CoursePage",
        "metadata": null,
        "argumentDefinitions": [],
        "selections": [
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "edges",
                "storageKey": null,
                "args": null,
                "concreteType": "Course",
                "plural": true,
                "selections": [
                    {
                        "kind": "ScalarField",
                        "alias": "ident",
                        "name": "id",
                        "args": null,
                        "storageKey": null
                    },
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
(node as any).hash = '172bd14a3c1370079363ff4a71e47496';
export default node;
