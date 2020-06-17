import * as React from "react";
import PaymentPage from "./PaymentPage";
import { withKnobs } from "@storybook/addon-knobs";
import BillingCard, { BillingDetails } from "./BillingCard";
import PaymentHeader from "./PaymentHeader";

export default {
  title: "Overview/PaymentPage",
  decorators: [withKnobs],
};

const empyBillingDetails: BillingDetails = {
    email: "",
    companyName: "",
    phoneNumber: "",
    firstName: "",
    lastName: "",
    adressOne: "",
    adressTwo: "",
    city: "",
    postcode: "",
    country: "",
    contact: false,
}

export const billing = () => {
    return <BillingCard billingDetails={empyBillingDetails} />;
  };

export const header = () => {
    return <PaymentHeader
        text="Checkout Securely"
        imageURL={require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')}
    />;
  };

export const page = () => {
  return <PaymentPage />;
};