import * as React from 'react';
import { withKnobs, text, number } from '@storybook/addon-knobs';
import SearchInput from './SearchInput';

export default {
  title: 'Search/SearchInput',
  decorators: [withKnobs]
};

export const plain = () => {
  return (
    <SearchInput
      onChange={() => {}}
      placeholder={text('placeholder', 'Search for Courses...')}
    />
  );
};
