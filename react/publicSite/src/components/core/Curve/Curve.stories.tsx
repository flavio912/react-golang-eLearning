import * as React from 'react';
import Curve, { Link } from './Curve';
import { withKnobs, text, boolean, object } from '@storybook/addon-knobs';
// Modules for mocking router
//@ts-ignore
import { Resolver } from 'found-relay';
//@ts-ignore
import { MemoryProtocol } from 'farce';
import { createFarceRouter, createRender } from 'found';
import environment from '../../../api/environment';
import TrustedCard from '../Cards/TrustedCard';
import VideoPlayer from '../VideoPlayer';

export default {
  title: 'core/Curve',
  decorators: [withKnobs]
};

// Thumbnail styles
const root = {
  alignSelf: 'flex-end',
  display: 'flex',
  flexDirection: 'column',
  justifyContent: 'flex-end',
  alignItems: 'center',
  width: '100%'
};
const playCircle = {
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  width: '76px',
  height: '76px',
  borderRadius: '76px',
  backgroundColor: 'white',
  opacity: 0.8,
  marginBottom: '40px'
};
const playTriangle = {
  width: 0,
  height: 0,
  marginLeft: 6,
  borderTop: '13.5px solid transparent',
  borderBottom: '13.5px solid transparent',
  borderLeft: '27px solid #0E66E0'
};

const normal = () => {
  const blue: boolean = boolean('Blue', true);
  const description: string = text(
    'Description',
    '“ TTC Lorem ipsum dolor sit amet, consectetur adipiscing elit, se oddo eiusmod tempor incididunt ut labore Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusm.”'
  );
  const link: Link = object('Link', {
    title: 'Read & watch their story',
    link: '/'
  });
  return <Curve logo="Dhl" description={description} link={link} blue={blue} />;
};

const video = () => {
  const blue: boolean = boolean('Blue', true);
  const description: string = text(
    'Description',
    '“ TTC Lorem ipsum dolor sit amet, consectetur adipiscing elit, se oddo eiusmod tempor incididunt ut labore Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusm.”'
  );
  const link: Link = object('Link', {
    title: 'Read & watch their story',
    link: '/'
  });
  return (
    <VideoPlayer
      source={require('assets/Stock_Video.mp4')}
      thumbnail={
        // @ts-ignore
        <div style={root}>
          <div style={playCircle}>
            <div style={playTriangle} />
          </div>
          <Curve
            width={853}
            logo="Dhl"
            description={description}
            link={link}
            blue={blue}
          />
        </div>
      }
      width={853}
    />
  );
};

const home = () => {
  return (
    <div style={{ backgroundColor: 'black' }}>
      <Curve
        height={140}
        children={
          <TrustedCard
            text="Trusted by more than 1,000 businesses in 120 countries."
            noShadow
          />
        }
      />
    </div>
  );
};

// Mock Router
const StoryRouter = createFarceRouter({
  historyProtocol: new MemoryProtocol('/'),
  routeConfig: [{ path: '/', Component: normal }],
  render: createRender({})
});

export const standard = () => (
  <StoryRouter resolver={new Resolver(environment)} />
);

// Mock Router
const Router = createFarceRouter({
  historyProtocol: new MemoryProtocol('/'),
  routeConfig: [{ path: '/', Component: home }],
  render: createRender({})
});

export const homepage = () => <Router resolver={new Resolver(environment)} />;

// Mock Router
const VideoRouter = createFarceRouter({
  historyProtocol: new MemoryProtocol('/'),
  routeConfig: [{ path: '/', Component: video }],
  render: createRender({})
});

export const thumbnail = () => (
  <VideoRouter resolver={new Resolver(environment)} />
);
