/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type UserType = "delegate" | "individual" | "manager" | "%future added value";
export type AppHolder_user = {
    readonly type: UserType;
    readonly firstName: string;
    readonly lastName: string;
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
            "name": "type",
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
        }
    ]
} as any);
(node as any).hash = '6b5a69f30315f2d9cebccd9277859fec';
export default node;
