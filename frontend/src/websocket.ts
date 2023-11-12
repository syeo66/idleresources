import { z } from "zod";

let websocket: WebSocket;

const messageSchema = z.object({
  type: z.enum(["collect"]),
  payload: z.string(), 
});

type Message = z.infer<typeof messageSchema>;

export const connect = () => {
  if (websocket?.readyState !== WebSocket.OPEN) {
    websocket = new WebSocket("ws://localhost:8080/ws");
  }

  return websocket;
}

export const send = (message: Message) => {
  if (websocket?.readyState !== WebSocket.OPEN) {
    return;
  }

  websocket.send(JSON.stringify(message));
}

