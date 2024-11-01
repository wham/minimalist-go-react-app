import React, { useState, useEffect } from "react";
import { Form } from "./Form";
import { List } from "./List";
import { Todo } from "./todo";

export function App() {
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    fetch("/api/todos").then((response) => {
      response.json().then((todos) => setTodos(todos || []));
    });
  }, []);

  const onAdd = (todo: Todo) => {
    setTodos((todos) => [...todos, todo]);
    fetch("/api/todos", { method: "PUT", body: JSON.stringify(todo) });
  };

  return (
    <>
      <div className="header">
        <img src="/static/logo.svg" />
        <h1>My TODO List</h1>
      </div>
      <Form onAdd={onAdd} />
      <List todos={todos} />
    </>
  );
}
