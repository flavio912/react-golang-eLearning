import * as React from 'react';
import { withKnobs, text } from '@storybook/addon-knobs';
import RegisterCompany from './RegisterCompany';

export default {
  title: 'Overview/Registration/RegisterCompany',
  decorators: [withKnobs]
};

export const normal = () => {
  return <RegisterCompany onSubmit={() => {}} />;
};
