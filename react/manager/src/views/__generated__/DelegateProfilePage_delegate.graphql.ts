/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type DelegateProfilePage_delegate = {
    readonly firstName: string;
    readonly lastName: string;
    readonly " $refType": "DelegateProfilePage_delegate";
};
export type DelegateProfilePage_delegate$data = DelegateProfilePage_delegate;
export type DelegateProfilePage_delegate$key = {
    readonly " $data"?: DelegateProfilePage_delegate$data;
    readonly " $fragmentRefs": FragmentRefs<"DelegateProfilePage_delegate">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "DelegateProfilePage_delegate",
    "type": "Delegate",
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
        }
    ]
} as any);
(node as any).hash = 'a1517f51eb0cc04f9cddcfd61c349240';
export default node;
