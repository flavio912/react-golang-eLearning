import * as React from 'react';
import { withKnobs, select, text } from '@storybook/addon-knobs';
import LabelledCard from './LabelledCard';

export default {
  title: 'Core/Cards/LabelledCard',
  decorators: [withKnobs]
};

export const plain = () => {
  return (
    <LabelledCard
      label={text('Card label', 'Example label')}
      labelBackground={text('BackgroundColor', 'black')}
    >
      {text('Card content', 'A card with a fancy label')}
    </LabelledCard>
  );
};
