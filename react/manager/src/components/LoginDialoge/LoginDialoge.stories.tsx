import * as React from 'react';
import LoginDialoge from './LoginDialoge';
import { withKnobs } from '@storybook/addon-knobs';

export default {
  title: 'Login/Login Dialoge',
  decorators: [withKnobs]
}

export const normal = () => {
  return <LoginDialoge/>
}