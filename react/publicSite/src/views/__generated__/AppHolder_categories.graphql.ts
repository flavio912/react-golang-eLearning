/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type AppHolder_categories = {
    readonly edges: ReadonlyArray<{
        readonly uuid: string | null;
        readonly name: string;
        readonly color: string;
    } | null> | null;
    readonly pageInfo: {
        readonly total: number;
    } | null;
    readonly " $refType": "AppHolder_categories";
};
export type AppHolder_categories$data = AppHolder_categories;
export type AppHolder_categories$key = {
    readonly " $data"?: AppHolder_categories$data;
    readonly " $fragmentRefs": FragmentRefs<"AppHolder_categories">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "AppHolder_categories",
    "type": "CategoryPage",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
        {
            "kind": "LinkedField",
            "alias": null,
            "name": "edges",
            "storageKey": null,
            "args": null,
            "concreteType": "Category",
            "plural": true,
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
        },
        {
            "kind": "LinkedField",
            "alias": null,
            "name": "pageInfo",
            "storageKey": null,
            "args": null,
            "concreteType": "PageInfo",
            "plural": false,
            "selections": [
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "total",
                    "args": null,
                    "storageKey": null
                }
            ]
        }
    ]
} as any);
(node as any).hash = 'bc1f9d404304b42f934381eee0a17d35';
export default node;
