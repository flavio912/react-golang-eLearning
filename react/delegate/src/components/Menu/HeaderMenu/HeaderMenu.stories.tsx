import * as React from 'react';
import HeaderMenu from './HeaderMenu';
import { withKnobs, text } from '@storybook/addon-knobs';

export default {
  title: 'Menu/HeaderMenu',
  decorators: [withKnobs]
};

export const normal = () => {
  // Header Menu knobs
  const name: string = text('Name', 'Fred Ecceleston');
  const url: string = text(
    'Profile Image',
    require('../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png')
  );
  const user = { name, url };
  return (
    <HeaderMenu
      user={user}
      onProfileClick={() => console.log('Profile Pressed')}
    />
  );
};
