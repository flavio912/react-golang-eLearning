import * as React from 'react';
import ActivityName from './ActivityName';
import { withKnobs, text } from '@storybook/addon-knobs';

export default {
  title: 'Delegate/ActivityName',
  decorators: [withKnobs]
};

export const normal = () => {
  return (
    <ActivityName
      avatar={text('Text', 'https://picsum.photos/id/1/200/300')}
      title={text(
        'Text',
        'Bruce failed the Dangerous Goods by Road Awareness Course'
      )}
    />
  );
};
