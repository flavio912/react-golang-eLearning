/* tslint:disable */
/* eslint-disable */

import { ReaderFragment } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type CertGenerator_certificateInfo = {
    readonly courseTitle: string;
    readonly expiryDate: string;
    readonly completionDate: string;
    readonly companyName: string | null;
    readonly takerFirstName: string;
    readonly takerLastName: string;
    readonly certificateBodyURL: string | null;
    readonly regulationText: string;
    readonly CAANo: string | null;
    readonly title: string;
    readonly instructorName: string;
    readonly instructorCIN: string;
    readonly instructorSignatureURL: string | null;
    readonly certificateNumber: string;
    readonly " $refType": "CertGenerator_certificateInfo";
};
export type CertGenerator_certificateInfo$data = CertGenerator_certificateInfo;
export type CertGenerator_certificateInfo$key = {
    readonly " $data"?: CertGenerator_certificateInfo$data;
    readonly " $fragmentRefs": FragmentRefs<"CertGenerator_certificateInfo">;
};



const node: ReaderFragment = ({
    "kind": "Fragment",
    "name": "CertGenerator_certificateInfo",
    "type": "CertificateInfo",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "courseTitle",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "expiryDate",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "completionDate",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "companyName",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "takerFirstName",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "takerLastName",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "certificateBodyURL",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "regulationText",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "CAANo",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "title",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "instructorName",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "instructorCIN",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "instructorSignatureURL",
            "args": null,
            "storageKey": null
        },
        {
            "kind": "ScalarField",
            "alias": null,
            "name": "certificateNumber",
            "args": null,
            "storageKey": null
        }
    ]
} as any);
(node as any).hash = '687d78ac064d0f84120e6baffb16bdca';
export default node;
