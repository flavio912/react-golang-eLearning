import React from "react";
import { graphql, createFragmentContainer } from "react-relay";

import ExampleComponent from "./ExampleComponent/ExampleComponent";

type Props = {
  data: any;
};

const ExamplePage = ({ data }: Props) => {
  return (
    <section>
      <ExampleComponent info={data} />
    </section>
  );
};

export default ExamplePage;
