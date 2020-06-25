import * as React from "react";
import Homepage from "./Homepage";
import { withKnobs } from "@storybook/addon-knobs";
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
  title: "Overview/Homepage",
  decorators: [withKnobs],
};

const normal = () => {
  return <Homepage
        imageURL={require('assets/Homepage-stock-photo.png')}
        onView={() => console.log("View")}
        onDemo={() => console.log("Demo")}
  />;
};

// Mock Router
const StoryRouter = createFarceRouter({
    historyProtocol: new MemoryProtocol('/'),
    routeConfig: [{ path: '/', Component: normal }],
    render: createRender({})
  })
  
export const standard = () => (<StoryRouter resolver={new Resolver(environment)} />);