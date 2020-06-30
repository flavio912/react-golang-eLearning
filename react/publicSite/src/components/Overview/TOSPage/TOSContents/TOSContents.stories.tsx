import * as React from "react";
import TOSContents, { LinkDetails } from "./TOSContents";
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
import environment from '../../../../api/environment';

export default {
  title: "Overview/TOSPage/TOSContents",
  decorators: [withKnobs],
};

const defaultLinks: LinkDetails[] = [
    { title: 'Terms of Service', link: '/' },
    { title: 'Privacy Policy', link: '/' },
    { title: 'Security Policy', link: '/' },
    { title: 'GDPR', link: '/' },
    { title: 'Cookie Policy', link: '/' },
    { title: 'List of Sub-processors', link: '/' }
]


const ContentsExample = () => {
    const [selected, setSelected] = React.useState(defaultLinks[0]);
    return <TOSContents
        links={defaultLinks}
        selected={selected}
        setSelected={setSelected}
    />;
  };
  
const normal = () => {
    return (
      <ContentsExample />
    );
  };
  

// Mock Router
const StoryRouter = createFarceRouter({
    historyProtocol: new MemoryProtocol('/'),
    routeConfig: [{ path: '/', Component: normal }],
    render: createRender({})
  })
  
  export const withRouter = () => (<StoryRouter resolver={new Resolver(environment)} />);