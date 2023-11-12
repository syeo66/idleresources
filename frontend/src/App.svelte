<script lang="ts">
  import Sources from "./components/Sources.svelte";
  import Resources from "./components/Resources.svelte";
  import StateDebugger from "./components/StateDebugger.svelte";
  import Tools from "./components/Tools.svelte";
</script>

<main>
  <h1>Game</h1>
  <Resources />
  <Sources />
  <Tools />
  <StateDebugger />
</main>

<style>
  main {
    display: grid;
    width: 100%;
    grid-template-columns: 3fr 1fr 1fr;
    grid-gap: 1rem;
    grid-template-areas: "header header header" "sources resources tools" "debugger resources tools";
  }

  h1 {
    grid-area: header;
  }
</style>
