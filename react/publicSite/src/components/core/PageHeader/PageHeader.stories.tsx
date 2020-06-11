import * as React from "react";
import PageHeader, { Archetypes, ButtonLink } from "./PageHeader";
import { withKnobs, text, array, object, select } from "@storybook/addon-knobs";
import CoursePreview from "components/Misc/CoursePreview";

// Modules for mocking router
//@ts-ignore
import { Resolver } from 'found-relay';
//@ts-ignore
import { MemoryProtocol } from 'farce';
import {
  createFarceRouter,
  createRender
} from 'found';
import environment from '../../../api/environment';

export default {
  title: "core/PageHeader",
  decorators: [withKnobs],
};

const archetypes: Archetypes[] = ["default", "buttons", "course"];

const defaultButtons: ButtonLink[] = [
  { title: "Regulated Agents", link: "/"},
  { title: "Known Consignor", link: "/"},
  { title: "GSAT", link: "/"},
]

// Course Preview data
const defaultDetails: string[] = [
  " 6 hours of on-demand video",
  "15 modules",
  "104 lessons",
  "4 examinations",
  "Full lifetime access",
  "Access on mobile/tablet/computer",
  "Industry-approved certificate"
]

const normal = () => {
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
            sideComponent={<CoursePreview price="£310.00" details={defaultDetails} video={require("assets/Stock_Video.mp4")}/>}
          />;
};

// Mock Router
const StoryRouter = createFarceRouter({
  historyProtocol: new MemoryProtocol('/'),
  routeConfig: [{ path: '/', Component: normal }],
  render: createRender({})
})

export const withRouter = () => (<StoryRouter resolver={new Resolver(environment)} />);