import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { TextField, Card, CardContent } from '@material-ui/core';
import {TreeView, TreeItem} from '@material-ui/lab';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';
import DragIndicatorIcon from '@material-ui/icons/DragIndicator';
import ReoderableList from './ReorderableList';

const useStyles = makeStyles({
  root: {
    flexGrow: 1,
  },
});

export default function ReoderableDropdown({ value, onChange, placeholder, label, items, setItems }) {
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
            label={
              <TextField
                fullWidth
                label={label}
                onChange={onChange}
                placeholder={placeholder}
                value={value}
              />
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