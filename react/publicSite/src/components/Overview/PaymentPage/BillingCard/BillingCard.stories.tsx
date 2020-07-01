import * as React from "react";
import BillingCard, { BillingDetails } from "./BillingCard";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "Overview/PaymentPage/BillingCard",
  decorators: [withKnobs],
};

const emptyBillingDetails: BillingDetails = {
    addressOne: "",
    addressTwo: "",
    city: "",
    postcode: "",
    country: "",
    contact: false,
};

export const normal = () => {
    return <BillingCard billingDetails={emptyBillingDetails} />;
  };