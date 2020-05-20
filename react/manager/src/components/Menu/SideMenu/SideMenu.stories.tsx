import * as React from "react";
import SideMenu, { Tab } from "./SideMenu";
import { withKnobs, object } from "@storybook/addon-knobs";

export default {
  title: "Menu/SideMenu",
  decorators: [withKnobs],
};

// Menu props
const defaultTabs: Array<Tab> = [
    {id: 0, icon: "LeftNav_Icon_Dashboard" },
    {id: 1, icon: "LeftNav_Icon_Delegates" },
    {id: 2, icon: "LeftNav_Icon_Courses" },
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