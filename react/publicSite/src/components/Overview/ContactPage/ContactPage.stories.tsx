import * as React from "react";
import ContactPage, { ContactDetails } from "./ContactPage";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "Overview/ContactPage",
  decorators: [withKnobs],
};

const emptyContactDetails: ContactDetails = {
    firstName: "",
    lastName: "",
    companyName: "",
    email: "",
    type: "",
    extra: "",
}

export const normal = () => {
  return <ContactPage
    contactDetails={emptyContactDetails}
    onHelp={() => console.log('Help')}
    onSubmit={() => console.log('Submit')}
  />;
};