/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CoursesPage_classroomCourses = {
    readonly edges: ReadonlyArray<{
        readonly uuid: string;
    } | null>;
    readonly pageInfo: {
        readonly total: number;
        readonly offset: number;
        readonly limit: number;
        readonly given: number;
    } | null;
    readonly " $refType": "CoursesPage_classroomCourses";
};
export type CoursesPage_classroomCourses$data = CoursesPage_classroomCourses;
export type CoursesPage_classroomCourses$key = {
    readonly " $data"?: CoursesPage_classroomCourses$data;
    readonly " $fragmentRefs": FragmentRefs<"CoursesPage_classroomCourses">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "CoursesPage_classroomCourses",
    "type": "ClassroomCoursePage",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
        {
            "kind": "LinkedField",
            "alias": null,
            "name": "edges",
            "storageKey": null,
            "args": null,
            "concreteType": "ClassroomCourse",
            "plural": true,
            "selections": [
                {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "uuid",
                    "args": null,
                    "storageKey": null
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
                    "name": "offset",
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
                    "name": "given",
                    "args": null,
                    "storageKey": null
                }
            ]
        }
    ]
} as any);
(node as any).hash = '5fbd5dd17f377dfe16492f3eacd3b456';
export default node;
