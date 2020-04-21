import React from "react";
import { graphql, createFragmentContainer } from "react-relay";

type Props = {
  info: {
    info: string;
  };
};

const ExampleComponent = ({ info }: Props) => {
  console.log(info && info);
  return <div>{info && info.info}</div>;
};

export default createFragmentContainer(
  ExampleComponent,
  // Each key specified in this object will correspond to a prop available to the component
  {
    info: graphql`
      # As a convention, we name the fragment as '<ComponentFileName>_<propName>'
      fragment ExampleComponent_info on Query {
        info
      }
    `,
  }
);
