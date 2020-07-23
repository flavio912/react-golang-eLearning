/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type AppHolder_courses = {
    readonly edges: ReadonlyArray<{
        readonly notId: number;
        readonly name: string;
        readonly bannerImageURL: string | null;
        readonly introduction: string | null;
    } | null> | null;
    readonly pageInfo: {
        readonly total: number;
        readonly offset: number;
        readonly limit: number;
        readonly given: number;
    } | null;
    readonly " $refType": "AppHolder_courses";
};
export type AppHolder_courses$data = AppHolder_courses;
export type AppHolder_courses$key = {
    readonly " $data"?: AppHolder_courses$data;
    readonly " $fragmentRefs": FragmentRefs<"AppHolder_courses">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "AppHolder_courses",
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
                    "alias": "notId",
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
                    "name": "bannerImageURL",
                    "args": null,
                    "storageKey": null
                },
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "introduction",
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
(node as any).hash = 'f589a60a9d6ddc6b716bfc026b24bde5';
export default node;
