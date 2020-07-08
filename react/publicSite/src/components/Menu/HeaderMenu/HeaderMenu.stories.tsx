import * as React from 'react';
import HeaderMenu from './HeaderMenu';
import { withKnobs, object } from '@storybook/addon-knobs';
import CheckoutPopup, { BasketItem } from './CheckoutPopup';
import { Tab } from './TabOption';

export default {
  title: 'Menu/HeaderMenu',
  decorators: [withKnobs]
};

// Menu props
const defaultTabs: Array<Tab> = [
  { id: 0, title: 'Features', options: [
    { title: 'For Companies', text: "We're training the finest minds in air, road and sea - are you on the list?", link: '/'},
    { title: 'For Individuals', text: "Do you need trainingsolutions for your next role in Transport Compliance?", link: '/'}
  ]},
  { id: 1, title: 'Courses', options: [
    { title: 'Aviation Security', text: "Training courses specifically designed for those who work in Aviation Security", link: '/'},
    { title: 'Dangerous Goods', text: "Courses for both air and road, all in accordance with CAA Regulations", link: '/'},
    { title: 'Health & Safety', text: "All our courses can be taken online in conjunction withyour internal policies", link: '/'},
    { title: 'Classroom Courses', text: "All classroom courses are delivered in London at our purpose built facility", link: '/'}
  ]},
  { id: 2, title: 'Resources', options: [
    { title: 'For Companies', text: "We're training the finest minds in air, road and sea - are you on the list?", link: '/'},
    { title: 'For Individuals', text: "Do you need trainingsolutions for your next role in Transport Compliance?", link: '/'}
  ]},
  { id: 3, title: 'Consultancy', link: '/' }
];

const defaultItems: BasketItem[] = [
  {
    id: 0,
    name: 'Cargo Manager Recurrent (CM) – VC, HS, XRY, EDS',
    price: 65.0,
    imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
  },
  {
    id: 1,
    name: 'Cargo Aircraft Protection',
    price: 65.0,
    imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
  }
];

export const normal = () =>
  React.createElement(() => {
    const [selected, setSelected] = React.useState<Tab | undefined>();
    const tabs: Array<Tab> = object('Options', defaultTabs);

    return (
      <HeaderMenu
        selected={selected}
        setSelected={(tab?: Tab) => setSelected(tab)}
        tabs={tabs}
        onLogoClick={() => console.log('Logo')}
        onClick={(link: string) => console.log(link)}
        onCheckout={() => console.log('Checkout')}
        basketItems={defaultItems}
      />
    );
  });

export const popup = () =>
  React.createElement(() => {
    return (
      <CheckoutPopup
        showPopup={true}
        basketItems={defaultItems}
        onCheckout={() => console.log('Checkout')}
      />
    );
  });
