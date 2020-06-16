import * as React from 'react';
import HeaderMenu, { Tab } from './HeaderMenu';
import { withKnobs, object } from '@storybook/addon-knobs';
import CheckoutPopup, { BasketItem } from './CheckoutPopup';

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

const defaultItems: BasketItem[] = [
  {
    id: 0, name: "Cargo Manager Recurrent (CM) – VC, HS, XRY, EDS", price: 65.00, imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
  },
  {
    id: 1, name: "Cargo Aircraft Protection", price: 65.00, imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
  }
]

export const normal = () => React.createElement(() => {
  const [selected, setSelected] = React.useState(defaultTabs[0]);
  const tabs: Array<Tab> = object("Options", defaultTabs);

  return (
      <HeaderMenu
          selected={selected}
          tabs={tabs}
          onClick={(tab: Tab) => setSelected(tab)}
          onCheckout={() => console.log("Checkout")}
          basketItems={defaultItems}
      />
  );
});

export const popup = () => React.createElement(() => {
  return (
      <CheckoutPopup
          basketItems={defaultItems}
          onCheckout={() => console.log("Checkout")}
      />
  );
});
