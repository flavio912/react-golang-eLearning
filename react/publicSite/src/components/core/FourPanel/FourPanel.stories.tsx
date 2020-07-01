import * as React from "react";
import FourPanel, { Panel } from "./FourPanel";
import { withKnobs, boolean, number } from "@storybook/addon-knobs";

export default {
  title: "Core/FourPanel",
  decorators: [withKnobs],
};

const defaultIconPanels: Panel[] = [
    {
        title: "Addressing annual reports", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqu Lorem ipsum dolor sit amet, consectetur adipisci Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod", iconName: "TTC_Logo_Icon"
    },
    {
        title: "Cutting compliance costs", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqu Lorem ipsum dolor sit amet, consectetur adipisci Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod", iconName: "TTC_Logo_Icon"
    },
    {
        title: "Supporting security managers", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqu Lorem ipsum dolor sit amet, consectetur adipisci Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod", iconName: "TTC_Logo_Icon"
    },
    {
        title: "Continually evolving processes", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqu Lorem ipsum dolor sit amet, consectetur adipisci Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod", iconName: "TTC_Logo_Icon"
    }
];

const defaultImagePanels: Panel[] = [
    {
        title: "Start with this thing", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim.", imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
    },
    {
        title: "Then do this thing", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim.", imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
    },
    {
        title: "Then give this a shot", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim.", imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
    },
    {
        title: "Finish with some of that", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim.", imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
    }
];

export const normal = () => {
    const showImagePanels: boolean = boolean("Show Image Panels", false);
  return <FourPanel
    panels={showImagePanels ? defaultImagePanels : defaultIconPanels}
    noBorders={showImagePanels}
  />;
};