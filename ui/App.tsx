import React, { useEffect, useState } from "react";
import { Form } from "./Form";
import { List } from "./List";
import { Todo } from "./todo";

export function App() {
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    fetch("/api/todos").then((response) => {
      response.json().then((todos) => setTodos(todos));
    });
  }, []);

  const onAdd = (todo: Todo) => {
    setTodos((todos) => [...todos, todo]);
    fetch("/api/todos", { method: "PUT", body: JSON.stringify(todo) });
  };

  return (
    <div id="container">
      <h1>TODO</h1>
      <Form onAdd={onAdd} />
      <List todos={todos} />
    </div>
  );
}
