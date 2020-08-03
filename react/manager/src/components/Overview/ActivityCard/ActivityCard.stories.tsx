import * as React from 'react';
import ActvityCard, { Statistic } from './ActivityCard';
import { PaddingOptions } from '../../../sharedComponents/core/Cards/Card/index';
import { withKnobs, select, text, object, array } from '@storybook/addon-knobs';

export default {
  title: 'Overview/ActvityCard',
  decorators: [withKnobs]
};

const paddingOptions: PaddingOptions[] = ['none', 'small', 'medium', 'large'];

const defaultInfo: any = [
  { name: 'James Green', course: 'Cargo Aircraft Piloting', time: new Date() },
  { name: 'Bryony Turner', course: 'Aviation Security', time: new Date() },
  { name: 'Robert Johnson', course: 'Cargo Management', time: new Date() },
  { name: 'James Green', course: 'Cargo Aircraft Piloting', time: new Date() },
  { name: 'Bryony Turner', course: 'Aviation Security', time: new Date() },
  { name: 'Robert Johnson', course: 'Cargo Management', time: new Date() },
  { name: 'James Green', course: 'Cargo Aircraft Piloting', time: new Date() },
  { name: 'Bryony Turner', course: 'Aviation Security', time: new Date() },
  { name: 'Robert Johnson', course: 'Cargo Management', time: new Date() }
];

const defaultHeaders = ['This Month', 'All Time'];

const defaultData: Statistic = {
  outerRing: {
    name: 'Active',
    value: 154
  },
  innerRing: {
    name: 'Inactive',
    value: 64
  }
};

// TODO
// export const plain = () => {
//   const padding: PaddingOptions = select('Padding', paddingOptions, 'medium');
//   const leftHeadingText: string = text('Left Heading', '');
//   const rightHeadingText: string = text('Right Heading', '');
//   const optionsValues: string[] = array('Header Options', [...defaultHeaders]);
//   const dataValues: Statistic = object('Data Values', defaultData);
//   const updateData: any = object('Updates', [...defaultInfo]);
//   return (
//     <ActvityCard
//       leftHeading={leftHeadingText || 'Delegates activity'}
//       rightHeading={rightHeadingText || 'Recent updates'}
//       options={optionsValues || ['This Month', 'All Time']}
//       updates={{}}
//       data={dataValues}
//       onClick={() => console.log('Pressed')}
//       padding={padding}
//     />
//   );
// };
