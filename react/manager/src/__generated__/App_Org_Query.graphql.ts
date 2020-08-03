/* tslint:disable */
/* eslint-disable */
/* @relayHash 5d1925c54c9902fd35e229d60b7075b0 */

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

fragment ActivityCard_activity on ActivityPage {
  edges {
    ...UserUpdate_activity
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
    delegates {
      pageInfo {
        total
      }
    }
    activity(page: {limit: 6}) {
      ...ActivityCard_activity
    }
  }
}

fragment UserUpdate_activity on Activity {
  type
  createdAt
  course {
    name
  }
  user {
    firstName
    lastName
    profileImageUrl
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "firstName",
        "args": null,
        "storageKey": null
    } as any), v1 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "lastName",
        "args": null,
        "storageKey": null
    } as any), v2 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "createdAt",
        "args": null,
        "storageKey": null
    } as any), v3 = ({
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
                        (v0 /*: any*/),
                        (v1 /*: any*/),
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
                        (v2 /*: any*/),
                        {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "company",
                            "storageKey": null,
                            "args": null,
                            "concreteType": "Company",
                            "plural": false,
                            "selections": [
                                (v3 /*: any*/),
                                {
                                    "kind": "LinkedField",
                                    "alias": null,
                                    "name": "delegates",
                                    "storageKey": null,
                                    "args": null,
                                    "concreteType": "DelegatePage",
                                    "plural": false,
                                    "selections": [
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
                                },
                                {
                                    "kind": "LinkedField",
                                    "alias": null,
                                    "name": "activity",
                                    "storageKey": "activity(page:{\"limit\":6})",
                                    "args": [
                                        {
                                            "kind": "Literal",
                                            "name": "page",
                                            "value": {
                                                "limit": 6
                                            }
                                        }
                                    ],
                                    "concreteType": "ActivityPage",
                                    "plural": false,
                                    "selections": [
                                        {
                                            "kind": "LinkedField",
                                            "alias": null,
                                            "name": "edges",
                                            "storageKey": null,
                                            "args": null,
                                            "concreteType": "Activity",
                                            "plural": true,
                                            "selections": [
                                                {
                                                    "kind": "ScalarField",
                                                    "alias": null,
                                                    "name": "type",
                                                    "args": null,
                                                    "storageKey": null
                                                },
                                                (v2 /*: any*/),
                                                {
                                                    "kind": "LinkedField",
                                                    "alias": null,
                                                    "name": "course",
                                                    "storageKey": null,
                                                    "args": null,
                                                    "concreteType": "Course",
                                                    "plural": false,
                                                    "selections": [
                                                        (v3 /*: any*/)
                                                    ]
                                                },
                                                {
                                                    "kind": "LinkedField",
                                                    "alias": null,
                                                    "name": "user",
                                                    "storageKey": null,
                                                    "args": null,
                                                    "concreteType": "User",
                                                    "plural": false,
                                                    "selections": [
                                                        (v0 /*: any*/),
                                                        (v1 /*: any*/),
                                                        {
                                                            "kind": "ScalarField",
                                                            "alias": null,
                                                            "name": "profileImageUrl",
                                                            "args": null,
                                                            "storageKey": null
                                                        }
                                                    ]
                                                }
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
            "name": "App_Org_Query",
            "id": null,
            "text": "query App_Org_Query {\n  manager {\n    ...OrgOverview_manager\n  }\n}\n\nfragment ActivityCard_activity on ActivityPage {\n  edges {\n    ...UserUpdate_activity\n  }\n}\n\nfragment OrgOverview_manager on Manager {\n  firstName\n  lastName\n  email\n  telephone\n  createdAt\n  company {\n    name\n    delegates {\n      pageInfo {\n        total\n      }\n    }\n    activity(page: {limit: 6}) {\n      ...ActivityCard_activity\n    }\n  }\n}\n\nfragment UserUpdate_activity on Activity {\n  type\n  createdAt\n  course {\n    name\n  }\n  user {\n    firstName\n    lastName\n    profileImageUrl\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'f3bdd6ff4674f9f0f2b1e33b22045368';
export default node;
