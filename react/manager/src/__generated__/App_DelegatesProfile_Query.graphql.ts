/* tslint:disable */
/* eslint-disable */
/* @relayHash 7411ad3eb2942b46845d555835c34cde */

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
  myCourses {
    status
    course {
      name
      category {
        name
      }
    }
  }
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
    ], v2 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "name",
        "args": null,
        "storageKey": null
    } as any);
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
                        },
                        {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "myCourses",
                            "storageKey": null,
                            "args": null,
                            "concreteType": "MyCourse",
                            "plural": true,
                            "selections": [
                                {
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "status",
                                    "args": null,
                                    "storageKey": null
                                },
                                {
                                    "kind": "LinkedField",
                                    "alias": null,
                                    "name": "course",
                                    "storageKey": null,
                                    "args": null,
                                    "concreteType": "Course",
                                    "plural": false,
                                    "selections": [
                                        (v2 /*: any*/),
                                        {
                                            "kind": "LinkedField",
                                            "alias": null,
                                            "name": "category",
                                            "storageKey": null,
                                            "args": null,
                                            "concreteType": "Category",
                                            "plural": false,
                                            "selections": [
                                                (v2 /*: any*/)
                                            ]
                                        }
                                    ]
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        "params": {
            "operationKind": "query",
            "name": "App_DelegatesProfile_Query",
            "id": null,
            "text": "query App_DelegatesProfile_Query(\n  $uuid: UUID!\n) {\n  delegate(uuid: $uuid) {\n    ...DelegateProfilePage_delegate\n  }\n}\n\nfragment DelegateProfilePage_delegate on Delegate {\n  firstName\n  lastName\n  myCourses {\n    status\n    course {\n      name\n      category {\n        name\n      }\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '8d71d48492ccccdeb87bd3c6e6ee8ae7';
export default node;
