/* tslint:disable */
/* eslint-disable */
/* @relayHash e2410d54af8caafe26fdf621f7e61e28 */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_CourseDetailsPage_QueryVariables = {
    id: number;
};
export type App_CourseDetailsPage_QueryResponse = {
    readonly course: {
        readonly " $fragmentRefs": FragmentRefs<"CourseDetailsPage_course">;
    } | null;
};
export type App_CourseDetailsPage_Query = {
    readonly response: App_CourseDetailsPage_QueryResponse;
    readonly variables: App_CourseDetailsPage_QueryVariables;
};



/*
query App_CourseDetailsPage_Query(
  $id: Int!
) {
  course(id: $id) {
    ...CourseDetailsPage_course
  }
}

fragment CourseDetailsPage_course on Course {
  ident: id
  name
  price
  excerpt
  introduction
  bannerImageURL
  howToComplete
  hoursToComplete
  whatYouLearn
  requirements
  category {
    name
    color
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
        } as any)
    ], v1 = [
        ({
            "kind": "Variable",
            "name": "id",
            "variableName": "id"
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
            "name": "App_CourseDetailsPage_Query",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "course",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
                    "concreteType": "Course",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "FragmentSpread",
                            "name": "CourseDetailsPage_course",
                            "args": null
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "App_CourseDetailsPage_Query",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "course",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
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
                            "name": "price",
                            "args": null,
                            "storageKey": null
                        },
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
                            "name": "bannerImageURL",
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
                            "name": "hoursToComplete",
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
                            "name": "requirements",
                            "args": null,
                            "storageKey": null
                        },
                        {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "category",
                            "storageKey": null,
                            "args": null,
                            "concreteType": "Category",
                            "plural": false,
                            "selections": [
                                (v2 /*: any*/),
                                {
                                    "kind": "ScalarField",
                                    "alias": null,
                                    "name": "color",
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
            "name": "App_CourseDetailsPage_Query",
            "id": null,
            "text": "query App_CourseDetailsPage_Query(\n  $id: Int!\n) {\n  course(id: $id) {\n    ...CourseDetailsPage_course\n  }\n}\n\nfragment CourseDetailsPage_course on Course {\n  ident: id\n  name\n  price\n  excerpt\n  introduction\n  bannerImageURL\n  howToComplete\n  hoursToComplete\n  whatYouLearn\n  requirements\n  category {\n    name\n    color\n  }\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'b3146e50486583eee7942695c661a87e';
export default node;
