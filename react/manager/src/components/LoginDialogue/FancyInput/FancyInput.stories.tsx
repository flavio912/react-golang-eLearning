import * as React from 'react';
import FancyInput from './FancyInput';
import { withKnobs, text, color } from '@storybook/addon-knobs';

export default {
  title: 'Login/Fancy Input',
  decorators: [withKnobs]
}

export const normal = () => {
  return (
    <FancyInput
      label={text('Label', 'Email')}
      labelColor={color('Label Color', 'black')}
      placeholder={text('Placeholder', 'joe@joe.com')}
    />
  )
}