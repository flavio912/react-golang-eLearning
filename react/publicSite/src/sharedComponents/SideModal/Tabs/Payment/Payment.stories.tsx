import * as React from 'react';
import Payment from './Payment';
import { withKnobs, number, text } from '@storybook/addon-knobs';

export default {
  title: 'Core/Modals/SlideModal/Payment',
  decorators: [withKnobs]
};

export const normal = () => {
  return (
    <Payment
      courses={[
        {
          id: number('id', 1),
          name: text(
            'Course name 1',
            'Cargo Operative Screener (COS) Recurrent – VC, HS, XRY, ETD'
          ),
          price: number('price', 55),
          sku: text('sku', '082739428374')
        }
      ]}
      userUUIDs={['123']}
      isContract={false}
      onPurchase={() => false}
      onSuccess={() => {}}
      onError={() => {}}
    />
  );
};
