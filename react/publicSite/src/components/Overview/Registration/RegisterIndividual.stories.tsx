import * as React from 'react';
import { withKnobs, text } from '@storybook/addon-knobs';
import RegisterIndividual from './RegisterIndividual';

export default {
  title: 'Overview/Registration/RegisterIndividual',
  decorators: [withKnobs]
};

export const normal = () => {
  return <RegisterIndividual onSubmit={() => {}} />;
};
