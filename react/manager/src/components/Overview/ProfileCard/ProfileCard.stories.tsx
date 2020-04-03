import * as React from 'react';
import ProfileCard from './ProfileCard';
import { PaddingOptions } from '../../core/Card';
import { Field } from '../../core/InfoField';
import { withKnobs, select, text, object } from '@storybook/addon-knobs';

export default {
  title: 'Overview/ProfileCard',
  decorators: [withKnobs]
}

const paddingOptions: PaddingOptions[] = ["none", "small", "medium", "large"];

const defaultInfo = [{ fieldName: 'Name', value: 'Fred Eccleston' }, { fieldName: 'Role', value: 'Group Leader' }];

export const plain = () => {
  const padding: PaddingOptions = select("Padding", paddingOptions, "medium");
  const headingText: string = text("Heading", "");
  const fieldsData: Field[] = object("Data", [ ...defaultInfo ]);
  return <ProfileCard heading={headingText || "Profile"} fields={fieldsData} padding={padding} />
}