import * as React from "react";
import CoursePreview from "./CoursePreview";
import { withKnobs, text, array } from "@storybook/addon-knobs";

export default {
  title: "Misc/CoursePreview",
  decorators: [withKnobs],
};

const defaultDetails: string[] = [
    " 6 hours of on-demand video",
    "15 modules",
    "104 lessons",
    "4 examinations",
    "Full lifetime access",
    "Access on mobile/tablet/computer",
    "Industry-approved certificate"
]

export const normal = () => {
    const price: string = text("Price", "£310.00");
    const details: string[] = array("Price", defaultDetails);
  return <CoursePreview
        price={price}
        details={details}
    />;
};