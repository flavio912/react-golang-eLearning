/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type Module_module = {
    readonly name: string;
    readonly uuid: string;
    readonly voiceoverURL: string | null;
    readonly description: string;
    readonly transcript: string;
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
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "voiceoverURL",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "description",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "transcript",
            "args": null,
            "storageKey": null
        }
    ]
} as any);
(node as any).hash = '6a5f20da6dd5241ee1ab2092f56a76c8';
export default node;
