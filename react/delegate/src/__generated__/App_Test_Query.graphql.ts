/* tslint:disable */
/* eslint-disable */
/* @relayHash 848df1fc0257102992284c02f7b3fb05 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_Test_QueryVariables = {
    id: number;
    uuid: string;
};
export type App_Test_QueryResponse = {
    readonly user: {
        readonly myActiveCourse: {
            readonly " $fragmentRefs": FragmentRefs<"Test_myActiveCourse">;
        } | null;
    } | null;
    readonly test: {
        readonly " $fragmentRefs": FragmentRefs<"Test_test">;
    } | null;
};
export type App_Test_Query = {
    readonly response: App_Test_QueryResponse;
    readonly variables: App_Test_QueryVariables;
};



/*
query App_Test_Query(
  $id: Int!
  $uuid: UUID!
) {
  user {
    myActiveCourse(id: $id) {
      ...Test_myActiveCourse
    }
  }
  test(uuid: $uuid) {
    ...Test_test
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

fragment Test_myActiveCourse on MyCourse {
  course {
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

fragment Test_test on Test {
  name
  uuid
  questions {
    uuid
    text
    questionType
    answers {
      uuid
      text
      imageURL
    }
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
    } as any), v8 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "text",
        "args": null,
        "storageKey": null
    } as any);
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "App_Test_Query",
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
                                    "name": "Test_myActiveCourse",
                                    "args": null
                                }
                            ]
                        }
                    ]
                },
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "test",
                    "storageKey": null,
                    "args": (v2 /*: any*/),
                    "concreteType": "Test",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "FragmentSpread",
                            "name": "Test_test",
                            "args": null
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "App_Test_Query",
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
                                                                (v6 /*: any*/),
                                                                (v5 /*: any*/),
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
                    "name": "test",
                    "storageKey": null,
                    "args": (v2 /*: any*/),
                    "concreteType": "Test",
                    "plural": false,
                    "selections": [
                        (v4 /*: any*/),
                        (v6 /*: any*/),
                        {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "questions",
                            "storageKey": null,
                            "args": null,
                            "concreteType": "Question",
                            "plural": true,
                            "selections": [
                                (v6 /*: any*/),
                                (v8 /*: any*/),
                                {
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "questionType",
                                    "args": null,
                                    "storageKey": null
                                },
                                {
                                    "kind": "LinkedField",
                                    "alias": null,
                                    "name": "answers",
                                    "storageKey": null,
                                    "args": null,
                                    "concreteType": "Answer",
                                    "plural": true,
                                    "selections": [
                                        (v6 /*: any*/),
                                        (v8 /*: any*/),
                                        {
                                            "kind": "ScalarField",
                                            "alias": null,
                                            "name": "imageURL",
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
        },
        "params": {
            "operationKind": "query",
            "name": "App_Test_Query",
            "id": null,
            "text": "query App_Test_Query(\n  $id: Int!\n  $uuid: UUID!\n) {\n  user {\n    myActiveCourse(id: $id) {\n      ...Test_myActiveCourse\n    }\n  }\n  test(uuid: $uuid) {\n    ...Test_test\n  }\n}\n\nfragment CourseSyllabusCardFrag_course on Course {\n  syllabus {\n    __typename\n    name\n    type\n    uuid\n    complete\n    ... on Module {\n      syllabus {\n        __typename\n        name\n        type\n        uuid\n        complete\n      }\n    }\n  }\n}\n\nfragment Test_myActiveCourse on MyCourse {\n  course {\n    syllabus {\n      __typename\n      name\n      type\n      uuid\n      ... on Module {\n        syllabus {\n          __typename\n          name\n          uuid\n          type\n        }\n      }\n    }\n    ...CourseSyllabusCardFrag_course\n  }\n}\n\nfragment Test_test on Test {\n  name\n  uuid\n  questions {\n    uuid\n    text\n    questionType\n    answers {\n      uuid\n      text\n      imageURL\n    }\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '06b614ac00a6c2c6e42c25d8edfd2307';
export default node;
