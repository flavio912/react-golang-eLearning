/* tslint:disable */
/* eslint-disable */
/* @relayHash 4238750ceb49f1d8d3550ff9a241a095 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_Holder_QueryVariables = {};
export type App_Holder_QueryResponse = {
    readonly manager: {
        readonly " $fragmentRefs": FragmentRefs<"AppHolder_manager">;
    } | null;
};
export type App_Holder_Query = {
    readonly response: App_Holder_QueryResponse;
    readonly variables: App_Holder_QueryVariables;
};



/*
query App_Holder_Query {
  manager {
    ...AppHolder_manager
  }
}

fragment AppHolder_manager on Manager {
  firstName
  lastName
  profileImageUrl
  company {
    logoURL
  }
}
*/

const node: ConcreteRequest = ({
    "kind": "Request",
    "fragment": {
        "kind": "Fragment",
        "name": "App_Holder_Query",
        "type": "Query",
        "metadata": null,
        "argumentDefinitions": [],
        "selections": [
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "manager",
                "storageKey": null,
                "args": null,
                "concreteType": "Manager",
                "plural": false,
                "selections": [
                    {
                        "kind": "FragmentSpread",
                        "name": "AppHolder_manager",
                        "args": null
                    }
                ]
            }
        ]
    },
    "operation": {
        "kind": "Operation",
        "name": "App_Holder_Query",
        "argumentDefinitions": [],
        "selections": [
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "manager",
                "storageKey": null,
                "args": null,
                "concreteType": "Manager",
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
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "profileImageUrl",
                        "args": null,
                        "storageKey": null
                    },
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
                                "name": "logoURL",
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
        "name": "App_Holder_Query",
        "id": null,
        "text": "query App_Holder_Query {\n  manager {\n    ...AppHolder_manager\n  }\n}\n\nfragment AppHolder_manager on Manager {\n  firstName\n  lastName\n  profileImageUrl\n  company {\n    logoURL\n  }\n}\n",
        "metadata": {}
    }
} as any);
(node as any).hash = 'ae1f32a58cdf66e3f2f3dc9abd5b102d';
export default node;
