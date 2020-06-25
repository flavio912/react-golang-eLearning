import * as React from "react";
import TickBullet from "./TickBullet";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "Misc/TickBullet",
  decorators: [withKnobs],
};

export const normal = () => {
    const title: string = text("Text", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor ")
  return <TickBullet
    text={title}
  />;
};