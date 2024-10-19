import React, { useState } from "react";
import { Todo } from "./todo";

interface FormProps {
  onAdd: (todo: Todo) => void;
}

export function Form({ onAdd }: FormProps) {
  const [text, setText] = useState("");

  const onClick = () => {
    onAdd({ text });
  };

  return (
    <div>
      <input
        type="text"
        value={text}
        onInput={(e) => setText(e.target.value)}
      />
      <button onClick={onClick}>Add</button>
    </div>
  );
}