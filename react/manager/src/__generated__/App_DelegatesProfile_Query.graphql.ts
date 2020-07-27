/* tslint:disable */
/* eslint-disable */
/* @relayHash 1b2a79c3686ac395a228f3a04d2aac46 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_DelegatesProfile_QueryVariables = {
    uuid: string;
    offset?: number | null;
    limit?: number | null;
};
export type App_DelegatesProfile_QueryResponse = {
    readonly delegate: {
        readonly activity: {
            readonly " $fragmentRefs": FragmentRefs<"DelegateProfilePage_activity">;
        } | null;
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
  $offset: Int
  $limit: Int
) {
  delegate(uuid: $uuid) {
    activity(page: {offset: $offset, limit: $limit}) {
      ...DelegateProfilePage_activity
    }
    ...DelegateProfilePage_delegate
  }
}

fragment DelegateProfilePage_activity on ActivityPage {
  edges {
    type
    createdAt
    course {
      ident: id
      name
    }
  }
  pageInfo {
    total
    limit
    offset
  }
}

fragment DelegateProfilePage_delegate on Delegate {
  uuid
  firstName
  lastName
  email
  jobTitle
  telephone
  TTC_ID
  lastLogin
  myCourses {
    status
    minutesTracked
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
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "offset",
            "type": "Int",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "limit",
            "type": "Int",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "Variable",
            "name": "uuid",
            "variableName": "uuid"
        } as any)
    ], v2 = [
        ({
            "kind": "ObjectValue",
            "name": "page",
            "fields": [
                {
                    "kind": "Variable",
                    "name": "limit",
                    "variableName": "limit"
                },
                {
                    "kind": "Variable",
                    "name": "offset",
                    "variableName": "offset"
                }
            ]
        } as any)
    ], v3 = ({
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
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "activity",
                            "storageKey": null,
                            "args": (v2 /*: any*/),
                            "concreteType": "ActivityPage",
                            "plural": false,
                            "selections": [
                                {
                                    "kind": "FragmentSpread",
                                    "name": "DelegateProfilePage_activity",
                                    "args": null
                                }
                            ]
                        },
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
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "activity",
                            "storageKey": null,
                            "args": (v2 /*: any*/),
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
                                            "name": "course",
                                            "storageKey": null,
                                            "args": null,
                                            "concreteType": "Course",
                                            "plural": false,
                                            "selections": [
                                                {
                                                    "kind": "ScalarField",
                                                    "alias": "ident",
                                                    "name": "id",
                                                    "args": null,
                                                    "storageKey": null
                                                },
                                                (v3 /*: any*/)
                                            ]
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
                                        },
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "limit",
                                            "args": null,
                                            "storageKey": null
                                        },
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "offset",
                                            "args": null,
                                            "storageKey": null
                                        }
                                    ]
                                }
                            ]
                        },
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
                            "name": "jobTitle",
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
                            "name": "TTC_ID",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "lastLogin",
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
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "minutesTracked",
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
                                        (v3 /*: any*/),
                                        {
                                            "kind": "LinkedField",
                                            "alias": null,
                                            "name": "category",
                                            "storageKey": null,
                                            "args": null,
                                            "concreteType": "Category",
                                            "plural": false,
                                            "selections": [
                                                (v3 /*: any*/)
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
            "text": "query App_DelegatesProfile_Query(\n  $uuid: UUID!\n  $offset: Int\n  $limit: Int\n) {\n  delegate(uuid: $uuid) {\n    activity(page: {offset: $offset, limit: $limit}) {\n      ...DelegateProfilePage_activity\n    }\n    ...DelegateProfilePage_delegate\n  }\n}\n\nfragment DelegateProfilePage_activity on ActivityPage {\n  edges {\n    type\n    createdAt\n    course {\n      ident: id\n      name\n    }\n  }\n  pageInfo {\n    total\n    limit\n    offset\n  }\n}\n\nfragment DelegateProfilePage_delegate on Delegate {\n  uuid\n  firstName\n  lastName\n  email\n  jobTitle\n  telephone\n  TTC_ID\n  lastLogin\n  myCourses {\n    status\n    minutesTracked\n    course {\n      name\n      category {\n        name\n      }\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'f86767931d343d55d024c4bd5d9da63d';
export default node;
