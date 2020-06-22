import * as React from "react";
import VideoPlayer from "./VideoPlayer";
import { withKnobs, object, boolean } from "@storybook/addon-knobs";
import FloatingVideo, { Author } from "./FloatingVideo";

export default {
  title: "core/VideoPlayer",
  decorators: [withKnobs],
};

const defaultAuthor: Author = {
    name: "Kristian Durhuus",
    title: "Chief Executive Officer",
    quote: "TTC Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore ."
}

export const normal = () => {
  return <VideoPlayer
        width={832}
        height={459}
        source={require("assets/Stock_Video.mp4")}
    />;
};

export const floating = () => {
    const showQuote = boolean("Show Quote", true);
    const author: Author = object("Quote", defaultAuthor);
  return <FloatingVideo
        width={628}
        height={352}
        source={require("assets/Stock_Video.mp4")}
        author={showQuote ? author : undefined}
    />;
};