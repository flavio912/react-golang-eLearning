import * as React from "react";
import HeaderMenu from "./HeaderMenu"; 
import { withKnobs, object, text } from "@storybook/addon-knobs";
import CourseCard from '../../../sharedComponents/Overview/CourseCard/CourseCard';
import SideMenu, { Tab } from '../SideMenu/SideMenu';

export default {
  title: "Menu/HeaderMenu",
  decorators: [withKnobs],
};

// Course card props
const defaultCourse = {
  type: "DANGEROUS GOODS AIR",
  colour: "#8C1CB4",
  url: require("../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png"),
  title: "Dangerous goods by air category 7",
  price: 60,
  description: "This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage",
  assigned: 40,
  expiring: 9,
  date: "MAR 3rd 2020",
  location: "TTC at Hilton T4"
};

// Side Menu props
const defaultTabs: Tab[] = [
    {id: 0, icon: "LeftNav_Icon_Dashboard", children: <CourseCard course={defaultCourse} filterColour="#8C1CB4" onClick={() => console.log('Pressed')} size="small" />},
    {id: 1, icon: "LeftNav_Icon_Delegates", children: <CourseCard course={defaultCourse} filterColour="#8C1CB4" onClick={() => console.log('Pressed')} size="large" />},
    {id: 2, icon: "LeftNav_Icon_Courses", children: <CourseCard course={defaultCourse} filterColour="#8C1CB4" onClick={() => console.log('Pressed')} size="small" />},
];

export const withContent = () => React.createElement(() => {
  // Header Menu knobs
  const logo: string = text("Logo", require("../../../assets/logo/ttc-logo.svg"));
  const name: string = text("Name", "Fred Ecceleston");
  const url: string = text("Profile Image", require("../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png"));
  const user = { name, url };

  // SideMenu state
  const [selected, setSelected] = React.useState(defaultTabs[0]);
  const tabs: Array<Tab> = object("Options", defaultTabs);
  return (
      <HeaderMenu
          logo={logo}
          user={user}
          onLogoClick={() => console.log('Logo Pressed')}
          onProfileClick={() => console.log('Profile Pressed')}
          children={
            <SideMenu
              selected={selected}
              tabs={tabs}
              onClick={(tab: Tab) => setSelected(tab)}
          />
          }
      />
  );
});

export const normal = () => {
  // Header Menu knobs
  const logo: string = text("Logo", require("../../../assets/logo/ttc-logo.svg"));
  const name: string = text("Name", "Fred Ecceleston");
  const url: string = text("Profile Image", require("../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png"));
  const user = { name, url };
  return (
    <HeaderMenu
      logo={logo}
      user={user}
      onLogoClick={() => console.log('Logo Pressed')}
      onProfileClick={() => console.log('Profile Pressed')}
  />
  );
};
