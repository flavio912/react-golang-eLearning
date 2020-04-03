import * as React from "react";
import Button, { Archetypes } from "./Button";
import { withKnobs, select, text, boolean } from "@storybook/addon-knobs";

export default {
  title: "Core/Button",
  decorators: [withKnobs],
};

const archetypes: Archetypes[] = ["default", "grey", "submit"];

export const flexible = () => {
  const archetypeTexts = {
    default: "Cancel",
    grey: "Add Delegate",
    submit: "Continue to Terms",
  };
  const archetype: Archetypes = select("Type", archetypes, "default");
  const btnText: string = text("Text", "");
  const bold: boolean = boolean("Bold", false);
  const small: boolean = boolean("Small", false);
  // icons will be text (but not yet)
  const iconLeft: boolean = boolean("Left Icon", false);
  const iconRight: boolean = boolean("Right Icon", false);

  return (
    <Button
      archetype={archetype}
      iconLeft={iconLeft}
      iconRight={iconRight}
      bold={bold}
      small={small}
    >
      {btnText || archetypeTexts[archetype]}
    </Button>
  );
};
