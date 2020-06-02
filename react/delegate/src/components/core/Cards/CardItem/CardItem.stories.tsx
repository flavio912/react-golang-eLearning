import * as React from 'react';
import { withKnobs, text } from '@storybook/addon-knobs';
import CardItem from './CardItem';

export default {
  title: 'Core/Cards/CardItem',
  decorators: [withKnobs]
};

export const plain = () => {
  return (
    <div>
      <CardItem
        title={text('Title', 'All Coursees')}
        description={text(
          'Description',
          'Find all of the online courses you have been enrolled on here'
        )}
        buttonProps={{
          title: text('ButtonTitle', 'View Courses'),
          onClick: () => {}
        }}
      />
    </div>
  );
};
