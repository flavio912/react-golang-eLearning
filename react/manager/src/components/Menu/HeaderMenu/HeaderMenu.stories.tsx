import * as React from "react";
import HeaderMenu from "./HeaderMenu";
import { withKnobs, object, text } from "@storybook/addon-knobs";
import SideMenu, { Tab } from "../SideMenu/SideMenu";

export default {
  title: "Menu/HeaderMenu",
  decorators: [withKnobs],
};

// Side Menu props
const defaultTabs: Tab[] = [
  { id: 0, icon: "LeftNav_Icon_Dashboard" },
  { id: 1, icon: "LeftNav_Icon_Delegates" },
  { id: 2, icon: "LeftNav_Icon_Courses" },
];

export const withContent = () =>
  React.createElement(() => {
    // Header Menu knobs
    const logo: string = text(
      "Logo",
      require("../../../assets/logo/ttc-logo.svg")
    );
    const name: string = text("Name", "Fred Ecceleston");
    const url: string = text(
      "Profile Image",
      require("../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png")
    );
    const user = { name, url };

    // SideMenu state
    const [selected, setSelected] = React.useState(defaultTabs[0]);
    const tabs: Array<Tab> = object("Options", defaultTabs);
    return (
      <HeaderMenu
        logo={logo}
        user={user}
        onLogoClick={() => console.log("Logo Pressed")}
        onProfileClick={() => console.log("Profile Pressed")}
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
  const logo: string = text(
    "Logo",
    require("../../../assets/logo/ttc-logo.svg")
  );
  const name: string = text("Name", "Fred Ecceleston");
  const url: string = text(
    "Profile Image",
    require("../../../assets/SampleImage_ClassroomCoursesDetail_Feat.png")
  );
  const user = { name, url };
  return (
    <HeaderMenu
      logo={logo}
      user={user}
      onLogoClick={() => console.log("Logo Pressed")}
      onProfileClick={() => console.log("Profile Pressed")}
    />
  );
};
