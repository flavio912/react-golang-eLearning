/* tslint:disable */
/* eslint-disable */
/* @relayHash 69c27049fdf0d590c266f5029b15b446 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_Course_QueryVariables = {
    ident: number;
};
export type App_Course_QueryResponse = {
    readonly user: {
        readonly myActiveCourse: {
            readonly " $fragmentRefs": FragmentRefs<"OnlineCourse_myActiveCourse">;
        } | null;
    } | null;
};
export type App_Course_Query = {
    readonly response: App_Course_QueryResponse;
    readonly variables: App_Course_QueryVariables;
};



/*
query App_Course_Query(
  $ident: Int!
) {
  user {
    myActiveCourse(id: $ident) {
      ...OnlineCourse_myActiveCourse
    }
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

fragment OnlineCourse_myActiveCourse on MyCourse {
  status
  enrolledAt
  upTo
  course {
    ident: id
    name
    excerpt
    introduction
    howToComplete
    whatYouLearn
    hoursToComplete
    syllabus {
      __typename
      name
      type
      uuid
      ... on Module {
        syllabus {
          __typename
          name
          uuid
          type
        }
      }
    }
    ...CourseSyllabusCardFrag_course
  }
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "ident",
            "type": "Int!",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "Variable",
            "name": "id",
            "variableName": "ident"
        } as any)
    ], v2 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "name",
        "args": null,
        "storageKey": null
    } as any), v3 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "__typename",
        "args": null,
        "storageKey": null
    } as any), v4 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "type",
        "args": null,
        "storageKey": null
    } as any), v5 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "uuid",
        "args": null,
        "storageKey": null
    } as any), v6 = ({
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
            "name": "App_Course_Query",
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
                                    "name": "OnlineCourse_myActiveCourse",
                                    "args": null
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "App_Course_Query",
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
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "status",
                                    "args": null,
                                    "storageKey": null
                                },
                                {
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "enrolledAt",
                                    "args": null,
                                    "storageKey": null
                                },
                                {
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "upTo",
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
                                        (v2 /*: any*/),
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "excerpt",
                                            "args": null,
                                            "storageKey": null
                                        },
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "introduction",
                                            "args": null,
                                            "storageKey": null
                                        },
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "howToComplete",
                                            "args": null,
                                            "storageKey": null
                                        },
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "whatYouLearn",
                                            "args": null,
                                            "storageKey": null
                                        },
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "hoursToComplete",
                                            "args": null,
                                            "storageKey": null
                                        },
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
                                                (v2 /*: any*/),
                                                (v4 /*: any*/),
                                                (v5 /*: any*/),
                                                (v6 /*: any*/),
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
                                                                (v2 /*: any*/),
                                                                (v5 /*: any*/),
                                                                (v4 /*: any*/),
                                                                (v6 /*: any*/)
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
                }
            ]
        },
        "params": {
            "operationKind": "query",
            "name": "App_Course_Query",
            "id": null,
            "text": "query App_Course_Query(\n  $ident: Int!\n) {\n  user {\n    myActiveCourse(id: $ident) {\n      ...OnlineCourse_myActiveCourse\n    }\n  }\n}\n\nfragment CourseSyllabusCardFrag_course on Course {\n  syllabus {\n    __typename\n    name\n    type\n    uuid\n    complete\n    ... on Module {\n      syllabus {\n        __typename\n        name\n        type\n        uuid\n        complete\n      }\n    }\n  }\n}\n\nfragment OnlineCourse_myActiveCourse on MyCourse {\n  status\n  enrolledAt\n  upTo\n  course {\n    ident: id\n    name\n    excerpt\n    introduction\n    howToComplete\n    whatYouLearn\n    hoursToComplete\n    syllabus {\n      __typename\n      name\n      type\n      uuid\n      ... on Module {\n        syllabus {\n          __typename\n          name\n          uuid\n          type\n        }\n      }\n    }\n    ...CourseSyllabusCardFrag_course\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '9af72eeb279e1a5e923760216eba6bc2';
export default node;
