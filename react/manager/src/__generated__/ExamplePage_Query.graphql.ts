/* tslint:disable */
/* eslint-disable */
/* @relayHash d4035b6fd3545e62fe25c1db0a1b38fb */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type ExamplePage_QueryVariables = {};
export type ExamplePage_QueryResponse = {
    readonly " $fragmentRefs": FragmentRefs<"ExampleComponent_info">;
};
export type ExamplePage_Query = {
    readonly response: ExamplePage_QueryResponse;
    readonly variables: ExamplePage_QueryVariables;
};



/*
query ExamplePage_Query {
  ...ExampleComponent_info
}

fragment ExampleComponent_info on Query {
  info
}
*/

const node: ConcreteRequest = ({
    "kind": "Request",
    "fragment": {
        "kind": "Fragment",
        "name": "ExamplePage_Query",
        "type": "Query",
        "metadata": null,
        "argumentDefinitions": [],
        "selections": [
            {
                "kind": "FragmentSpread",
                "name": "ExampleComponent_info",
                "args": null
            }
        ]
    },
    "operation": {
        "kind": "Operation",
        "name": "ExamplePage_Query",
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
    },
    "params": {
        "operationKind": "query",
        "name": "ExamplePage_Query",
        "id": null,
        "text": "query ExamplePage_Query {\n  ...ExampleComponent_info\n}\n\nfragment ExampleComponent_info on Query {\n  info\n}\n",
        "metadata": {}
    }
} as any);
(node as any).hash = '52089d878ca3b08a3d9eddc69c2a7738';
export default node;
