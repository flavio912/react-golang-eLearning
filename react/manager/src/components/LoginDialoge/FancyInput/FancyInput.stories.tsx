import * as React from 'react';
import FancyInput from './FancyInput';
import { withKnobs } from '@storybook/addon-knobs';

export default {
  title: 'Login/Fancy Input',
  decorators: [withKnobs]
}

export const normal = () => {
  return <FancyInput/>
}