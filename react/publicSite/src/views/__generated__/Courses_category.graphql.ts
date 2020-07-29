/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type Courses_category = {
    readonly uuid: string | null;
    readonly name: string;
    readonly color: string;
    readonly " $refType": "Courses_category";
};
export type Courses_category$data = Courses_category;
export type Courses_category$key = {
    readonly " $data"?: Courses_category$data;
    readonly " $fragmentRefs": FragmentRefs<"Courses_category">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "Courses_category",
    "type": "Category",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
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
            "name": "name",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "color",
            "args": null,
            "storageKey": null
        }
    ]
} as any);
(node as any).hash = 'ab10a98627f9574a3d1cd7125c4cb98b';
export default node;
