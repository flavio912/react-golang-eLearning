/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CourseDetailsPage_course = {
    readonly ident: number;
    readonly name: string;
    readonly price: number;
    readonly excerpt: string | null;
    readonly introduction: string | null;
    readonly bannerImageURL: string | null;
    readonly howToComplete: string | null;
    readonly hoursToComplete: number | null;
    readonly whatYouLearn: ReadonlyArray<string> | null;
    readonly requirements: ReadonlyArray<string> | null;
    readonly category: {
        readonly name: string;
        readonly color: string;
    } | null;
    readonly " $refType": "CourseDetailsPage_course";
};
export type CourseDetailsPage_course$data = CourseDetailsPage_course;
export type CourseDetailsPage_course$key = {
    readonly " $data"?: CourseDetailsPage_course$data;
    readonly " $fragmentRefs": FragmentRefs<"CourseDetailsPage_course">;
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
        "name": "CourseDetailsPage_course",
        "type": "Course",
        "metadata": null,
        "argumentDefinitions": [],
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
                "name": "hoursToComplete",
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
                "name": "requirements",
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
                    {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "color",
                        "args": null,
                        "storageKey": null
                    }
                ]
            }
        ]
    } as any;
})();
(node as any).hash = '67313fa1af64b0283da08d2a7187006f';
export default node;
