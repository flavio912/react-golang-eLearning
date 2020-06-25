import * as React from "react";
import CoursePreview from "./CoursePreview";
import { withKnobs, text, array, boolean } from "@storybook/addon-knobs";

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
    const useImage: boolean = boolean("Use Image", false);
    const price: string = text("Price", "£310.00");
    const details: string[] = array("Price", defaultDetails);
  return <CoursePreview
        price={price}
        details={details}
        video={!useImage && require("assets/Stock_Video.mp4")}
        image={useImage && require("assets/Homepage-stock-photo.png")}
    />;
};