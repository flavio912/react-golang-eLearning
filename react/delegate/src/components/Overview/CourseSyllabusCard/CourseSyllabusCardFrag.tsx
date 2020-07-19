import * as React from 'react';
import { createFragmentContainer, graphql } from 'react-relay';
import CourseSyllabusCard from './CourseSyllabusCard';
import { CourseSyllabusCardFrag_course } from './__generated__/CourseSyllabusCardFrag_course.graphql';

type Props = {
  course?: CourseSyllabusCardFrag_course;
};

function SyllabusCard({ course }: Props) {
  const syllabusSections = (course?.syllabus ?? []).map((courseElement) => {
    if (courseElement.type == 'module') {
      return {
        sections: (courseElement?.syllabus ?? []).map((moduleItem) => {
          return {
            name: moduleItem.name ?? '',
            uuid: moduleItem.uuid ?? '',
            complete: moduleItem.complete ?? false
          };
        })
      };
    } else {
      return {
        sections: [
          {
            name: courseElement.name ?? '',
            uuid: courseElement.uuid ?? '',
            complete: courseElement.complete ?? false
          }
        ]
      };
    }
  });

  const syllabusProp = {
    completePercentage: 23,
    modules: syllabusSections
  };
  return <CourseSyllabusCard courseSyllabus={syllabusProp} />;
}

export default createFragmentContainer(SyllabusCard, {
  course: graphql`
    fragment CourseSyllabusCardFrag_course on Course {
      syllabus {
        name
        type
        uuid
        complete
        ... on Module {
          syllabus {
            name
            type
            uuid
            complete
          }
        }
      }
    }
  `
});
