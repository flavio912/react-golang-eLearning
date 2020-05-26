import * as React from "react";
import CourseSyllabusCard, { CourseSyllabus } from "./CourseSyllabusCard";
import {
  withKnobs,
	object,
} from "@storybook/addon-knobs";

export default {
	title: "Overview/Course Syllabus",
	decorators: [withKnobs],
}
const defaultSyllabus: CourseSyllabus = {
	completePercentage: 24,
	modules: [
		{
			sections: [
				{
					name: "Lesson 1-1",
					uuid: "00000-0000-00000-0000",
					complete: false,
				},
				{
					name: "Lesson 1-2",
					uuid: "00000-0000-00000-0000",
					complete: false,
				}
			]
		},
		{
			sections: [
				{
					name: "Lesson 2-1",
					uuid: "00000-0000-00000-0000",
					complete: false,
				},
				{
					name: "Lesson 2-2",
					uuid: "00000-0000-00000-0000",
					complete: false,
				}
			]
		},
		{
			sections: [
				{
					name: "Lesson 3-1",
					uuid: "00000-0000-00000-0000",
					complete: false,
				},				
				{
					name: "Lesson 3-2",
					uuid: "00000-0000-00000-0000",
					complete: false,
				}
			]
		}
	]
}


export const normal = () => {
  const syllabus = object("CourseSyllabus", {...defaultSyllabus});
	return (
		<CourseSyllabusCard courseSyllabus={syllabus} />
	)
}