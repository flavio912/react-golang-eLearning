import * as React from "react";
import SideMenu, { Tab } from "./SideMenu";
import { withKnobs, object } from "@storybook/addon-knobs";
import CourseCard from '../../Overview/CourseCard/CourseCard';

export default {
  title: "Core/SideMenu",
  decorators: [withKnobs],
};

// Course card props
const defaultCourse = {
  type: "DANGEROUS GOODS AIR",
  url: require("../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png"),
  title: "Dangerous goods by air category 7",
  price: 60,
  description: "This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage",
  assigned: 40,
  expiring: 9,
  date: "MAR 3rd 2020",
  location: "TTC at Hilton T4"
};

// Menu props
const defaultTabs: Array<Tab> = [
    {id: 0, icon: "LeftNav_Icon_Dashboard", children: <CourseCard course={defaultCourse} color="#8C1CB4" onClick={() => console.log('Pressed')} size="small" />},
    {id: 1, icon: "LeftNav_Icon_Delegates", children: <CourseCard course={defaultCourse} color="#8C1CB4" onClick={() => console.log('Pressed')} size="large" />},
    {id: 2, icon: "LeftNav_Icon_Courses", children: <CourseCard course={defaultCourse} color="#8C1CB4" onClick={() => console.log('Pressed')} size="small" />},
];

export const normal = () => React.createElement(() => {
    const [selected, setSelected] = React.useState(defaultTabs[0]);
    const tabs: Array<Tab> = object("Options", defaultTabs);

    return (
        <SideMenu
            selected={selected}
            tabs={tabs}
            onClick={(tab: Tab) => setSelected(tab)}
        />
    );
});