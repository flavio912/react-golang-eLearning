import * as React from 'react';
import { withKnobs, text } from '@storybook/addon-knobs';
import Question from './Question';

export default {
  title: 'Misc/Question',
  decorators: [withKnobs]
};
const questionImage = {
  id: 1,
  title: text(
    'Question title',
    'What is the authority with responsibility for the transport of dangerous goods by air in the UK?'
  ),
  options: [
    {
      id: '1',
      image: text(
        'option image',
        'https://media.istockphoto.com/photos/chef-knife-picture-id874095794'
      )
    },
    {
      id: '2',
      image: text(
        'option image',
        'https://media.istockphoto.com/photos/traditional-chefs-knife-isolated-on-a-white-background-picture-id832724072'
      )
    }
  ]
};
const questionText = {
  id: 1,
  title: text(
    'Question title',
    'What is the authority with responsibility for the transport of dangerous goods by air in the UK?'
  ),
  options: [
    {
      id: '1',
      title: text('option title', 'The Department for Transport (DfT)'),
      index: text('option index', 'A')
    },
    {
      id: '2',
      title: text('option title', 'The Ministry of Justice (MoJ)'),
      index: text('option index', 'B')
    }
  ]
};
export const plain = () => {
  return (
    <div>
      <Question question={questionText} type="text" onSelected={() => {}} />
      <Question question={questionImage} type="image" onSelected={() => {}} />
    </div>
  );
};
