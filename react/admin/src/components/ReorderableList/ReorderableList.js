import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { DragDropContext, Droppable, Draggable } from "react-beautiful-dnd";

const useStyles = makeStyles(theme => ({
  padding: {
    padding: `${theme.spacing(2)}px 0`
  },
}));

export default function ReoderableList({ multiple, uuid, items, setItems, newItem }) {
  const classes = useStyles();

  const getIndices = (list, location, id) => {
    let topIndex, itemIndex, hasItems;
    // if outermost
    if (location.droppableId === 'droppable') {
      topIndex = location.index;
      hasItems = list.find(x => x.uuid === id)?.items?.length > 0;
      // if top level item
    } else if (list.find(x => x.uuid === location.droppableId)) {
      itemIndex = list.findIndex(x => x.uuid === location.droppableId);
      hasItems = list[itemIndex]?.items.findIndex(x => x.uuid === id) < 0;
    }
    return {topIndex, itemIndex, hasItems};
  }

  const swap = (list, result) => {
    const source = getIndices(list, result.source, result.draggableId);
    let dest = getIndices(list, result.destination, result.draggableId);
    let removed;

    // Remove from top level
    if (source.topIndex !== undefined) {
      removed = list.splice(source.topIndex, 1);
      // Remove from item's items
    } else if (source.itemIndex !== undefined) {
      removed = list[source.itemIndex].items.splice(result.source.index, 1)
    }

    // Update indices
    dest = getIndices(list, result.destination, result.draggableId);
    removed = removed && (removed[0]?.item ? removed[0].item : removed[0]);
  
    // Add to top level
    if (removed && dest.topIndex !== undefined) {
      list.splice(dest.topIndex, 0, newItem(removed));
      // Add to item's items
    } else if (removed && dest.itemIndex !== undefined && !source.hasItems && dest.hasItems) {
      list[dest.itemIndex].items.splice(result.destination.index, 0, removed)
    }
    return list;
  }
  
  const onDragEnd = (setItems, items, result) => {
    // dropped outside the list
    if (!result.destination) {
      return;
    }
    return setItems(swap(Array.from(items), result));
  }

  const Context = ({children}) => (
    uuid ? (children) :
    <DragDropContext onDragEnd={(result) => onDragEnd(setItems, items, result)}>
      {children}
    </DragDropContext>
  );

  return (
    <Context>
      <Droppable droppableId={uuid ? uuid : 'droppable'}>
        {(provided) => (
            <div
            {...provided.droppableProps}
            ref={provided.innerRef}
            >
            {items && items.map((item, index) => (
              <Draggable
                draggableId={item.item ? item.item.uuid : item.uuid}
                index={index}
                key={item.item ? item.item.uuid : item.uuid}
              >
                {(provided) => (
                  <div
                    className={classes.padding}
                    ref={provided.innerRef}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}
                  >
                    {multiple ? (
                      <Droppable droppableId={item.item ? item.item.uuid : item.uuid}>
                        {(provided) => (
                          <div
                          {...provided.droppableProps}
                          ref={provided.innerRef}
                          >
                            {item.component}
                            {provided.placeholder}
                          </div>
                        )}
                      </Droppable>
                    ) : (
                      item.component
                    )}
                  </div>
                )}
              </Draggable>
            ))}
            {provided.placeholder}
          </div>
        )}
      </Droppable>
    </Context>
  );
}