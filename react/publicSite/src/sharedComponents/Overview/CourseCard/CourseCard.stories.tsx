import * as React from 'react';
import CourseCard, { Course, SizeOptions, Completion } from './CourseCard';
import { withKnobs, select, object } from '@storybook/addon-knobs';

export default {
  title: 'Overview/CourseCard',
  decorators: [withKnobs]
};

const sizeOptions: SizeOptions[] = ['small', 'large'];

const defaultCourse: Course = {
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
  video_time: 4
};

const defaultProgress: Completion = {
  total: 100,
  complete: 40
};

export const plain = () => {
  const size: SizeOptions = select('Size', sizeOptions, 'small');
  const progress: Completion = object('Progress Bar', defaultProgress);
  const courseData: Course = object('Data', defaultCourse);
  return (
    <CourseCard
      course={courseData}
      onClick={() => console.log('Pressed')}
      size={size}
      progress={progress}
    />
  );
};
