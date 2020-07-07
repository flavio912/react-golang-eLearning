import * as React from 'react';
import FancyBorder from './FancyBorder';
import { withKnobs } from '@storybook/addon-knobs';

export default {
  title: 'Certificate/Fancy',
  decorators: [withKnobs]
}

export const border = () => {
  return (
    <FancyBorder children={<div>This is a test</div>} />
  )
}