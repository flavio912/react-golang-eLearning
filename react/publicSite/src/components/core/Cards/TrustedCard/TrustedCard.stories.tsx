import * as React from "react";
import TrustedCard from "./TrustedCard";
import { withKnobs, text, boolean } from "@storybook/addon-knobs";

export default {
  title: "core/Cards/TrustedCard",
  decorators: [withKnobs],
};

export const normal = () => {
    const textValue: string = text("Text", "Trusted by more than 1,000 businesses in 120 countries.");
    const noShadow: boolean = boolean("No Shadow", false);
  return <TrustedCard
        text={textValue}
        noShadow={noShadow}
    />;
};