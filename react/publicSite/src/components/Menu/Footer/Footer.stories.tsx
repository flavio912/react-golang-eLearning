import * as React from "react";
import Footer, { Column, Link } from "./Footer";
import { withKnobs, object } from "@storybook/addon-knobs";

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
  title: "Menu/Footer",
  decorators: [withKnobs],
};

const defaultLink: Link = {
    id: 0,
    name: "Feature",
    link: ""
}

const defaultNew: Link = {
    id: 1,
    name: "Feature",
    link: "",
    alert: {
        type: "new",
        value: "NEW"
    }
}

const defaultIncrease: Link = {
    id: 2,
    name: "Feature",
    link: "",
    alert: {
        type: "increase",
        value: "99.9%"
    }
}

const defaultColumns: Column[] = [
    {
        id: 0,
        header: "Features", links: [
            defaultLink, defaultLink, defaultLink, defaultLink, defaultNew, defaultLink, defaultLink, defaultLink, defaultLink
        ]
    },
    {
        id: 1,
        header: "Learn More", links: [
            defaultLink, defaultLink, defaultLink, defaultLink, defaultLink, defaultLink
        ]
    },
    {
        id: 2,
        header: "Company", links: [
            defaultLink, defaultLink, defaultLink, defaultLink, defaultLink, defaultLink
        ]
    },
    {
        id: 3,
        header: "Get Help", links: [
            defaultLink, defaultLink, defaultLink, defaultLink, defaultIncrease
        ]
    }
]

const normal = () => {
    const columns: Column[] = object("Columns", defaultColumns);
  return <Footer columns={columns} />;
};

// Mock Router
const StoryRouter = createFarceRouter({
    historyProtocol: new MemoryProtocol('/'),
    routeConfig: [{ path: '/', Component: normal }],
    render: createRender({})
  })
  
  export const withRouter = () => (<StoryRouter resolver={new Resolver(environment)} />);