import React from "react";

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
