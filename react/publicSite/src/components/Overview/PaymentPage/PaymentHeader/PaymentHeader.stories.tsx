import * as React from "react";
import PaymentHeader from "./PaymentHeader";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "Overview/PaymentPage/PaymentHeader",
  decorators: [withKnobs],
};

export const normal = () => {
    return <PaymentHeader
        text="Checkout Securely"
        imageURL={require('assets/StripePayment.svg')}
    />;
  };