import * as React from 'react';
import PaymentForm from './PaymentForm';
import { withKnobs } from '@storybook/addon-knobs';

export default {
  title: 'Core/SlideModal/PaymentForm',
  decorators: [withKnobs]
};

export const normal = () => {
  return <PaymentForm onPurchase={() => {}} />;
};
