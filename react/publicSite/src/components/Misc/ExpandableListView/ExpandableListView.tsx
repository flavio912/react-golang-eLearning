import React, { useState, ReactElement } from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';

import { ExpandItemView } from './ExpandItemView';

const useStyles = createUseStyles((theme: Theme) => ({
  rootExpandableListView: {
    backgroundColor: theme.colors.primaryWhite,
    width: '100%'
  },
}));

export type ExpandItemType = {
  id: number;
  title: string;
  description: string;
  isExpanded: boolean;
};

type Props = {
  data: ExpandItemType[];
};

function ExpandableListView({ data }: Props): ReactElement {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [expandList, setExpandList] = useState(data);

  function onClickItem(item: ExpandItemType): () => void {
    return (): void => {
      setExpandList(
        expandList.map(
          (listItem): ExpandItemType => {
            if (listItem.id == item.id) {
              return { ...listItem, isExpanded: !listItem.isExpanded };
            }
            return listItem;
          },
        ),
      );
    };
  }
  return (
    <div className={classes.rootExpandableListView}>
      {expandList.map((item) => (
        <ExpandItemView
          key={item.id}
          item={item}
          onClickItem={onClickItem(item)}
        ></ExpandItemView>
      ))}
    </div>
  );
}

export default ExpandableListView;
