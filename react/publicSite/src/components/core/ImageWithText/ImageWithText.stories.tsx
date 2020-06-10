import * as React from "react";
import  ImageWithText, { Link, Row } from "./ImageWithText";
import { withKnobs, text, object, boolean } from "@storybook/addon-knobs";

export default {
  title: "core/ImageWithText",
  decorators: [withKnobs],
};

const defaultStack: Row[] = [
    {
        iconName: "CourseCertificates", text: "All of our friendly and knowledgable team are available via email and live chat.",
        link: { title: "World Class 24x7 Support", link: "TBD"}
    },
    {
        iconName: "CourseCertificates", text: "Stay tuned for regular webinars and live QA sessions with the TTC team.",
        link: { title: "Webinars &amp; Live Sessions", link: "TBD"}
    },
    {
        iconName: "CourseCertificates", text: "Got a question that needs an immediate answer? Try our knowledge base.",
        link: { title: "Knowledge Base", link: "TBD"}
    },
]

export const normal = () => {
    const showStack: boolean = boolean("Show Stack", false);
    const textRight: boolean = boolean("Show Text On Right", false);
    const stack: Row[] = object("Stack", defaultStack)
    const title: string = text("Title", "Online Courses");
    const subtitle: string = text("Subtitle", "Get certified online");
    const description: string = text("Description", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqu Lorem ipsum dolor sit amet, consectetur adipiscing.");
    const link: Link = object("Link", { title: "See all Online Courses", link: "TBD" });
    const image: string = text("Image", require("assets/SampleImage_ClassroomCoursesDetail_Feat.png"));
  return <ImageWithText 
    title={title}
    subtitle={subtitle}
    description={description}
    link={link}
    image={image}
    stack={showStack ? stack : undefined}
    textRight={textRight}
  />;
};