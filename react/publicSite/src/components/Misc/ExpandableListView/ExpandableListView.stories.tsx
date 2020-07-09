import * as React from 'react';
import ExpandableListView, { ExpandItemType } from './ExpandableListView';
import { withKnobs } from '@storybook/addon-knobs';

export default {
  title: 'Misc/ExpandableListView',
  decorators: [withKnobs],
};

const defaultCourse: ExpandItemType = {
  id: 3,
  title: 'Dangerous goods by air category 7',
  description:
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Ex ea commodo consequat.Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum iplorem ipsum Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut.',
  isExpanded: false,
};

export const plain = (): React.ReactElement => {
  const data = [1, 2, 3, 4, 5, 6].map(
    (index): ExpandItemType => ({ ...defaultCourse, id: index }),
  );
  return <ExpandableListView data={data} />;
};
