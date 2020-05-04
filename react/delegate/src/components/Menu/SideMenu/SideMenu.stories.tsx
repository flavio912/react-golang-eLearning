import * as React from "react";
import SideMenu, { Tab } from "./SideMenu";
import { withKnobs, text, object } from "@storybook/addon-knobs";

export default {
  title: "Menu/SideMenu",
  decorators: [withKnobs],
};

// Menu props
const defaultTabs: Array<Tab> = [
    {id: 0, icon: "LeftNav_Icon_Dashboard", title: "Dashboard" },
    {id: 1, icon: "LeftNav_Icon_Courses", title: "Online Courses" },
];

export const normal = () => React.createElement(() => {
    const [selected, setSelected] = React.useState(defaultTabs[0]);
    const tabs: Array<Tab> = object("Options", defaultTabs);
    const logo: string = text("Logo", require("../../../assets/logo/ttc-logo.svg"));

    return (
        <SideMenu
            logo={logo}
            selected={selected}
            tabs={tabs}
            onClick={(tab: Tab) => setSelected(tab)}
        />
    );
});