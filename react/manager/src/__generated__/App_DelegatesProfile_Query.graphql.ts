/* tslint:disable */
/* eslint-disable */
/* @relayHash cf5d02d75554a3684e45b339f756e1e4 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_DelegatesProfile_QueryVariables = {
    uuid: string;
};
export type App_DelegatesProfile_QueryResponse = {
    readonly delegate: {
        readonly " $fragmentRefs": FragmentRefs<"DelegateProfilePage_delegate">;
    } | null;
};
export type App_DelegatesProfile_Query = {
    readonly response: App_DelegatesProfile_QueryResponse;
    readonly variables: App_DelegatesProfile_QueryVariables;
};



/*
query App_DelegatesProfile_Query(
  $uuid: UUID!
) {
  delegate(uuid: $uuid) {
    ...DelegateProfilePage_delegate
  }
}

fragment DelegateProfilePage_delegate on Delegate {
  firstName
  lastName
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "uuid",
            "type": "UUID!",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "Variable",
            "name": "uuid",
            "variableName": "uuid"
        } as any)
    ];
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "App_DelegatesProfile_Query",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "delegate",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
                    "concreteType": "Delegate",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "FragmentSpread",
                            "name": "DelegateProfilePage_delegate",
                            "args": null
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "App_DelegatesProfile_Query",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "delegate",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
                    "concreteType": "Delegate",
                    "plural": false,
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
                        }
                    ]
                }
            ]
        },
        "params": {
            "operationKind": "query",
            "name": "App_DelegatesProfile_Query",
            "id": null,
            "text": "query App_DelegatesProfile_Query(\n  $uuid: UUID!\n) {\n  delegate(uuid: $uuid) {\n    ...DelegateProfilePage_delegate\n  }\n}\n\nfragment DelegateProfilePage_delegate on Delegate {\n  firstName\n  lastName\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '8d71d48492ccccdeb87bd3c6e6ee8ae7';
export default node;
