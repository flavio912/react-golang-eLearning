/* tslint:disable */
/* eslint-disable */
/* @relayHash 94bc53d93a4c82c12f483cd0d72c6e3f */

import { ConcreteRequest } from "relay-runtime";
export type RegisterCompany_CompanyRequestMutationVariables = {
    companyName: string;
    addr1: string;
    addr2: string;
    county: string;
    postcode: string;
    country: string;
    firstName: string;
    lastName: string;
    email: string;
    password: string;
};
export type RegisterCompany_CompanyRequestMutationResponse = {
    readonly createCompanyRequest: boolean;
};
export type RegisterCompany_CompanyRequestMutation = {
    readonly response: RegisterCompany_CompanyRequestMutationResponse;
    readonly variables: RegisterCompany_CompanyRequestMutationVariables;
};



/*
mutation RegisterCompany_CompanyRequestMutation(
  $companyName: String!
  $addr1: String!
  $addr2: String!
  $county: String!
  $postcode: String!
  $country: String!
  $firstName: String!
  $lastName: String!
  $email: String!
  $password: String!
) {
  createCompanyRequest(company: {companyName: $companyName, addressLine1: $addr1, addressLine2: $addr2, county: $county, postCode: $postcode, country: $country, contactEmail: $email}, manager: {firstName: $firstName, lastName: $lastName, email: $email, jobTitle: "", telephone: "", password: $password}, recaptcha: "")
}
*/

const node: ConcreteRequest = (function () {
    var v0 = [
        ({
            "kind": "LocalArgument",
            "name": "companyName",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "addr1",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "addr2",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "county",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "postcode",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "country",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "firstName",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "lastName",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "email",
            "type": "String!",
            "defaultValue": null
        } as any),
        ({
            "kind": "LocalArgument",
            "name": "password",
            "type": "String!",
            "defaultValue": null
        } as any)
    ], v1 = [
        ({
            "kind": "ScalarField",
            "alias": null,
            "name": "createCompanyRequest",
            "args": [
                {
                    "kind": "ObjectValue",
                    "name": "company",
                    "fields": [
                        {
                            "kind": "Variable",
                            "name": "addressLine1",
                            "variableName": "addr1"
                        },
                        {
                            "kind": "Variable",
                            "name": "addressLine2",
                            "variableName": "addr2"
                        },
                        {
                            "kind": "Variable",
                            "name": "companyName",
                            "variableName": "companyName"
                        },
                        {
                            "kind": "Variable",
                            "name": "contactEmail",
                            "variableName": "email"
                        },
                        {
                            "kind": "Variable",
                            "name": "country",
                            "variableName": "country"
                        },
                        {
                            "kind": "Variable",
                            "name": "county",
                            "variableName": "county"
                        },
                        {
                            "kind": "Variable",
                            "name": "postCode",
                            "variableName": "postcode"
                        }
                    ]
                },
                {
                    "kind": "ObjectValue",
                    "name": "manager",
                    "fields": [
                        {
                            "kind": "Variable",
                            "name": "email",
                            "variableName": "email"
                        },
                        {
                            "kind": "Variable",
                            "name": "firstName",
                            "variableName": "firstName"
                        },
                        {
                            "kind": "Literal",
                            "name": "jobTitle",
                            "value": ""
                        },
                        {
                            "kind": "Variable",
                            "name": "lastName",
                            "variableName": "lastName"
                        },
                        {
                            "kind": "Variable",
                            "name": "password",
                            "variableName": "password"
                        },
                        {
                            "kind": "Literal",
                            "name": "telephone",
                            "value": ""
                        }
                    ]
                },
                {
                    "kind": "Literal",
                    "name": "recaptcha",
                    "value": ""
                }
            ],
            "storageKey": null
        } as any)
    ];
    return {
        "kind": "Request",
        "fragment": {
            "kind": "Fragment",
            "name": "RegisterCompany_CompanyRequestMutation",
            "type": "Mutation",
            "metadata": null,
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "operation": {
            "kind": "Operation",
            "name": "RegisterCompany_CompanyRequestMutation",
            "argumentDefinitions": (v0 /*: any*/),
            "selections": (v1 /*: any*/)
        },
        "params": {
            "operationKind": "mutation",
            "name": "RegisterCompany_CompanyRequestMutation",
            "id": null,
            "text": "mutation RegisterCompany_CompanyRequestMutation(\n  $companyName: String!\n  $addr1: String!\n  $addr2: String!\n  $county: String!\n  $postcode: String!\n  $country: String!\n  $firstName: String!\n  $lastName: String!\n  $email: String!\n  $password: String!\n) {\n  createCompanyRequest(company: {companyName: $companyName, addressLine1: $addr1, addressLine2: $addr2, county: $county, postCode: $postcode, country: $country, contactEmail: $email}, manager: {firstName: $firstName, lastName: $lastName, email: $email, jobTitle: \"\", telephone: \"\", password: $password}, recaptcha: \"\")\n}\n",
            "metadata": {}
        }
    } as any;
})();
(node as any).hash = 'cbb84b3bff02684da62cba9669b4f856';
export default node;
