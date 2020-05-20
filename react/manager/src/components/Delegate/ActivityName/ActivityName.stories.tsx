import * as React from "react";
import ActivityName from "./ActivityName";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "Delegate/ActivityName",
  decorators: [withKnobs],
};

export const normal = () => {
  return (
    <ActivityName
      userName={text("Text", "Bruce Willis")}
      title={text(
        "Text",
        "Bruce failed the Dangerous Goods by Road Awareness Course"
      )}
    />
  );
};
