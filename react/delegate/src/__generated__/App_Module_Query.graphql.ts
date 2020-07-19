/* tslint:disable */
/* eslint-disable */
/* @relayHash 257ca53a1543349c605d8e048d5a6ac6 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_Module_QueryVariables = {
    id: number;
    uuid: string;
};
export type App_Module_QueryResponse = {
    readonly user: {
        readonly myActiveCourse: {
            readonly " $fragmentRefs": FragmentRefs<"Module_myActiveCourse">;
        } | null;
    } | null;
    readonly module: {
        readonly " $fragmentRefs": FragmentRefs<"Module_module">;
    } | null;
};
export type App_Module_Query = {
    readonly response: App_Module_QueryResponse;
    readonly variables: App_Module_QueryVariables;
};



/*
query App_Module_Query(
  $id: Int!
  $uuid: UUID!
) {
  user {
    myActiveCourse(id: $id) {
      ...Module_myActiveCourse
    }
  }
  module(uuid: $uuid) {
    ...Module_module
  }
}

fragment CourseSyllabusCardFrag_course on Course {
  syllabus {
    __typename
    name
    type
    uuid
    complete
    ... on Module {
      syllabus {
        __typename
        name
        type
        uuid
        complete
      }
    }
  }
}

fragment Module_module on Module {
  name
  uuid
}

fragment Module_myActiveCourse on MyCourse {
  course {
    ...CourseSyllabusCardFrag_course
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "id",
            "type": "Int!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "uuid",
            "type": "UUID!",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "Variable",
            "name": "id",
            "variableName": "id"
        } as any)
    ], v2 = [
        ({
            "kind": "Variable",
            "name": "uuid",
            "variableName": "uuid"
        } as any)
    ], v3 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "__typename",
        "args": null,
        "storageKey": null
    } as any), v4 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "name",
        "args": null,
        "storageKey": null
    } as any), v5 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "type",
        "args": null,
        "storageKey": null
    } as any), v6 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "uuid",
        "args": null,
        "storageKey": null
    } as any), v7 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "complete",
        "args": null,
        "storageKey": null
    } as any);
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "App_Module_Query",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "user",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "User",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "myActiveCourse",
                            "storageKey": null,
                            "args": (v1 /*: any*/),
                            "concreteType": "MyCourse",
                            "plural": false,
                            "selections": [
                                {
                                    "kind": "FragmentSpread",
                                    "name": "Module_myActiveCourse",
                                    "args": null
                                }
                            ]
                        }
                    ]
                },
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "module",
                    "storageKey": null,
                    "args": (v2 /*: any*/),
                    "concreteType": "Module",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "FragmentSpread",
                            "name": "Module_module",
                            "args": null
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "App_Module_Query",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "user",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "User",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "myActiveCourse",
                            "storageKey": null,
                            "args": (v1 /*: any*/),
                            "concreteType": "MyCourse",
                            "plural": false,
                            "selections": [
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
                                            "kind": "LinkedField",
                                            "alias": null,
                                            "name": "syllabus",
                                            "storageKey": null,
                                            "args": null,
                                            "concreteType": null,
                                            "plural": true,
                                            "selections": [
                                                (v3 /*: any*/),
                                                (v4 /*: any*/),
                                                (v5 /*: any*/),
                                                (v6 /*: any*/),
                                                (v7 /*: any*/),
                                                {
                                                    "kind": "InlineFragment",
                                                    "type": "Module",
                                                    "selections": [
                                                        {
                                                            "kind": "LinkedField",
                                                            "alias": null,
                                                            "name": "syllabus",
                                                            "storageKey": null,
                                                            "args": null,
                                                            "concreteType": null,
                                                            "plural": true,
                                                            "selections": [
                                                                (v3 /*: any*/),
                                                                (v4 /*: any*/),
                                                                (v5 /*: any*/),
                                                                (v6 /*: any*/),
                                                                (v7 /*: any*/)
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
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "module",
                    "storageKey": null,
                    "args": (v2 /*: any*/),
                    "concreteType": "Module",
                    "plural": false,
                    "selections": [
                        (v4 /*: any*/),
                        (v6 /*: any*/)
                    ]
                }
            ]
        },
        "params": {
            "operationKind": "query",
            "name": "App_Module_Query",
            "id": null,
            "text": "query App_Module_Query(\n  $id: Int!\n  $uuid: UUID!\n) {\n  user {\n    myActiveCourse(id: $id) {\n      ...Module_myActiveCourse\n    }\n  }\n  module(uuid: $uuid) {\n    ...Module_module\n  }\n}\n\nfragment CourseSyllabusCardFrag_course on Course {\n  syllabus {\n    __typename\n    name\n    type\n    uuid\n    complete\n    ... on Module {\n      syllabus {\n        __typename\n        name\n        type\n        uuid\n        complete\n      }\n    }\n  }\n}\n\nfragment Module_module on Module {\n  name\n  uuid\n}\n\nfragment Module_myActiveCourse on MyCourse {\n  course {\n    ...CourseSyllabusCardFrag_course\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'a53a24ba4a72b40e073bfd1b77e7592a';
export default node;
