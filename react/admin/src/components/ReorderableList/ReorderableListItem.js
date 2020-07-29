import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import {
  ListItem,
  ListItemIcon,
  ListItemText,
  IconButton,
  Typography
} from '@material-ui/core';
import DragHandleIcon from '@material-ui/icons/DragHandle';
import DeleteIcon from '@material-ui/icons/Delete';

const useStyles = makeStyles({
  bold: {
    fontWeight: 'bold'
  },
});

export default function ReoderableListItem({ className, uuid, text, onDelete  }) {
  const classes = useStyles();

  return (
    <ListItem button className={className}>
      <ListItemIcon>
        <DragHandleIcon />
      </ListItemIcon>
      <ListItemText
        primary={
          <Typography
            variant="body1"
            color="textPrimary"
            className={classes.bold}
          >
            {text}
          </Typography>
        }
      />
      {onDelete && (
        <ListItemIcon>
          <IconButton
            onClick={() => onDelete(uuid)}
            edge="end"
          >
            <DeleteIcon color="disabled" />
          </IconButton>
        </ListItemIcon>
      )}
    </ListItem>
  );
}