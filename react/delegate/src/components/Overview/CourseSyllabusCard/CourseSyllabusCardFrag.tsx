import * as React from 'react';
import { createFragmentContainer, graphql } from 'react-relay';
import CourseSyllabusCard from './CourseSyllabusCard';
import { CourseSyllabusCardFrag_course } from './__generated__/CourseSyllabusCardFrag_course.graphql';

type Props = {
  course?: CourseSyllabusCardFrag_course;
  upTo?: String;
  completePercent?: number;
};

function SyllabusCard({ course, upTo, completePercent = 0 }: Props) {
  const syllabusSections = (course?.syllabus ?? []).map((courseElement) => {
    if (courseElement.type == 'module') {
      return {
        sections: [
          {
            name: courseElement.name ?? '',
            uuid: courseElement.uuid ?? '',
            complete: false
          },
          ...(courseElement?.syllabus ?? []).map((moduleItem) => {
            return {
              name: moduleItem.name ?? '',
              uuid: moduleItem.uuid ?? '',
              complete: false
            };
          })
        ]
      };
    } else {
      return {
        sections: [
          {
            name: courseElement.name ?? '',
            uuid: courseElement.uuid ?? '',
            complete: false
          }
        ]
      };
    }
  });

  var revSyllabus = syllabusSections.reverse();
  var found = false;
  for (const i in revSyllabus) {
    var revSections = syllabusSections[i].sections.reverse();
    for (const item in revSections) {
      if (revSections[item].uuid == upTo) {
        found = true;
      }
      if (found) {
        revSections[item] = { ...revSections[item], complete: true };
      }
    }
    revSyllabus[i].sections = revSections.reverse();
  }

  const syllabus = revSyllabus.reverse();
  const syllabusProp = {
    completePercentage: completePercent,
    modules: syllabus
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
