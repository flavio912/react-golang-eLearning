/* tslint:disable */
/* eslint-disable */
/* @relayHash b6c1010ed5ded464d927b1477e04962b */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_Lesson_QueryVariables = {
    id: number;
    uuid: string;
};
export type App_Lesson_QueryResponse = {
    readonly user: {
        readonly myActiveCourse: {
            readonly " $fragmentRefs": FragmentRefs<"Lesson_myActiveCourse">;
        } | null;
    } | null;
    readonly lesson: {
        readonly " $fragmentRefs": FragmentRefs<"Lesson_lesson">;
    } | null;
};
export type App_Lesson_Query = {
    readonly response: App_Lesson_QueryResponse;
    readonly variables: App_Lesson_QueryVariables;
};



/*
query App_Lesson_Query(
  $id: Int!
  $uuid: UUID!
) {
  user {
    myActiveCourse(id: $id) {
      ...Lesson_myActiveCourse
    }
  }
  lesson(uuid: $uuid) {
    ...Lesson_lesson
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

fragment Lesson_lesson on Lesson {
  name
  uuid
  voiceoverURL
  description
  transcript
}

fragment Lesson_myActiveCourse on MyCourse {
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
  upTo
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
            "name": "App_Lesson_Query",
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
                                    "name": "Lesson_myActiveCourse",
                                    "args": null
                                }
                            ]
                        }
                    ]
                },
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "lesson",
                    "storageKey": null,
                    "args": (v2 /*: any*/),
                    "concreteType": "Lesson",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "FragmentSpread",
                            "name": "Lesson_lesson",
                            "args": null
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "App_Lesson_Query",
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
                                },
                                {
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "upTo",
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
                    "name": "lesson",
                    "storageKey": null,
                    "args": (v2 /*: any*/),
                    "concreteType": "Lesson",
                    "plural": false,
                    "selections": [
                        (v4 /*: any*/),
                        (v6 /*: any*/),
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "voiceoverURL",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "description",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "ScalarField",
                            "alias": null,
                            "name": "transcript",
                            "args": null,
                            "storageKey": null
                        }
                    ]
                }
            ]
        },
        "params": {
            "operationKind": "query",
            "name": "App_Lesson_Query",
            "id": null,
            "text": "query App_Lesson_Query(\n  $id: Int!\n  $uuid: UUID!\n) {\n  user {\n    myActiveCourse(id: $id) {\n      ...Lesson_myActiveCourse\n    }\n  }\n  lesson(uuid: $uuid) {\n    ...Lesson_lesson\n  }\n}\n\nfragment CourseSyllabusCardFrag_course on Course {\n  syllabus {\n    __typename\n    name\n    type\n    uuid\n    complete\n    ... on Module {\n      syllabus {\n        __typename\n        name\n        type\n        uuid\n        complete\n      }\n    }\n  }\n}\n\nfragment Lesson_lesson on Lesson {\n  name\n  uuid\n  voiceoverURL\n  description\n  transcript\n}\n\nfragment Lesson_myActiveCourse on MyCourse {\n  course {\n    syllabus {\n      __typename\n      name\n      type\n      uuid\n      ... on Module {\n        syllabus {\n          __typename\n          name\n          uuid\n          type\n        }\n      }\n    }\n    ...CourseSyllabusCardFrag_course\n  }\n  upTo\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'd88c4b7c9bcbf67db1012b8fa7832617';
export default node;
