/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CourseStatus = "complete" | "failed" | "incomplete" | "%future added value";
export type TrainingZone_user = {
    readonly firstName: string;
    readonly myCourses: ReadonlyArray<{
        readonly status: CourseStatus;
        readonly enrolledAt: string;
    }> | null;
    readonly " $refType": "TrainingZone_user";
};
export type TrainingZone_user$data = TrainingZone_user;
export type TrainingZone_user$key = {
    readonly " $data"?: TrainingZone_user$data;
    readonly " $fragmentRefs": FragmentRefs<"TrainingZone_user">;
};



const node: ReaderFragment = ({
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
                }
            ]
        }
    ]
} as any);
(node as any).hash = '6c8ec02fd709f6a57443b0bc50c0bf9b';
export default node;
