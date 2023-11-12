<script lang="ts">
  import gamestate from "../store/gamestate";
</script>

<div class="resources">
  {#each $gamestate.resources as resource}
    <div class="resource">
      <div class="resource__name">{resource.name}</div>
      <div class="resource__amount">{resource.amount}</div>
    </div>
  {/each}
</div>

<style>
  .resources {
    grid-area: resources;
    display: flex;
    flex-direction: column;
  }

  .resource {
    display: flex;
    gap: 1rem;
    flex-direction: row;
  }
</style>
