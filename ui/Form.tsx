import React, { useState } from "react";
import { Todo } from "./todo";

interface FormProps {
  onAdd: (todo: Todo) => void;
}

export function Form({ onAdd }: FormProps) {
  const [text, setText] = useState("");

  const onClick = () => {
    onAdd({ text });
    setText("");
  };

  return (
    <div className="form">
      <input
        type="text"
        placeholder="TODO"
        value={text}
        onInput={(e) => setText((e.target as HTMLInputElement).value)}
      />
      <button onClick={onClick}>Add</button>
    </div>
  );
}
