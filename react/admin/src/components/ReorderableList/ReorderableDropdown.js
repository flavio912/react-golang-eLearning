import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import {
  Typography,
  Card,
  Collapse,
  ListItem,
  ListItemText,
  ListItemIcon,
  IconButton
} from '@material-ui/core';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import ExpandLessIcon from '@material-ui/icons/ExpandLess';
import DragIndicatorIcon from '@material-ui/icons/DragIndicator';
import DeleteIcon from '@material-ui/icons/Delete';
import ReoderableList from './ReorderableList';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  bold: {
    fontWeight: 'bold'
  },
  margin: {
    marginRight: theme.spacing(1)
  }
}));

export default function ReoderableDropdown({ uuid, title, items, setItems, onDelete }) {
  const classes = useStyles();
  const [ expanded, setExpanded ] = React.useState(false);

  return (
    <Card>
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
              variant="subtitle1"
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
            className={classes.margin}
          >
            {expanded
              ? <ExpandLessIcon />
              : <ExpandMoreIcon />}
          </IconButton>
        </ListItemIcon>
        <ListItemIcon>
          <IconButton
          onClick={() => onDelete(uuid)}
          edge="end"
          >
            <DeleteIcon color="disabled" />
          </IconButton>
        </ListItemIcon>
      </ListItem>
      <Collapse in={expanded} timeout="auto" unmountOnExit>
        <ReoderableList uuid={uuid} items={items} setItems={setItems} />
      </Collapse>
    </Card>
  );
}