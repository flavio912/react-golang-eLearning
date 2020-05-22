import * as React from "react";
import { withKnobs, text } from "@storybook/addon-knobs";
import FlatCard from "./FlatCard";

export default {
  title: "FlatCard",
  decorators: [withKnobs],
};

export const plain = () => {
  return (
    <FlatCard
      className={text("className", "flat-card")}
      children={<div>This is a test</div>}
      backgroundColor={"#fff"}
      style={{
        boxShadow: `2px 2px 10px 0 rgba(0,0,0,0.07)`,
      }}
    />
  );
};
