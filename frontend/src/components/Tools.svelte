<script lang="ts">
  import gamestate from "../store/gamestate";
  import { send } from "../websocket";
  import { type Tool } from "../types";

  const handleTool = (tool: Tool) => {
    send({ type: tool.id, payload: tool.id });
  };
</script>

<div class="tools">
  {#each $gamestate.tools as tool}
    <button
      class="tool"
      disabled={!tool.is_enabled}
      on:click={() => handleTool(tool)}>{tool.name}</button
    >
  {/each}
</div>

<style>
  div {
    grid-area: tools;
    text-align: left;
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }
</style>
