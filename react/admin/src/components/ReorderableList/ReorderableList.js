import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { DragDropContext, Droppable, Draggable } from "react-beautiful-dnd";

const useStyles = makeStyles(theme => ({
  padding: {
    padding: `${theme.spacing(1)}px 0`
  },
}));

export default function ReoderableList({ items, setItems }) {
  const classes = useStyles();

  const reorder = (list, startIndex, endIndex) => {
    const result = Array.from(list);
    const [removed] = result.splice(startIndex, 1);
    result.splice(endIndex, 0, removed);
  
    return result;
  };
  
  const onDragEnd = (setItems, items, result) => {
    // dropped outside the list
    if (!result.destination) {
      return;
    }
  
    setItems(reorder(
      items,
      result.source.index,
      result.destination.index
    ));
  }

  return (
    <DragDropContext onDragEnd={(result) => onDragEnd(setItems, items, result)}>
      <Droppable droppableId="droppable">
        {(provided) => (
            <div
            {...provided.droppableProps}
            ref={provided.innerRef}
            >
            {items.map((item, index) => (
              <Draggable
                draggableId={`item-${item.id}`}
                index={index}
                key={item.id}
              >
                {(provided) => (
                  <div
                    className={classes.padding}
                    ref={provided.innerRef}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}
                  >
                    {item.component}
                  </div>
                )}
              </Draggable>
            ))}
            {provided.placeholder}
          </div>
        )}
      </Droppable> 
    </DragDropContext>
  );
}