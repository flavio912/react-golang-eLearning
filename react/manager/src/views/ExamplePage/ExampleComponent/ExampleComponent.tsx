import React from 'react';
import {graphql, createFragmentContainer} from 'react-relay'

type Props = {
    queryProp: {
        complete: boolean;
        text: string;
    };
};

const ExampleComponent = ({queryProp: { complete, text }}: Props) => (
    <li>
        <div>
          <input
            checked={complete}
            type="checkbox"
          />
          <label>
            {text}
          </label>
        </div>
    </li>
)

export default createFragmentContainer(
  ExampleComponent,
  // Each key specified in this object will correspond to a prop available to the component
  {
    queryProp: graphql`
      # As a convention, we name the fragment as '<ComponentFileName>_<propName>'
      fragment ExampleComponent_props on Todo {
        complete
        text
      }
    `
  },
)