/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type Module_module = {
    readonly name: string;
    readonly uuid: string;
    readonly " $refType": "Module_module";
};
export type Module_module$data = Module_module;
export type Module_module$key = {
    readonly " $data"?: Module_module$data;
    readonly " $fragmentRefs": FragmentRefs<"Module_module">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "Module_module",
    "type": "Module",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "name",
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
(node as any).hash = 'f44909eaa4c3c9b89669bcec523b082e';
export default node;
