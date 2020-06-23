import * as React from "react";
import PaymentCard from "./PaymentCard";
import { withKnobs } from "@storybook/addon-knobs";

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
import environment from '../../../../api/environment';

export default {
  title: "Overview/PaymentPage/PaymentCard",
  decorators: [withKnobs],
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
  
export const normal = () => (<StoryRouter resolver={new Resolver(environment)} />);