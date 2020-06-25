/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type AppHolder_user = {
    readonly firstName: string;
    readonly lastName: string;
    readonly profileImageUrl: string | null;
    readonly " $refType": "AppHolder_user";
};
export type AppHolder_user$data = AppHolder_user;
export type AppHolder_user$key = {
    readonly " $data"?: AppHolder_user$data;
    readonly " $fragmentRefs": FragmentRefs<"AppHolder_user">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "AppHolder_user",
    "type": "User",
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
(node as any).hash = '11adf77ff9dfbd8bf5257acbc2b981f7';
export default node;