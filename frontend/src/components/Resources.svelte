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
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    gap: 1rem;
  }

  .resource {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 75px;
    height: 75px;
    border: 1px solid white;
    border-radius: 10px;
  }

  .resource__name {
    font-size: 1rem;
  }

  .resource__amount {
    font-size: 1.3rem;
  }
</style>
