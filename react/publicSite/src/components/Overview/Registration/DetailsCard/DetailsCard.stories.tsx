import * as React from "react";
import DetailsCard, { IndividualDetails, CompanyDetails } from "./DetailsCard";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "Overview/Registration/DetailsCard",
  decorators: [withKnobs],
};

const emptyIndividualDetails: IndividualDetails = {
    firstName: "",
    lastName: "",
    email: "",
    number: "",
    role: "",
    contact: false,
};

const emptyCompanyDetails: CompanyDetails = {
    firstName: "",
    lastName: "",
    companyName: "",
    companyEmail: "",
    number: "",
    companyType: "",
    role: "",
    contact: false,
};

export const normal = React.createElement(() => {
    const title: string = text('Title', 'Register your team today');
    const subtitle: string = text('Subtitle', 'If you’re looking to get access to the finest level of compliance training register with TTC Hub below');
   
    // Internal state
    const [individual, setIndividual] = React.useState(false);
    const [company, setCompany] = React.useState(false);
   
    return <DetailsCard
    title={title}
    subtitle={subtitle}
    onRegister={() => console.log('Register')}
    onBook={() => console.log('Book')}
    onIndividual={() => setIndividual(true)}
    onCompany={() => setCompany(true)}
    individual={individual}
    company={company}
    individualDetails={emptyIndividualDetails}
    companyDetails={emptyCompanyDetails}
  />;
});