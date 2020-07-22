import { StructureElement } from './OnlineCourse/__generated__/OnlineCourse_myActiveCourse.graphql';

type CourseSyllabus = ReadonlyArray<{
  readonly name: string;
  readonly type: StructureElement;
  readonly uuid: string;
  readonly syllabus?: ReadonlyArray<{
    readonly name: string;
    readonly uuid: string;
    readonly type: StructureElement;
  }> | null;
}> | null;

export function goToNextURL(
  courseID: number,
  courseSyllabus: CourseSyllabus,
  currentUUID: string | undefined
) {
  if (courseSyllabus === null) {
    return '/app/courses';
  }

  if (currentUUID === undefined) {
    const type = courseSyllabus[0].type;
    const uuid = courseSyllabus[0].uuid;
    return `/app/courses/${courseID}/${type}/${uuid}`;
  }

  const flatSyllabus: {
    name: string;
    uuid: string;
    type: string;
  }[] = [];

  courseSyllabus.forEach((item, index) => {
    flatSyllabus.push(item);
    if (item.type === 'module') {
      item.syllabus?.forEach((modItem) => {
        flatSyllabus.push(modItem);
      });
    }
  });
  console.log('flat', flatSyllabus);
  console.log('n', currentUUID);
  var found = false;
  for (const item of flatSyllabus) {
    if (found) {
      return `/app/courses/${courseID}/${item.type}/${item.uuid}`;
    }
    if (item.uuid === currentUUID) {
      found = true;
    }
  }

  // Assumed finished
  return `/app/courses`;
}
