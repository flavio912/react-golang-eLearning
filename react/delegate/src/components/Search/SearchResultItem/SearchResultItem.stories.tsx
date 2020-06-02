import * as React from 'react';
import { withKnobs, text, number } from '@storybook/addon-knobs';
import SearchResultItem from './SearchResultItem';

export default {
  title: 'Search/SearchResultItem',
  decorators: [withKnobs]
};

export const plain = () => {
  const course = {
    id: number('Id', 1),
    image: text('Image', 'https://www.gstatic.com/webp/gallery/1.jpg'),
    title: text('Title', 'Cargo Manager (CM) – VC, HS, XRY, EDS'),
    description: text(
      'Description',
      'This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…'
    )
  };
  return <SearchResultItem course={course} onClick={() => {}} />;
};
