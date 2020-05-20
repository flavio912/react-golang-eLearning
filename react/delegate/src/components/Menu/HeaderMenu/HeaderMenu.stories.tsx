import * as React from "react";
import HeaderMenu from "./HeaderMenu";
import { withKnobs, text } from "@storybook/addon-knobs";
import CourseCard from "../../Overview/CourseCard/CourseCard";

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
  description:
    "This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage",
  assigned: 40,
  expiring: 9,
  date: "MAR 3rd 2020",
  location: "TTC at Hilton T4",
};

export const withContent = () =>
  React.createElement(() => {
    // Header Menu knobs
    const name: string = text("Name", "Fred Ecceleston");
    const url: string = text(
      "Profile Image",
      require("../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png")
    );
    const user = { name, url };
    return (
      <HeaderMenu
        user={user}
        onProfileClick={() => console.log("Profile Pressed")}
      />
    );
  });

export const normal = () => {
  // Header Menu knobs
  const name: string = text("Name", "Fred Ecceleston");
  const url: string = text(
    "Profile Image",
    require("../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png")
  );
  const user = { name, url };
  return (
    <HeaderMenu
      user={user}
      onProfileClick={() => console.log("Profile Pressed")}
    />
  );
};
