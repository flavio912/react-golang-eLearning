import * as React from "react";
import CourseSearch, { Tab } from "./CourseSearch";
import { withKnobs, text, object } from "@storybook/addon-knobs";
import CourseItem, { CourseProps } from "./CourseItem";

export default {
  title: "CourseSearch",
  decorators: [withKnobs],
};

const defaultTabs: Tab[] = [
    {
        name: 'All Courses', value: ''
    },
    {
        name: 'Regulated Agents', value: 'Regulated Agent'
    },
    {
        name: 'Known Consignor', value: 'Known Consignor'
    },
    {
        name: 'GSAT', value: 'GSAT'
    }
];

const defaultCourseItem: CourseProps = {
    title: "Cargo Manager (CM) – VC, HS, XRY, EDS",
    description: "We have over 100 Aviation Security Courses lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
    price: "£310.00",
    type: "Regulated Agents",
    colour: "#8C1CB4",
    imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png'),
    viewCourse: () => console.log("View Course"),
    addToBasket: () => console.log("Add to Basket")
}

const defaultCourseItems: CourseProps[] = [
    defaultCourseItem,defaultCourseItem,defaultCourseItem,defaultCourseItem,defaultCourseItem,defaultCourseItem,defaultCourseItem,defaultCourseItem,defaultCourseItem,defaultCourseItem
]

export const search = () => {
    const tabs: Tab[] = object("Tabs", defaultTabs);
  return <CourseSearch 
        tabs={tabs}
        courses={defaultCourseItems}
    />;
};

export const item = () => {
    const title: string = text("Title", "Cargo Manager (CM) – VC, HS, XRY, EDS");
    const description: string = text("Description", "We have over 100 Aviation Security Courses lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua");
    const price: string = text("Price", "£310.00");
    const type: string = text("Type", "Regulated Agents");
    const colour: string = text("Colour", "#8C1CB4");
    return <CourseItem 
        title={title}
        description={description}
        price={price}
        colour={colour}
        type={type}
        imageURL={require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')}
        viewCourse={() => console.log("View Course")}
        addToBasket={() => console.log("Add to Basket")}
  />
};
