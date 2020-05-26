import * as React from 'react';
import PaymentSuccess from './PaymentSuccess';
import { withKnobs, number, text } from '@storybook/addon-knobs';

export default {
  title: 'Core/SlideModal/PaymentSuccess',
  decorators: [withKnobs]
};

export const normal = () => {
  return (
    <PaymentSuccess
      total={number('total', 165)}
      transactionId={text('transactionId', '34239rCD')}
    />
  );
};
