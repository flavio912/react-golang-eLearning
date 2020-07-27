import * as React from 'react';
import CertInfo from './CertInfo';
import { withKnobs, text } from '@storybook/addon-knobs';

export default {
  title: 'Certificate/CertInfo',
  decorators: [withKnobs]
}

export const plain = () => {
  return (
    <CertInfo 
      certName={text('Certificate Name', 'General Security Awareness Training (GSAT)')}
      moduleDeliver={text('Module Deliver', 'Module Delivered: 1-20')}
      forEu={text('For EU', 'For EU 1998/2015: 11.2.2')}
      certNo={text('Certificate No.', '000000054321')}
      trainingDate={text('Date of Training', 'DD MMM YYYY')}
      expiryDate={text('Expiry Date', 'DD MMM YYYY')}
    />
  )
}