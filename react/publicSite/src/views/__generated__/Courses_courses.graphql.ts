/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type Courses_courses = {
    readonly edges: ReadonlyArray<{
        readonly ident: number;
        readonly name: string;
        readonly price: number;
        readonly excerpt: string | null;
        readonly introduction: string | null;
        readonly bannerImageURL: string | null;
    } | null> | null;
    readonly pageInfo: {
        readonly total: number;
        readonly offset: number;
        readonly limit: number;
        readonly given: number;
    } | null;
    readonly " $refType": "Courses_courses";
};
export type Courses_courses$data = Courses_courses;
export type Courses_courses$key = {
    readonly " $data"?: Courses_courses$data;
    readonly " $fragmentRefs": FragmentRefs<"Courses_courses">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "Courses_courses",
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
                    "name": "price",
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
                    "name": "bannerImageURL",
                    "args": null,
                    "storageKey": null
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
} as any);
(node as any).hash = '1ed3be92146747e1179fd667a3eb89f4';
export default node;
