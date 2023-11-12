import { readable } from 'svelte/store'
import { connect } from '../websocket'
import { gameStateSchema, type GameState } from '../types';


export const initialGameState: GameState = {
  resources: []
}

const gamestate = readable<GameState>(initialGameState, function start(set) {
  const socket = connect();

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
