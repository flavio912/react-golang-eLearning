/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CourseType = "classroom" | "online" | "%future added value";
export type OnlineCourses_user = {
    readonly firstName: string;
    readonly activeCourses: ReadonlyArray<{
        readonly course: {
            readonly name: string;
            readonly excerpt: string | null;
            readonly color: string | null;
            readonly type: CourseType;
            readonly ident: number;
        };
        readonly currentAttempt: number;
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
            "name": "activeCourses",
            "storageKey": null,
            "args": null,
            "concreteType": "ActiveCourse",
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
                            "alias": "ident",
                            "name": "id",
                            "args": null,
                            "storageKey": null
                        }
                    ]
                },
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "currentAttempt",
                    "args": null,
                    "storageKey": null
                }
            ]
        }
    ]
} as any);
(node as any).hash = '0550d2b2269ed130be1fa38a8d43c1cf';
export default node;
