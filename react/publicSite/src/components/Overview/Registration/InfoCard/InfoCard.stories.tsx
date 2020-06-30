import * as React from "react";
import InfoCard from "./InfoCard";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "Overview/Registration/InfoCard",
  decorators: [withKnobs],
};

export const normal = () => {
    const title: string = text('Title', 'New to TTC Hub?');
    const subtitle: string = text('Subtitle', 'Book a 30 minute demo and get your questions answered with one of our customer champions');
    const imageTitle: string = text('Image Title', 'How does it work?');
    const imageSubtitle: string = text('Image Subtitle', 'Simply select a date and time thats suitable for you and we’ll send you through the video call details to confirm.');
  return <InfoCard
    title={title}
    subtitle={subtitle}
    imageURL={require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')}
    imageTitle={imageTitle}
    imageSubtitle={imageSubtitle}
  />;
};