import React from 'react';
import {graphql, createFragmentContainer} from 'react-relay'

type Props = {
  propName: {
    complete: boolean;
    text: string;
  };
};

const ExampleComponent = ({ propName: { complete, text } }: Props) => (
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
    propName: graphql`
      # As a convention, we name the fragment as '<ComponentFileName>_<propName>'
      fragment ExampleComponent_propName on Todo {
        complete
        text
      }
    `
  },
)