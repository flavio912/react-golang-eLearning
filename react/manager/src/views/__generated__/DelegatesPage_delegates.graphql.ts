/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type DelegatesPage_delegates = {
    readonly edges: ReadonlyArray<{
        readonly uuid: string;
        readonly email: string | null;
        readonly firstName: string;
        readonly lastName: string;
        readonly lastLogin: string;
        readonly createdAt: string | null;
    } | null> | null;
    readonly pageInfo: {
        readonly total: number;
        readonly offset: number;
        readonly limit: number;
        readonly given: number;
    } | null;
    readonly " $refType": "DelegatesPage_delegates";
};
export type DelegatesPage_delegates$data = DelegatesPage_delegates;
export type DelegatesPage_delegates$key = {
    readonly " $data"?: DelegatesPage_delegates$data;
    readonly " $fragmentRefs": FragmentRefs<"DelegatesPage_delegates">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "DelegatesPage_delegates",
    "type": "DelegatePage",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
        {
            "kind": "LinkedField",
            "alias": null,
            "name": "edges",
            "storageKey": null,
            "args": null,
            "concreteType": "Delegate",
            "plural": true,
            "selections": [
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "uuid",
                    "args": null,
                    "storageKey": null
                },
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "email",
                    "args": null,
                    "storageKey": null
                },
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
                    "name": "lastLogin",
                    "args": null,
                    "storageKey": null
                },
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "createdAt",
                    "args": null,
                    "storageKey": null
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
                    "name": "offset",
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
                    "name": "given",
                    "args": null,
                    "storageKey": null
                }
            ]
        }
    ]
} as any);
(node as any).hash = '8a7f29cca036c34433d36538addac5f0';
export default node;
