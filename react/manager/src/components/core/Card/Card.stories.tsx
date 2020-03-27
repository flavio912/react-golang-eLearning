import * as React from 'react';
import { storiesOf } from '@storybook/react';
import Card, {PaddingOptions} from './Card';
import { withKnobs, select } from '@storybook/addon-knobs';

export default {
  title: 'Core/Card',
  decorators: [withKnobs]
}

const paddingOptions: PaddingOptions[] = ['none', 'small', 'medium', 'large'];

export const plain = () => {
  const padding: PaddingOptions = select('Padding', paddingOptions, 'none');
  return (<Card padding={padding}>
    A plain card
  </Card>)
}