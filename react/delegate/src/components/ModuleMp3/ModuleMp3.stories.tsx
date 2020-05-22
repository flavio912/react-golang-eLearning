import * as React from 'react';
import { withKnobs, text } from '@storybook/addon-knobs';
import ModuleMp3 from './ModuleMp3';

export default {
  title: 'ModuleMp3',
  decorators: [withKnobs]
};

export const plain = () => {
  return (
    <ModuleMp3
      className={text('className', 'module-mp3')}
      module={{
        name: text('name', 'Module 1'),
        subTitle: text('subtitle', 'General Philosophy'),
        mp3Url: text(
          'mp3Url',
          'https://storage.googleapis.com/media-session/elephants-dream/the-wires.mp3'
        )
      }}
    />
  );
};
