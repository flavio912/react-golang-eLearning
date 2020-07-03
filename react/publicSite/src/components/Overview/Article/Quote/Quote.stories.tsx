import * as React from "react";
import Quote from "./Quote";
import { withKnobs, text, boolean } from "@storybook/addon-knobs";

export default {
  title: "Overview/Article/Quote",
  decorators: [withKnobs],
};

export const normal = () => {
    const leftSide: boolean = boolean("Left side", false);
    const quote: string = text("Quote", "“When we decided that our mission would be ‘to bring humanity back to air travel,’ we knew we would have to build a great culture internally or great customer service would never take root externally, with the traveling public. We needed to build a culture of respect, trust and communication, a culture where we take care of each other.”");
  return <Quote
        quote={quote}
        borderSide={leftSide ? 'left' : 'top'}
    />;
};