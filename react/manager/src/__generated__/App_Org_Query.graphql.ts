/* tslint:disable */
/* eslint-disable */
/* @relayHash 458b792d04485d5b0ddc1fcd782a669d */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_Org_QueryVariables = {};
export type App_Org_QueryResponse = {
    readonly manager: {
        readonly " $fragmentRefs": FragmentRefs<"OrgOverview_manager">;
    } | null;
};
export type App_Org_Query = {
    readonly response: App_Org_QueryResponse;
    readonly variables: App_Org_QueryVariables;
};



/*
query App_Org_Query {
  manager {
    ...OrgOverview_manager
  }
}

fragment OrgOverview_manager on Manager {
  firstName
  lastName
  uuid
}
*/

const node: ConcreteRequest = ({
    "kind": "Request",
    "fragment": {
        "kind": "Fragment",
        "name": "App_Org_Query",
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
                        "name": "OrgOverview_manager",
                        "args": null
                    }
                ]
            }
        ]
    },
    "operation": {
        "kind": "Operation",
        "name": "App_Org_Query",
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
                        "name": "uuid",
                        "args": null,
                        "storageKey": null
                    }
                ]
            }
        ]
    },
    "params": {
        "operationKind": "query",
        "name": "App_Org_Query",
        "id": null,
        "text": "query App_Org_Query {\n  manager {\n    ...OrgOverview_manager\n  }\n}\n\nfragment OrgOverview_manager on Manager {\n  firstName\n  lastName\n  uuid\n}\n",
        "metadata": {}
    }
} as any);
(node as any).hash = 'f3bdd6ff4674f9f0f2b1e33b22045368';
export default node;
