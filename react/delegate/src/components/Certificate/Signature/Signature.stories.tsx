import * as React from 'react';
import Signature from './Signature';
import { withKnobs, number } from '@storybook/addon-knobs';

export default {
  title: 'Certificate/Signature',
  decorators: [withKnobs]
}

export const plain = () => {
  return (
    <Signature 
      width={number('Width', 500)} 
      height={number('Height', 70)}
    />
  )
}