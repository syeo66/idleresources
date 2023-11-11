import { readable } from 'svelte/store'

const gamestate = readable({}, function start(set) {
  const socket = new WebSocket("ws://localhost:8080/ws");

  socket.addEventListener("open", () => {
    console.log("Opened");
  });

  socket.addEventListener("message", (event) => {
    set(JSON.parse(event.data));
  });

  return () => socket.close()
})

export default gamestate
