import React from "react";
import { graphql, createFragmentContainer } from "react-relay";

import ExampleComponent from "./ExampleComponent/ExampleComponent";

type Props = {
  propName: {
    totalCount: number;
    completedCount: number;
    todos: any;
  };
};

const ExamplePage = ({ propName: {totalCount, completedCount, todos} }: Props) => (
  <section>
    <input checked={totalCount === completedCount} type="checkbox" />
    <ul>
      {todos.edges.map((edge: any) => (
        <ExampleComponent propName={edge} />
      ))}
    </ul>
  </section>
);

export default createFragmentContainer(
  ExamplePage,
  {
    propName: graphql`
      # As a convention, we name the fragment as '<ComponentFileName>_<PropName>'
      fragment ExamplePage_propName on User {
        todos(
          first: 2147483647  # max GraphQLInt, to fetch all todos
        ) {
          edges {
            node {
              id,
              # We use the fragment defined by the child Todo component here
              ...ExampleComponent_propName,
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
