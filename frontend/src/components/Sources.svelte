<script lang="ts">
  import gamestate from "../store/gamestate";
  import { send } from "../websocket";
  import { type Resource } from "../types";

  const handleCollect = (resource: Resource) => {
    send({ type: "collect", payload: resource.id });
  };
</script>

<div class="sources">
  {#each $gamestate.resources as resource}
    <div class="source">
      <div class="source__name">{resource.name}</div>
      <div class="source__delta">{resource.delta}</div>
      {#if !resource.is_automated}
        <button on:click={() => handleCollect(resource)}>collect</button>
      {/if}
    </div>
  {/each}
</div>

<style>
  .sources {
    grid-area: sources;
    display: flex;
    flex-direction: row;
    gap: 1rem;
  }

  .source {
    font-size: 0.8rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 75px;
    height: 75px;
    border: 1px solid white;
    border-radius: 10px;
  }

  .source button {
    padding: 0.2rem 0.4rem;
  }

  .source__delta {
    font-size: 1rem;
  }
</style>
