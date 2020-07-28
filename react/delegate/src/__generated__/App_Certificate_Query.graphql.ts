/* tslint:disable */
/* eslint-disable */
/* @relayHash 9fc96a6410ef58b5e81dceb2efcbf06a */

import { ConcreteRequest } from "relay-runtime";
import { FragmentRefs } from "relay-runtime";
export type App_Certificate_QueryVariables = {
    token: string;
};
export type App_Certificate_QueryResponse = {
    readonly certificateInfo: {
        readonly " $fragmentRefs": FragmentRefs<"CertGenerator_certificateInfo">;
    };
};
export type App_Certificate_Query = {
    readonly response: App_Certificate_QueryResponse;
    readonly variables: App_Certificate_QueryVariables;
};



/*
query App_Certificate_Query(
  $token: String!
) {
  certificateInfo(token: $token) {
    ...CertGenerator_certificateInfo
  }
}

fragment CertGenerator_certificateInfo on CertificateInfo {
  courseTitle
  expiryDate
  completionDate
  companyName
  takerFirstName
  takerLastName
  certificateBodyURL
  regulationText
  CAANo
  title
  instructorName
  instructorCIN
  instructorSignatureURL
  certificateNumber
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "token",
            "type": "String!",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "Variable",
            "name": "token",
            "variableName": "token"
        } as any)
    ];
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "App_Certificate_Query",
            "type": "Query",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "certificateInfo",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
                    "concreteType": "CertificateInfo",
                    "plural": false,
                    "selections": [
                        {
                            "kind": "FragmentSpread",
                            "name": "CertGenerator_certificateInfo",
                            "args": null
                        }
                    ]
                }
            ]
        },
        "operation": {
            "kind": "Operation",
            "name": "App_Certificate_Query",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": [
                {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "certificateInfo",
                    "storageKey": null,
                    "args": (v1 /*: any*/),
                    "concreteType": "CertificateInfo",
                    "plural": false,
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
                }
            ]
        },
        "params": {
            "operationKind": "query",
            "name": "App_Certificate_Query",
            "id": null,
            "text": "query App_Certificate_Query(\n  $token: String!\n) {\n  certificateInfo(token: $token) {\n    ...CertGenerator_certificateInfo\n  }\n}\n\nfragment CertGenerator_certificateInfo on CertificateInfo {\n  courseTitle\n  expiryDate\n  completionDate\n  companyName\n  takerFirstName\n  takerLastName\n  certificateBodyURL\n  regulationText\n  CAANo\n  title\n  instructorName\n  instructorCIN\n  instructorSignatureURL\n  certificateNumber\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = '1857166aae02536f4c536cfe21484a30';
export default node;
