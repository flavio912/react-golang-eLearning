import * as React from 'react';
import { withKnobs, text } from '@storybook/addon-knobs';
import AccountTypeCard from './AccountTypeCard';

export default {
  title: 'Overview/Registration/AccountTypeCard',
  decorators: [withKnobs]
};

export const normal = () => {
  return <AccountTypeCard onIndividual={() => {}} onCompany={() => {}} />;
};
