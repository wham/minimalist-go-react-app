import React from "react";
import { Todo } from "./todo";

interface ListProps {
  todos: Todo[];
}

export function List({ todos }: ListProps) {
  return (
    <ul>
      {todos.map((todo) => {
        return <li>{todo.text}</li>;
      })}
    </ul>
  );
}
