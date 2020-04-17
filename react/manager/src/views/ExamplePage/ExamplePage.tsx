import React from 'react';
import {graphql, createFragmentContainer} from 'react-relay';

import ExampleComponent from './ExampleComponent/ExampleComponent'; 

type Props = {
    queryProp: {
        totalCount: number;
        completedCount: number;
        todos: any;
    };
}

const ExamplePage = ({queryProp: { totalCount, completedCount, todos }}: Props) => (
    <section>
        <input
            checked={totalCount === completedCount}
            type="checkbox"
        />
        <ul>
            {todos.edges.map((edge: any) =>
                <ExampleComponent
                    key={edge.node.id}
                    {/*We pass the data required by the component here*/}
                    todo={edge.node}
                />
            )}
        </ul>
    </section>
);

export default createFragmentContainer(
    ExamplePage,
  {
    queryProp: graphql`
      # As a convention, we name the fragment as '<ComponentFileName>_<PropName>'
      fragment ExamplePage_props on User {
        todos(
          first: 2147483647  # max GraphQLInt, to fetch all todos
        ) {
          edges {
            node {
              id,
              # We use the fragment defined by the child component here
              ...ExampleComponent_props,
            },
          },
        },
        id,
        totalCount,
        completedCount,
      }
    `,
  },
);