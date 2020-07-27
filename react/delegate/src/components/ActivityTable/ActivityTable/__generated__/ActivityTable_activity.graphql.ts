/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type ActivityType = "activated" | "completedCourse" | "failedCourse" | "newCourse" | "%future added value";
export type ActivityTable_activity = {
    readonly edges: ReadonlyArray<{
        readonly type: ActivityType;
        readonly createdAt: string;
        readonly course: {
            readonly ident: number;
            readonly name: string;
        } | null;
    } | null> | null;
    readonly pageInfo: {
        readonly total: number;
        readonly limit: number;
        readonly offset: number;
    } | null;
    readonly " $refType": "ActivityTable_activity";
};
export type ActivityTable_activity$data = ActivityTable_activity;
export type ActivityTable_activity$key = {
    readonly " $data"?: ActivityTable_activity$data;
    readonly " $fragmentRefs": FragmentRefs<"ActivityTable_activity">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "ActivityTable_activity",
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
                    "name": "offset",
                    "args": null,
                    "storageKey": null
                }
            ]
        }
    ]
} as any);
(node as any).hash = '0834b53af77015cdec392c6a08dca6f0';
export default node;
