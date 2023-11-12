<script lang="ts">
  import gamestate from "../store/gamestate";
</script>

<div class="sources">
  {#each $gamestate.resources as resource}
    <div class="source">
      <div class="source__name">{resource.name}</div>
      <div class="source__delta">{resource.delta}</div>
    </div>
  {/each}
</div>

<style>
  .sources {
    grid-area: sources;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    gap: 1rem;
  }

  .source {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 75px;
    height: 75px;
    border: 1px solid white;
    border-radius: 10px;
  }

  .source__name {
    font-size: 1rem;
  }

  .source__delta {
    font-size: 1.3rem;
  }
</style>
