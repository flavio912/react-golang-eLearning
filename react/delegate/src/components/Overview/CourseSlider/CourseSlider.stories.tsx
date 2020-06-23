import * as React from 'react';
import CourseSlider from './CourseSlider';
import { withKnobs, object, number } from '@storybook/addon-knobs';
import { Course as CourseCardProps } from 'sharedComponents/Overview/CourseCard/CourseCard';

export default {
  title: 'Overview/CourseSlider',
  decorators: [withKnobs]
};

const defaultCourse: CourseCardProps = {
  type: 'DANGEROUS GOODS AIR',
  colour: '#8C1CB4',
  url: require('../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png'),
  title: 'Dangerous goods by air category',
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

export const plain = () => {
  const courseData = object('Data', {
    ...defaultCourse,
    progress: number('Progress', 40)
  });
  const courses = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11].map((item) => {
    return {
      id: item,
      ...courseData,
      title: `${courseData.title} ${item}`
    };
  });
  return <CourseSlider courses={courses} slidesToShow={4} />;
};
