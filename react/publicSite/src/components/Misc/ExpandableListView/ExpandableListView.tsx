import React, { useState, ReactElement } from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import theme, { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon';

import 'react-multi-carousel/lib/styles.css';

import { ExpandItemView } from './ExpandItemView';
import { ExpandItemType } from './types';

const useStyles = createUseStyles((theme: Theme) => ({
  carousel: { backgroundColor: theme.colors.primaryWhite },
}));

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
    <div className={classes.carousel}>
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
