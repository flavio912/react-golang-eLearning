import * as React from "react";
import PageHeader, { Archetypes, ButtonLink } from "./PageHeader";
import { withKnobs, text, array, object, select } from "@storybook/addon-knobs";

export default {
  title: "core/PageHeader",
  decorators: [withKnobs],
};

const archetypes: Archetypes[] = ["default", "buttons", "course"];

const defaultButtons: ButtonLink[] = [
  { title: "Regulated Agents", link: "TBD"},
  { title: "Known Consignor", link: "TBD"},
  { title: "GSAT", link: "TBD"},
]

export const normal = () => {
  const archetype: Archetypes = select("Type", archetypes, "default");
  const title: string = text("Title", "About Us");
  const description: string = text("Description", "Our mission is to create the highest quality safety & compliance training in the world");
  const estimatedTime: string = text("Estimated Time", "6 hours");
  const lastUpdated: string = text("Last Updated", "May 2020");
  const history: string[] = array("History", ["Courses", "Aviation Security"]);
  const buttons: ButtonLink[] = object("Buttons", defaultButtons);
  return <PageHeader
            title={title}
            description={description}
            archetype={archetype}
            history={history}
            buttons={buttons}
            estimatedTime={estimatedTime}
            lastUpdated={lastUpdated}
          />;
};