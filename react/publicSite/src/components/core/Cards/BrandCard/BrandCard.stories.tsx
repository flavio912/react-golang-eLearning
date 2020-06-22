import * as React from "react";
import BrandCard, { Author } from "./BrandCard";
import { withKnobs, text, boolean, object } from "@storybook/addon-knobs";

// Modules for mocking router
//@ts-ignore
import { Resolver } from 'found-relay';
//@ts-ignore
import { MemoryProtocol } from 'farce';
import {
  createFarceRouter,
  createRender
} from 'found';
import environment from '../../../../api/environment';

export default {
  title: "core/Cards/BrandCard",
  decorators: [withKnobs],
};

const defaultAuthor: Author = {
    name: "James Mchale",
    title: "Head of Compliance",
    quote: "What TTC deliver is world class training at such a pace this industry have never seen before. Highly recommended!"
}

const normal = () => {
    const showQuote: boolean = boolean("Show Quote", false);
    const textValue: string = text("Text", "TTC Lorem ipsum dolor sit amet, consectetur adipiscing elitse oddo eiusmod tempor incididunt ut labore Lorem ipsum dolor sit amet, consecte");
    const author: Author = object("Author", defaultAuthor);

  return <BrandCard
        logoURL={require("assets/logo/ttc-logo.svg")}
        link="/"
        text={textValue}
        author={showQuote ? author : undefined}
    />;
};

// Mock Router
const StoryRouter = createFarceRouter({
    historyProtocol: new MemoryProtocol('/'),
    routeConfig: [{ path: '/', Component: normal }],
    render: createRender({})
  })
  
  export const withRouter = () => (<StoryRouter resolver={new Resolver(environment)} />);