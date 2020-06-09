/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type DelegatesPage_manager = {
    readonly company: {
        readonly name: string;
    };
    readonly " $refType": "DelegatesPage_manager";
};
export type DelegatesPage_manager$data = DelegatesPage_manager;
export type DelegatesPage_manager$key = {
    readonly " $data"?: DelegatesPage_manager$data;
    readonly " $fragmentRefs": FragmentRefs<"DelegatesPage_manager">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "DelegatesPage_manager",
    "type": "Manager",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
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
                    "name": "name",
                    "args": null,
                    "storageKey": null
                }
            ]
        }
    ]
} as any);
(node as any).hash = '86b48e75faa9fdda7be7d5d4c7ef6181';
export default node;
