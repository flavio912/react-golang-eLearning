import * as React from 'react';
import CarouselCourse from './CarouselCourse';
import { withKnobs, object } from '@storybook/addon-knobs';
import { Course } from 'sharedComponents/Overview/CourseCard';

export default {
  title: 'Misc/CarouselCourse',
  decorators: [withKnobs]
};

const defaultCourse: Course = {
  id: 3,
  type: 'DANGEROUS GOODS AIR',
  colour: '#8C1CB4',
  url: require('../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png'),
  title: 'Dangerous goods by air category 7',
  price: 60,
  description:
    'This course is for those involved in the handling, storage and loading of cargo or mail and baggage.',
  assigned: 40,
  expiring: 9,
  date: 'MAR 3rd 2020',
  location: 'TTC at Hilton T4',
  modules: 16,
  lessons: 144,
  videoTime: 4
};

export const plain = () => {
  const courseData: Course = object('Data', defaultCourse);
  const courses = [1, 2, 3, 4, 5, 6].map((item) => courseData);
  return <CarouselCourse courses={courses} />;
};
