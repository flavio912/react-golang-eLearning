import * as React from 'react';
import CheckboxSingle from './CheckboxSingle';
import { withKnobs, text, boolean } from '@storybook/addon-knobs';

export default {
  title: 'Core/Input/CheckboxSingle',
  decorators: [withKnobs]
};

const Wrapper = () => {
  return (
    <CheckboxSingle
      label={text('Label text', 'This is a label')}
      defaultChecked={boolean('Default checked', false)}
    />
  );
};

export const normal = () => {
  return <Wrapper />;
};
