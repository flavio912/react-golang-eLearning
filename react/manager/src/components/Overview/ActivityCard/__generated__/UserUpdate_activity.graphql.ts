/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type ActivityType = "activated" | "completedCourse" | "failedCourse" | "newCourse" | "%future added value";
export type UserUpdate_activity = {
    readonly type: ActivityType;
    readonly createdAt: string;
    readonly course: {
        readonly name: string;
    } | null;
    readonly user: {
        readonly firstName: string;
        readonly lastName: string;
        readonly profileImageUrl: string | null;
    } | null;
    readonly " $refType": "UserUpdate_activity";
};
export type UserUpdate_activity$data = UserUpdate_activity;
export type UserUpdate_activity$key = {
    readonly " $data"?: UserUpdate_activity$data;
    readonly " $fragmentRefs": FragmentRefs<"UserUpdate_activity">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "UserUpdate_activity",
    "type": "Activity",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
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
            "name": "createdAt",
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
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "name",
                    "args": null,
                    "storageKey": null
                }
            ]
        },
        {
            "kind": "LinkedField",
            "alias": null,
            "name": "user",
            "storageKey": null,
            "args": null,
            "concreteType": "User",
            "plural": false,
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
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "profileImageUrl",
                    "args": null,
                    "storageKey": null
                }
            ]
        }
    ]
} as any);
(node as any).hash = 'ced1d47dcfe89fc6467a96bd9614befe';
export default node;
