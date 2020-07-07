import * as React from 'react';
import CarouselImage, { Image } from './CarouselImage';
import { withKnobs, object } from '@storybook/addon-knobs';

export default {
  title: 'Misc/CarouselImage',
  decorators: [withKnobs]
};

const defaultCourse: Image = {
  url: 'https://picsum.photos/668/424',
  alt: 'Image'
};

export const plain = () => {
  const image: Image = object('Data', defaultCourse);
  const images = [1, 2, 3].map((item) => ({
    ...image,
    alt: `${image.alt} ${item}`
  }));
  return <CarouselImage images={images} />;
};
