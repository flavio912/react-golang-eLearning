import * as React from 'react';
import { withKnobs, text, boolean, select } from '@storybook/addon-knobs';
import FlatCard from './FlatCard';

export default {
  title: 'FlatCard',
  decorators: [withKnobs]
};

export const plain = () => {
  return (
    <FlatCard
      className={text('className', 'flat-card')}
      children={<div>This is a test</div>}
      backgroundColor={'white'}
      shadow={boolean('Shadow', false)}
      padding={select('Padding', ['small', 'medium', 'large', 'none'], 'small')}
      style={{
        alignItems: 'center',
        justifyContent: 'center'
      }}
    />
  );
};
