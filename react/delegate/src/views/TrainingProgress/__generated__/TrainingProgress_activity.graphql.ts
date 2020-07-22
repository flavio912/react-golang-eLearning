/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type TrainingProgress_activity = {
    readonly " $fragmentRefs": FragmentRefs<"ActivityTable_activity">;
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
            "kind": "FragmentSpread",
            "name": "ActivityTable_activity",
            "args": null
        }
    ]
} as any);
(node as any).hash = 'e13ff71d1b9d887aaaa01a4289fed4da';
export default node;
