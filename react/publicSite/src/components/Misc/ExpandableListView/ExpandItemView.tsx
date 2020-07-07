import React, { ReactElement } from 'react';
import { ExpandItemType } from './types';
import Icon from 'sharedComponents/core/Icon';
type Props = {
  item: ExpandItemType;
  onClickItem: () => void;
};

export const ExpandItemView = ({ item, onClickItem }: Props): ReactElement => (
  <div
    style={{
      padding: 20,
      paddingLeft: 33,
      paddingRight: 33,
      marginBottom: 20,
      backgroundColor: 'white',
      border: '1px solid #E9EBEB',
      boxShadow:
        '0 2px 10px 0 rgba(0,0,0,0.15), 4px 2px 10px -2px rgba(0,0,0,0.06)',
    }}
  >
    <div
      onClick={onClickItem}
      style={{
        flexDirection: 'row',
        display: 'flex',
        alignItems: 'center',
      }}
    >
      <Icon
        name={item.isExpanded ? 'Down_Arrow' : 'Right_Arrow'}
        size={25}
        color="#0E63E8"
      />
      <p
        style={{
          paddingLeft: 33,
          color: '#0C152E',
          fontFamily: 'Muli',
          fontSize: 25,
          fontWeight: 800,
        }}
      >
        {item.title}
      </p>
    </div>

    {item.isExpanded && (
      <p
        style={{
          paddingTop: 35,
          color: '#737988',
          fontFamily: 'Muli',
          fontSize: 25,
        }}
      >
        {item.description}
      </p>
    )}
  </div>
);
