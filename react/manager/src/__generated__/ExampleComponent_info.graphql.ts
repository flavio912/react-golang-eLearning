/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type ExampleComponent_info = {
    readonly info: string;
    readonly " $refType": "ExampleComponent_info";
};
export type ExampleComponent_info$data = ExampleComponent_info;
export type ExampleComponent_info$key = {
    readonly " $data"?: ExampleComponent_info$data;
    readonly " $fragmentRefs": FragmentRefs<"ExampleComponent_info">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "ExampleComponent_info",
    "type": "Query",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "info",
            "args": null,
            "storageKey": null
        }
    ]
} as any);
(node as any).hash = '8b6debfd46d67ae45d2fe233bc8a18b7';
export default node;
