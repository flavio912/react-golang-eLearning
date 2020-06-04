import * as React from 'react';
import HeaderMenu, { Tab } from './HeaderMenu';
import { withKnobs, object } from '@storybook/addon-knobs';

export default {
  title: 'Menu/HeaderMenu',
  decorators: [withKnobs]
};

// Menu props
const defaultTabs: Array<Tab> = [
  {id: 0, title: "Features", options: ["Some", "Different", "Options"] },
  {id: 1, title: "Courses", options: ["Some", "Different", "Options"]},
  {id: 2, title: "Resources", options: ["Some", "Different", "Options"]},
  {id: 3, title: "Consultancy"},
];

export const normal = () => React.createElement(() => {
  const [selected, setSelected] = React.useState(defaultTabs[0]);
  const tabs: Array<Tab> = object("Options", defaultTabs);

  return (
      <HeaderMenu
          selected={selected}
          tabs={tabs}
          onClick={(tab: Tab) => setSelected(tab)}
          basketItems={2}
      />
  );
});
