/* tslint:disable */
/* eslint-disable */
/* @relayHash bcd5a67afcca3352db500552cee0795e */

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
  email
  telephone
  createdAt
  company {
    name
  }
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
                        "name": "email",
                        "args": null,
                        "storageKey": null
                    },
                    {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "telephone",
                        "args": null,
                        "storageKey": null
                    },
                    {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "createdAt",
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
                                "name": "name",
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
        "name": "App_Org_Query",
        "id": null,
        "text": "query App_Org_Query {\n  manager {\n    ...OrgOverview_manager\n  }\n}\n\nfragment OrgOverview_manager on Manager {\n  firstName\n  lastName\n  email\n  telephone\n  createdAt\n  company {\n    name\n  }\n}\n",
        "metadata": {}
    }
} as any);
(node as any).hash = 'f3bdd6ff4674f9f0f2b1e33b22045368';
export default node;
