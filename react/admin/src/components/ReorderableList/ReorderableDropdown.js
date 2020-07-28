import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import {
  Typography,
  Card,
  CardContent,
  Collapse,
  ListItem,
  ListItemText,
  ListItemIcon,
  IconButton
} from '@material-ui/core';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import ExpandLessIcon from '@material-ui/icons/ExpandLess';
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
  },
  bold: {
    fontWeight: 'bold'
  },
});

export default function ReoderableDropdown({ uuid, title, items, setItems }) {
  const classes = useStyles();
  const [ expanded, setExpanded ] = React.useState(false);

  return (
    <Card>
      <CardContent>
        <ListItem
          className={classes.root}
          button
          onClick={() => setExpanded((previousState) => !previousState)}
        >
          <ListItemIcon>
            <IconButton
              edge="start"
            >
              <DragIndicatorIcon />
            </IconButton>
          </ListItemIcon>
          <ListItemText
            primary={
              <Typography
                variant="body1"
                color="textPrimary"
                className={classes.bold}
              >
                {title}
              </Typography>
            }
          />
          <ListItemIcon>
            <IconButton
              edge="end"
            >
              {expanded
                ? <ExpandLessIcon />
                : <ExpandMoreIcon />}
            </IconButton>
          </ListItemIcon>
        </ListItem>
        <Collapse in={expanded} timeout="auto" unmountOnExit>
          <ReoderableList uuid={uuid} items={items} setItems={setItems} />
        </Collapse>
      </CardContent>
    </Card>
  );
}