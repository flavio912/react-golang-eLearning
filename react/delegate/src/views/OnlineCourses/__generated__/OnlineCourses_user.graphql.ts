/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CourseStatus = "complete" | "failed" | "incomplete" | "%future added value";
export type CourseType = "classroom" | "online" | "%future added value";
export type OnlineCourses_user = {
    readonly firstName: string;
    readonly myCourses: ReadonlyArray<{
        readonly course: {
            readonly ident: number;
            readonly name: string;
            readonly excerpt: string | null;
            readonly color: string | null;
            readonly type: CourseType;
            readonly bannerImageURL: string | null;
        };
        readonly status: CourseStatus;
    }> | null;
    readonly " $refType": "OnlineCourses_user";
};
export type OnlineCourses_user$data = OnlineCourses_user;
export type OnlineCourses_user$key = {
    readonly " $data"?: OnlineCourses_user$data;
    readonly " $fragmentRefs": FragmentRefs<"OnlineCourses_user">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "OnlineCourses_user",
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
                            "name": "excerpt",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "color",
                            "args": null,
                            "storageKey": null
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
                }
            ]
        }
    ]
} as any);
(node as any).hash = '482aaec30f8d843aac4de3dc1343303b';
export default node;
