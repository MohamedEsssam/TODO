import React, { useEffect, useState } from "react";
import { getTodos, remove, update } from "../../services/todoServices";
import { currentUser } from "../../services/userServices";
import List from "./list";
import "./style.css";
import NewTodo from "../forms/customInput";
import DoughnutComponent from "../doughnut/doughnut";

const TodoList = React.memo(() => {
  const user = currentUser();
  const [todos, setTodos] = useState([]);
  const [inProgressTodos, setInProgressTodos] = useState([]);
  const [completedTodos, setCompletedTodos] = useState([]);
  const [pausedTodos, setPausedTodos] = useState([]);
  const [fetched, setFetched] = useState(false);

  useEffect(() => {
    if (!fetched) loadTodos();
  }, []);

  const loadTodos = async () => {
    const todos = await getTodos(user.userId);
    setTodos(todos.data.data);

    filterData(todos.data.data);
    setFetched(true);
  };

  const filterData = (todos) => {
    let inProgressTodos = todos.filter((todo) => {
      return todo.status === "inProgress";
    });
    setInProgressTodos(() => [...[], ...inProgressTodos]);

    let completedTodos = todos.filter((todo) => {
      return todo.status === "completed";
    });
    setCompletedTodos(() => [...[], ...completedTodos]);

    let pausedTodos = todos.filter((todo) => {
      return todo.status === "paused";
    });
    setPausedTodos(() => [...[], ...pausedTodos]);
  };

  const deleteTodo = async (todoId) => {
    let data = {};
    data.todoId = todoId;
    const todo = await remove(user.userId, data);

    if (todo) loadTodos();
  };

  const updateTodoStatus = async (todoId, status) => {
    let data = {};
    data.todoId = todoId;
    data.status = status;
    const todo = await update(user.userId, data);

    if (todo) loadTodos();
  };

  return (
    <>
      <DoughnutComponent
        dataSet={{
          data: [
            inProgressTodos.length,
            completedTodos.length,
            pausedTodos.length,
          ],
          backgroundColor: ["#FFCE56", "#36A2EB", "#FF6384"],
          hoverBackgroundColor: ["#FFCE56", "#36A2EB", "#FF6384"],
        }}
      />
      <img
        src={require("../../asset/animation.gif")}
        alt="loading..."
        style={{ width: "200px", position: "relative", left: "450px" }}
      />
      <div class="TodoContainer">
        <NewTodo initialValues={{ text: "" }} loadTodos={loadTodos} />
        <List
          label="In Progress Tasks"
          todos={inProgressTodos}
          type="inProgress"
          deleteTodo={deleteTodo}
          updateTodoStatus={updateTodoStatus}
        />
        <List
          label="Completed Tasks"
          todos={completedTodos}
          type="completed"
          deleteTodo={deleteTodo}
          updateTodoStatus={updateTodoStatus}
        />
        <List
          label="Paused Tasks"
          todos={pausedTodos}
          type="paused"
          deleteTodo={deleteTodo}
          updateTodoStatus={updateTodoStatus}
        />
      </div>
    </>
  );
});

export default TodoList;
