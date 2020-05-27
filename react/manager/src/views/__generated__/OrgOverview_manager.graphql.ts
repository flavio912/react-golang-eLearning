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
        }
    ]
} as any);
(node as any).hash = '12e21416a68d12eccb5d1cc3f1f45df2';
export default node;
