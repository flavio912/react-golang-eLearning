/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CourseStatus = "complete" | "failed" | "incomplete" | "%future added value";
export type UserType = "delegate" | "individual" | "manager" | "%future added value";
export type TrainingProgress_user = {
    readonly firstName: string;
    readonly type: UserType;
    readonly myCourses: ReadonlyArray<{
        readonly status: CourseStatus;
        readonly course: {
            readonly name: string;
            readonly category: {
                readonly name: string;
            } | null;
        };
    }> | null;
    readonly " $refType": "TrainingProgress_user";
};
export type TrainingProgress_user$data = TrainingProgress_user;
export type TrainingProgress_user$key = {
    readonly " $data"?: TrainingProgress_user$data;
    readonly " $fragmentRefs": FragmentRefs<"TrainingProgress_user">;
};



const node: ReaderFragment = (function () {
    var v0 = ({
        "kind": "ScalarField",
        "alias": null,
        "name": "name",
        "args": null,
        "storageKey": null
    } as any);
    return {
        "kind": "Fragment",
        "name": "TrainingProgress_user",
        "type": "User",
        "metadata": null,
        "argumentDefinitions": [],
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
                "name": "type",
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
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "course",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "Course",
                        "plural": false,
                        "selections": [
                            (v0 /*: any*/),
                            {
                                "kind": "LinkedField",
                                "alias": null,
                                "name": "category",
                                "storageKey": null,
                                "args": null,
                                "concreteType": "Category",
                                "plural": false,
                                "selections": [
                                    (v0 /*: any*/)
                                ]
                            }
                        ]
                    }
                ]
            }
        ]
    } as any;
})();
(node as any).hash = '67051e1e524681535447d2fbade72cab';
export default node;
