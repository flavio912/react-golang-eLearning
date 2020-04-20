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
          query ExamplePage_Query {
            user {
              ...ExamplePage_propName
            }
          }
        `}
        variables={{}}
        render={({ error, props }: Response) => {
          // Remove when connected to api
          /*if (error) {
            console.log("error" + error)
            return <div>Error!</div>;
          }
          if (!props) {
            return <div>Loading...</div>;
          }*/
          const temp = {
            totalCount: 1,
            completedCount: 1,
            todos: {
              edges: [{ complete: true, text: 'Todo' }]
            },
          };
          return <_ExamplePage propName={temp} /* propName={props} */ />;
        }}
    />
);

export default ExamplePage;