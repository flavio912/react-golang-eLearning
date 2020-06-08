/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type AppHolder_manager = {
    readonly firstName: string;
    readonly lastName: string;
    readonly profileImageUrl: string | null;
    readonly " $refType": "AppHolder_manager";
};
export type AppHolder_manager$data = AppHolder_manager;
export type AppHolder_manager$key = {
    readonly " $data"?: AppHolder_manager$data;
    readonly " $fragmentRefs": FragmentRefs<"AppHolder_manager">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "AppHolder_manager",
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
            "name": "profileImageUrl",
            "args": null,
            "storageKey": null
        }
    ]
} as any);
(node as any).hash = 'ec59545289870c174052616b7c79b5d3';
export default node;
