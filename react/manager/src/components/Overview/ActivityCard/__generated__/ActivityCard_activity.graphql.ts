/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type ActivityCard_activity = {
    readonly edges: ReadonlyArray<{
        readonly " $fragmentRefs": FragmentRefs<"UserUpdate_activity">;
    } | null> | null;
    readonly " $refType": "ActivityCard_activity";
};
export type ActivityCard_activity$data = ActivityCard_activity;
export type ActivityCard_activity$key = {
    readonly " $data"?: ActivityCard_activity$data;
    readonly " $fragmentRefs": FragmentRefs<"ActivityCard_activity">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "ActivityCard_activity",
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
                    "kind": "FragmentSpread",
                    "name": "UserUpdate_activity",
                    "args": null
                }
            ]
        }
    ]
} as any);
(node as any).hash = 'edefe2c1870374eb530f70240bcc72fe';
export default node;
