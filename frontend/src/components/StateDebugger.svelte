<script lang="ts">
  import gamestate from "../store/gamestate";
</script>

<details>
  <summary>Gamestate</summary>
  <pre>{JSON.stringify($gamestate, null, 2)}</pre>
</details>

<style>
  details {
    grid-area: debugger;
    text-align: left;
    background-color: #333;
    color: white;
    padding: 1rem;
    border-radius: 10px;
  }
</style>
