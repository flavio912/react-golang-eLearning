import * as React from "react";
import AccountCard, { AccountDetails } from "./AccountCard";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "Overview/PaymentPage/AccountCard",
  decorators: [withKnobs],
};

const emptyBillingDetails: AccountDetails = {
    firstName: "",
    lastName: "",
    emailAddress: "",
    companyName: "",
    phoneNumber: ""
};

export const normal = () => {
    return <AccountCard accountDetails={emptyBillingDetails} />;
  };