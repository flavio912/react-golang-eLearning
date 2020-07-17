import * as React from 'react';
import { withKnobs, text, number } from '@storybook/addon-knobs';
import SearchResults, { Result, Page } from './SearchResults';

export default {
  title: 'Search/SearchResults',
  decorators: [withKnobs]
};

const results = [
  {
    id: number('Id 1', 1),
    title: text('Title 1', 'Cargo Manager (CM) – VC, HS, XRY, EDS'),
    image: text('Image 1', 'https://www.gstatic.com/webp/gallery/1.jpg'),
    description: text(
      'Description 1',
      'This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…'
    )
  },
  {
    id: number('Id 2', 2),
    title: text('Title 2', 'Cargo Manager (CM) – VC, HS, XRY, EDS'),
    image: text('Image 2', 'https://www.gstatic.com/webp/gallery/1.jpg'),
    description: text(
      'Description 2',
      'This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…'
    )
  },
  {
    id: number('Id 3', 3),
    title: text('Title 3', 'Cargo Manager (CM) – VC, HS, XRY, EDS'),
    image: text('Image 3', 'https://www.gstatic.com/webp/gallery/1.jpg'),
    description: text(
      'Description 3',
      'This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…'
    )
  },
  {
    id: number('Id 4', 4),
    title: text('Title 4', 'Cargo Manager (CM) – VC, HS, XRY, EDS'),
    image: text('Image 4', 'https://www.gstatic.com/webp/gallery/1.jpg'),
    description: text(
      'Description 4',
      'This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…'
    )
  }
];

const result = {
  resultItems: results,
  pageInfo: {
    total: 10,
    offset: 1,
    limit: 10,
    given: 10
  } 
};

const searchFunc = async (query : string, page: Page) => {
  return await new Promise<Result>((resolve) => {
    setTimeout(() => resolve(result), 500)
  });
}

export const plain = () => {
  return <SearchResults searchFunction={searchFunc} />;
};
