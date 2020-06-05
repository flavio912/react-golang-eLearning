import * as React from 'react';
import CircleBorder, { User, BorderType } from './CircleBorder';
import { withKnobs, text, number, select } from '@storybook/addon-knobs';

export default {
  title: 'Core/CircleBorder',
  decorators: [withKnobs]
};

export const normal = () => {
  const name: string = text('Name', 'Fred Ecceleston');
  const url: string = text(
    'Profile Image',
    require('../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png')
  );
  const user: User = { name, url };
  const size: number = number('Size', 44);
  const fontSize: number = number('fontSize', 18);
  const type: BorderType = select('Border Type', ['fancy', 'plain'], 'fancy');

  return (
    <CircleBorder
      user={user}
      size={size}
      fontSize={fontSize}
      borderType={type}
    />
  );
};
