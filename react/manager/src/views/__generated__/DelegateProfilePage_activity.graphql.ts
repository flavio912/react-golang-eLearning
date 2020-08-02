/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type DelegateProfilePage_activity = {
    readonly " $fragmentRefs": FragmentRefs<"ActivityTable_activity">;
    readonly " $refType": "DelegateProfilePage_activity";
};
export type DelegateProfilePage_activity$data = DelegateProfilePage_activity;
export type DelegateProfilePage_activity$key = {
    readonly " $data"?: DelegateProfilePage_activity$data;
    readonly " $fragmentRefs": FragmentRefs<"DelegateProfilePage_activity">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "DelegateProfilePage_activity",
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
(node as any).hash = '1d987834673f0c45ef0b4307c86ebaa4';
export default node;
