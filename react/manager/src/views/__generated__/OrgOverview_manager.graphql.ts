/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type OrgOverview_manager = {
    readonly firstName: string;
    readonly lastName: string;
    readonly email: string;
    readonly telephone: string;
    readonly createdAt: string | null;
    readonly company: {
        readonly name: string;
        readonly delegates: {
            readonly pageInfo: {
                readonly total: number;
            } | null;
        };
    };
    readonly " $refType": "OrgOverview_manager";
};
export type OrgOverview_manager$data = OrgOverview_manager;
export type OrgOverview_manager$key = {
    readonly " $data"?: OrgOverview_manager$data;
    readonly " $fragmentRefs": FragmentRefs<"OrgOverview_manager">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "OrgOverview_manager",
    "type": "Manager",
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
            "kind": "ScalarField",
            "alias": null,
            "name": "lastName",
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
            "name": "telephone",
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
            "name": "company",
            "storageKey": null,
            "args": null,
            "concreteType": "Company",
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
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "delegates",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "DelegatePage",
                    "plural": false,
                    "selections": [
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
                }
            ]
        }
    ]
} as any);
(node as any).hash = '2b05b70b64616f7cdb49e3058abce406';
export default node;
