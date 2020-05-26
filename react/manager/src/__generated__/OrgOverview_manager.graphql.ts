/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type OrgOverview_manager = {
    readonly firstName: string;
    readonly lastName: string;
    readonly uuid: unknown;
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
            "name": "uuid",
            "args": null,
            "storageKey": null
        }
    ]
} as any);
(node as any).hash = 'f1de603aa606898b1b9e80fb88159bbc';
export default node;
