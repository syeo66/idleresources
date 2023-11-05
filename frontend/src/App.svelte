<script lang="ts">
  const socket = new WebSocket("ws://localhost:8080/ws");
  socket.addEventListener("open", () => {
    console.log("Opened");
  });

  socket.addEventListener("message", (event) => {
    console.log("Message from server", event.data);
  });
</script>

<main>
  <h1>Game</h1>
</main>

<style>
</style>
