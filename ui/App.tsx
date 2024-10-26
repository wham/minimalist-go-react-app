import React, { useEffect, useState } from "react";

export function App() {
  const [message, setMessage] = useState("");

  useEffect(() => {
    callServer();
  }, []);

  const callServer = () => {
    fetch("/api").then((response) => {
      response.json().then((response) => setMessage(response.message));
    });
  };

  return (
    <>
      <img src="/static/logo.svg" />
      <button onClick={callServer}>Call server</button>
      <p>{message}</p>
    </>
  );
}
