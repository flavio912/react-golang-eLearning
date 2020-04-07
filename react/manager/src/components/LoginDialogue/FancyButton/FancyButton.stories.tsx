import * as React from 'react';
import FancyButton from './FancyButton';
import { withKnobs, text } from '@storybook/addon-knobs';

export default {
  title: 'Login/Fancy',
  decorators: [withKnobs]
}

export const button = () => {
  return (
    <FancyButton
      text={text('Text', 'Login to TTC')}
    />
  )
}