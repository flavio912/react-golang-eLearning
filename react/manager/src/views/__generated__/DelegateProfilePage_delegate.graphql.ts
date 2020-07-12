/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type DelegateProfilePage_delegate = {
    readonly firstName: string;
    readonly lastName: string;
    readonly activeCourses: ReadonlyArray<{
        readonly currentAttempt: number;
        readonly course: {
            readonly name: string;
            readonly category: {
                readonly name: string;
            } | null;
        };
    }> | null;
    readonly " $refType": "DelegateProfilePage_delegate";
};
export type DelegateProfilePage_delegate$data = DelegateProfilePage_delegate;
export type DelegateProfilePage_delegate$key = {
    readonly " $data"?: DelegateProfilePage_delegate$data;
    readonly " $fragmentRefs": FragmentRefs<"DelegateProfilePage_delegate">;
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
        "name": "DelegateProfilePage_delegate",
        "type": "Delegate",
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
                "name": "lastName",
                "args": null,
                "storageKey": null
            },
            {
                "kind": "LinkedField",
                "alias": null,
                "name": "activeCourses",
                "storageKey": null,
                "args": null,
                "concreteType": "ActiveCourse",
                "plural": true,
                "selections": [
                    {
                        "kind": "ScalarField",
                        "alias": null,
                        "name": "currentAttempt",
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
(node as any).hash = '607d0b5bba25577535577a9dd89b82aa';
export default node;
