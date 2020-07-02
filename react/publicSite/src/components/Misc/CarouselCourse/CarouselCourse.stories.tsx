import * as React from 'react';
import CarouselCourse from './CarouselCourse';
import { withKnobs, object, text } from '@storybook/addon-knobs';
import { Course } from 'sharedComponents/Overview/CourseCard';
import CarouselWithDemo from './CarouselWithDemo';

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

export const withDemo = () => {
  const heading: string = text("Heading", "Explore our popular courses");
  const description: string = text("Description", "It’s time to remove the headache for you and your team, with TTC you could be logged in and learning in 24 hours.");
  const courseData: Course = object('Data', defaultCourse);
  const courses = [1, 2, 3, 4, 5, 6].map((item) => courseData);
  return <CarouselWithDemo heading={heading} description={description} courses={courses} />;
};
