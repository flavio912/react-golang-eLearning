/* tslint:disable */
/* eslint-disable */
/* @relayHash b4b01656357e4804ff9e509bcda2d9bf */

import { ConcreteRequest } from "relay-runtime";
export type AppHolderQueryVariables = {
    name: string;
};
export type AppHolderQueryResponse = {
    readonly courses: {
        readonly edges: ReadonlyArray<{
            readonly notId: number;
            readonly name: string;
            readonly bannerImageURL: string | null;
            readonly introduction: string | null;
        } | null> | null;
    } | null;
};
export type AppHolderQuery = {
    readonly response: AppHolderQueryResponse;
    readonly variables: AppHolderQueryVariables;
};



/*
query AppHolderQuery(
  $name: String!
) {
  courses(filter: {name: $name}, page: {limit: 4}) {
    edges {
      notId: id
      name
      bannerImageURL
      introduction
    }
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "name",
            "type": "String!",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "LinkedField",
            "alias": null,
            "name": "courses",
            "storageKey": null,
            "args": [
                {
                    "kind": "ObjectValue",
                    "name": "filter",
                    "fields": [
                        {
                            "kind": "Variable",
                            "name": "name",
                            "variableName": "name"
                        }
                    ]
                },
                {
                    "kind": "Literal",
                    "name": "page",
                    "value": {
                        "limit": 4
                    }
                }
            ],
            "concreteType": "CoursePage",
            "plural": false,
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "edges",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Course",
                    "plural": true,
                    "selections": [
                        {
                            "kind": "ScalarField",
                            "alias": "notId",
                            "name": "id",
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
                            "name": "bannerImageURL",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "introduction",
                            "args": null,
                            "storageKey": null
                        }
                    ]
                }
            ]
        } as any)
    ];
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "AppHolderQuery",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "AppHolderQuery",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "query",
            "name": "AppHolderQuery",
            "id": null,
            "text": "query AppHolderQuery(\n  $name: String!\n) {\n  courses(filter: {name: $name}, page: {limit: 4}) {\n    edges {\n      notId: id\n      name\n      bannerImageURL\n      introduction\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '9f10a3c6d550cc3712f1857c1b9242f3';
export default node;
