import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Typography, Card, CardContent } from '@material-ui/core';
import {TreeView, TreeItem} from '@material-ui/lab';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';
import DragIndicatorIcon from '@material-ui/icons/DragIndicator';
import ReoderableList from './ReorderableList';

const useStyles = makeStyles({
  root: {
    flexGrow: 1
  },
  noSelect: {
    '& .MuiTreeItem-content': {
      padding: '10px 0',
    },
    '& .Mui-selected': {
      // Currently not working, need to change this style:
      // .MuiTreeItem-root.Mui-selected > .MuiTreeItem-content .MuiTreeItem-label
      backgroundColor: 'transparent'
    }
  }
});

export default function ReoderableDropdown({ title, items, setItems }) {
  const classes = useStyles();
  const [ expanded, setExpanded ] = React.useState(false);

  return (
    <Card>
      <CardContent>
        <TreeView
            className={classes.root}
            defaultCollapseIcon={<DragIndicatorIcon />}
            defaultExpandIcon={<DragIndicatorIcon />}
            onNodeToggle={() => setExpanded(!expanded)}
        >
          <TreeItem
            nodeId="1"
            className={classes.noSelect}
            label={
              <Typography
                variant="h5"
                color="textPrimary"
              >
                {title}
              </Typography>
            }
            endIcon={
                expanded
                ? <ExpandMoreIcon />
                : <ChevronRightIcon />
            }
          >
            <ReoderableList items={items} setItems={setItems} />
          </TreeItem>
        </TreeView>
      </CardContent>
    </Card>
  );
}