import React from 'react';
import { DragDropContext, Droppable, Draggable } from "react-beautiful-dnd";

/**
 * To allow dragging inside a list, 'newItem' must be provided
 * To allow dragging between lists, a 'uuid' must be provided
 */
export default function ReoderableList({ className, multiple, uuid, items, setItems, newItem }) {

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
    // Get removed
    removed = removed && (removed[0]?.item ? removed[0].item : removed[0]);
    // Add to top level
    if (removed && dest.topIndex !== undefined) {
      list.splice(dest.topIndex, 0, newItem(removed));
      // Add to item's items
    } else if (uuid && removed && dest.itemIndex !== undefined && !source.hasItems && dest.hasItems) {
      list[dest.itemIndex].items.splice(result.destination.index, 0, removed);
      // Otherwise add back to the main list
    } else {
      list.splice(dest.itemIndex, 0, newItem(removed));
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
              className={className}
            >
            {items && items.map((item, index) => (
              <Draggable
                draggableId={item.item ? item.item.uuid : item.uuid}
                index={index}
                key={item.item ? item.item.uuid : item.uuid}
                isDragDisabled={newItem === undefined}
              >
                {(DraggableProvided) => (
                  <div
                    ref={DraggableProvided.innerRef}
                    {...DraggableProvided.draggableProps}
                    {...DraggableProvided.dragHandleProps}
                    style={DraggableProvided.draggableProps.style}
                  >
                    {multiple ? (
                      <Droppable droppableId={item.item ? item.item.uuid : item.uuid}>
                        {(DroppableProvided) => (
                          <div
                          {...DroppableProvided.droppableProps}
                          ref={DroppableProvided.innerRef}
                          >
                            {item.component}
                          </div>
                        )}
                      </Droppable>
                    ) : (
                      item.component
                    )}
                    {provided.placeholder}
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