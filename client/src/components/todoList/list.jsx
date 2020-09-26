import React, { useState } from "react";
import { Button, Collapse, Card } from "react-bootstrap";

const List = React.memo(
  ({ todos, label, type, deleteTodo, updateTodoStatus }) => {
    const [open, setOpen] = useState(false);
    const status = new Map();
    status.set("inProgress", "warning");
    status.set("completed", "primary");
    status.set("paused", "danger");
    return (
      <>
        <Button
          onClick={() => setOpen(!open)}
          id="onProgress"
          variant={status.get(type)}
        >
          {label}
        </Button>
        {todos.map((todo) => (
          <Collapse in={open} key={todo.todoId}>
            <Card
              id="onProgressList"
              key={todo.todoId}
              bg={status.get(type)}
              text="white"
              className="mb-2"
            >
              <Card.Body>
                <Card.Text>{todo.text}</Card.Text>
                <Button
                  variant="danger"
                  onClick={() => deleteTodo(todo.todoId)}
                >
                  Delete
                </Button>
                {type === "inProgress" ? (
                  <Button
                    variant="primary"
                    onClick={() => updateTodoStatus(todo.todoId, "paused")}
                  >
                    Paused
                  </Button>
                ) : (
                  <Button
                    variant="primary"
                    onClick={() => updateTodoStatus(todo.todoId, "inProgress")}
                  >
                    Redo
                  </Button>
                )}
                {type !== "completed" ? (
                  <Button
                    variant="primary"
                    onClick={() => updateTodoStatus(todo.todoId, "completed")}
                  >
                    Finished
                  </Button>
                ) : (
                  <Button
                    variant="primary"
                    onClick={() => updateTodoStatus(todo.todoId, "inProgress")}
                  >
                    Redo
                  </Button>
                )}
                <Button variant="dark">Edit</Button>
              </Card.Body>
            </Card>
          </Collapse>
        ))}
      </>
    );
  }
);

export default List;
