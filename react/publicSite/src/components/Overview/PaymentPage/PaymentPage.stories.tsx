import * as React from "react";
import { withKnobs } from "@storybook/addon-knobs";
import BillingCard, { BillingDetails } from "./BillingCard";
import PaymentHeader from "./PaymentHeader";
import OrderCard, { OrderItem } from "./OrderCard";
import PaymentCard from "./PaymentCard";

// Mock stripe
import { loadStripe } from '@stripe/stripe-js';
import { CardElement, Elements, useElements } from '@stripe/react-stripe-js';
// Modules for mocking router
//@ts-ignore
import { Resolver } from 'found-relay';
//@ts-ignore
import { MemoryProtocol } from 'farce';
import {
  createFarceRouter,
  createRender
} from 'found';
import environment from '../../../api/environment';

export default {
  title: "Overview/PaymentPage",
  decorators: [withKnobs],
};

const emptyBillingDetails: BillingDetails = {
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
};

const defaultOrderItems: OrderItem[] = [
    {
        id: '082739428373', name: 'Cargo Manager (CM) – VC, HS, XRY, EDS', quantity: 1, price: 200, imageURL: require("assets/SampleImage_ClassroomCoursesDetail_Feat.png")
    },
    {
        id: '082739428374', name: 'Cargo Aircraft Protection', quantity: 3, price: 55, imageURL: require("assets/SampleImage_ClassroomCoursesDetail_Feat.png")
    },
    {
        id: '082739428375', name: 'Cargo Manager Recurrent (CM) – VC, HS, XRY, EDS', quantity: 2, price: 25, imageURL: require("assets/SampleImage_ClassroomCoursesDetail_Feat.png")
    }
];

export const billing = () => {
    return <BillingCard billingDetails={emptyBillingDetails} />;
  };

export const order = () => {
    return <OrderCard orderItems={defaultOrderItems} />;
  };

export const header = () => {
    return <PaymentHeader
        text="Checkout Securely"
        imageURL={require('assets/StripePayment.svg')}
    />;
  };

const cardMock = React.createElement(() => {
    const elements = useElements();
    if (elements) {
        const card = elements.getElement(CardElement);
    }
  return (
    <PaymentCard total={336} />
  );
});

// Setup Stripe.js and the Elements provider
const stripePromise = loadStripe('pk_test_TYooMQauvdEDq54NiTphI7jx');

const stripeMock = () => {
    return (
      <Elements stripe={stripePromise}>
        {cardMock}
      </Elements>
    );
  }

// Mock Router
const StoryRouter = createFarceRouter({
    historyProtocol: new MemoryProtocol('/'),
    routeConfig: [{ path: '/', Component: stripeMock }],
    render: createRender({})
  })
  
export const card = () => (<StoryRouter resolver={new Resolver(environment)} />);