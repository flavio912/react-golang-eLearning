/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type AppHolder_manager = {
    readonly firstName: string;
    readonly lastName: string;
    readonly profileImageUrl: string | null;
    readonly company: {
        readonly logoURL: string | null;
    };
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
                    "name": "logoURL",
                    "args": null,
                    "storageKey": null
                }
            ]
        }
    ]
} as any);
(node as any).hash = 'ad6c5177293ee421311268539136ff1a';
export default node;
