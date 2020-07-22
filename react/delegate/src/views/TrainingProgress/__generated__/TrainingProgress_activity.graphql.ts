/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type ActivityType = "activated" | "completedCourse" | "failedCourse" | "newCourse" | "%future added value";
export type TrainingProgress_activity = {
    readonly edges: ReadonlyArray<{
        readonly type: ActivityType;
        readonly createdAt: string;
        readonly course: {
            readonly name: string;
        } | null;
    } | null> | null;
    readonly pageInfo: {
        readonly total: number;
    } | null;
    readonly " $refType": "TrainingProgress_activity";
};
export type TrainingProgress_activity$data = TrainingProgress_activity;
export type TrainingProgress_activity$key = {
    readonly " $data"?: TrainingProgress_activity$data;
    readonly " $fragmentRefs": FragmentRefs<"TrainingProgress_activity">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "TrainingProgress_activity",
    "type": "ActivityPage",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
        {
            "kind": "LinkedField",
            "alias": null,
            "name": "edges",
            "storageKey": null,
            "args": null,
            "concreteType": "Activity",
            "plural": true,
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
                }
            ]
        }
    ]
} as any);
(node as any).hash = 'c9d4f31dbfe952fe3623ce18b9b54398';
export default node;
