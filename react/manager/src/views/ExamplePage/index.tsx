import React from 'react';
import {graphql, QueryRenderer} from 'react-relay';

import environment from '../../api/environment';
import { default as _ExamplePage } from './ExamplePage';

type Response = {
    error: any,
    props: any,
  };

const ExamplePage = () => (
    <QueryRenderer
        environment={environment}
        query={graphql`
          query UserQuery {
            viewer {
              id
              ...ExamplePage_props
            }  
          }
        `}
        variables={{}}
        render={({ error, props }: Response) => {
          if (error) {
            return <div>Error!</div>;
          }
          if (!props) {
            return <div>Loading...</div>;
          }
          return <_ExamplePage userID={props.viewer.id} />;
        }}
    />
);

export default ExamplePage;