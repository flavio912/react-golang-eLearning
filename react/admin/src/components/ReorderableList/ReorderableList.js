import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { DragDropContext, Droppable, Draggable } from "react-beautiful-dnd";

const useStyles = makeStyles(theme => ({
  padding: {
    padding: `${theme.spacing(1)}px 0`
  },
}));

export default function ReoderableList({ multiple, uuid, items, setItems }) {
  const classes = useStyles();

  const reorder = (list, uuid, startIndex, endIndex) => {
    const result = Array.from(list);
    if (result.find(x => x.uuid === uuid)) {
      console.log('first layer', list, uuid)
      const [removed] = result.splice(startIndex, 1);
      result.splice(endIndex, 0, removed);
      return result;
    }
    result.map(inner => {
      console.log('inner layer', inner, uuid)
      if (inner.items.find(x => x.uuid === uuid)) {
        const [removed] = inner.items.splice(startIndex, 1);
        inner.items.splice(endIndex, 0, removed);
        return inner.items;
      }
    });
    return [];
  };
  
  const onDragEnd = (setItems, items, result) => {
    console.log('dropped', result, items)
    // dropped outside the list
    if (!result.destination) {
      return;
    }

    const temp = Array.from(items);

    // moving between containers
    if (result.source.droppableId !== result.destination.droppableId) {
      console.log('swapped', result, items)
      const sourceIndex = temp.findIndex(x => x.uuid === result.source.droppableId);
      const destIndex = temp.findIndex(x => x.uuid === result.destination.droppableId);
      console.log('ind', sourceIndex, destIndex)
      const deleteIndex = temp[sourceIndex].items.find(x => x.uuid === result.draggableId);
      const [removed] = temp[sourceIndex].items.splice(deleteIndex, 1)
      temp[destIndex].items.push(removed);
    }
  
    setItems(reorder(
      temp,
      result.draggableId,
      result.source.index,
      result.destination.index
    ));
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
            {items.map((item, index) => (
              <Draggable
                draggableId={item.uuid}
                index={index}
                key={item.uuid}
              >
                {(provided) => (
                  <div
                    className={classes.padding}
                    ref={provided.innerRef}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}
                  >
                    {multiple ? (
                      <Droppable droppableId={item.uuid}>
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