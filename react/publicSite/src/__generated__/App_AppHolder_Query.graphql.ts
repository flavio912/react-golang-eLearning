/* tslint:disable */
/* eslint-disable */
/* @relayHash 639c81ac28c4e3e1b61997b05f330add */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_AppHolder_QueryVariables = {};
export type App_AppHolder_QueryResponse = {
    readonly categories: {
        readonly " $fragmentRefs": FragmentRefs<"AppHolder_categories">;
    } | null;
};
export type App_AppHolder_Query = {
    readonly response: App_AppHolder_QueryResponse;
    readonly variables: App_AppHolder_QueryVariables;
};



/*
query App_AppHolder_Query {
  categories {
    ...AppHolder_categories
  }
}

fragment AppHolder_categories on CategoryPage {
  edges {
    uuid
    name
    color
  }
  pageInfo {
    total
  }
}
*/

const node: ConcreteRequest = ({
    "kind": "Request",
    "fragment": {
        "kind": "Fragment",
        "name": "App_AppHolder_Query",
        "type": "Query",
        "metadata": null,
        "argumentDefinitions": [],
        "selections": [
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "categories",
                "storageKey": null,
                "args": null,
                "concreteType": "CategoryPage",
                "plural": false,
                "selections": [
                    {
                        "kind": "FragmentSpread",
                        "name": "AppHolder_categories",
                        "args": null
                    }
                ]
            }
        ]
    },
    "operation": {
        "kind": "Operation",
        "name": "App_AppHolder_Query",
        "argumentDefinitions": [],
        "selections": [
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "categories",
                "storageKey": null,
                "args": null,
                "concreteType": "CategoryPage",
                "plural": false,
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
            }
        ]
    },
    "params": {
        "operationKind": "query",
        "name": "App_AppHolder_Query",
        "id": null,
        "text": "query App_AppHolder_Query {\n  categories {\n    ...AppHolder_categories\n  }\n}\n\nfragment AppHolder_categories on CategoryPage {\n  edges {\n    uuid\n    name\n    color\n  }\n  pageInfo {\n    total\n  }\n}\n",
        "metadata": {}
    }
} as any);
(node as any).hash = '483d604b9e9be2b467f41358b9ad188f';
export default node;
