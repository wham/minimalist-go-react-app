import React, { useState } from "react";
import { Form } from "./Form";
import { List } from "./List";
import { Todo } from "./todo";

export function App() {
  const [todos, setTodos] = useState<Todo[]>([]);

  const onAdd = (todo: Todo) => {
    setTodos((todos) => [...todos, todo]);
  };

  return (
    <div id="container">
      <Form onAdd={onAdd} />
      <List todos={todos} />
    </div>
  );
}
