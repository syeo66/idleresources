import { readable } from 'svelte/store'
import { z } from 'zod';

const resourceSchema = z.object({
  id: z.string(),
  name: z.string(),
  amount: z.number(),
  delta: z.number(),
  total: z.number(),
  is_automated: z.boolean(),
})

const gameStateSchema = z.object({
  resources: z.array(resourceSchema)
})
type GameState = z.infer<typeof gameStateSchema>

const initialGameState: GameState = {
  resources: []
}

const gamestate = readable<GameState>(initialGameState, function start(set) {
  const socket = new WebSocket("ws://localhost:8080/ws");

  socket.addEventListener("open", () => {
    console.log("Opened");
  });

  socket.addEventListener("message", (event) => {
    try 
  {
      const data = JSON.parse(event.data);
      const gameState = gameStateSchema.parse(data);
      set(gameState);
    } catch (error) {
      console.error(error);
    }
  });

  return () => socket.close()
})

export default gamestate
