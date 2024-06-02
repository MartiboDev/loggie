<template>
	<main class="log-view">
		<Button @click="sendWS()">Send</Button>
		<LogTable :logs="logs" />
	</main>
</template>

<script setup>
import { ref, onMounted } from "vue"

const logs = ref([])

onMounted(() => {
	logs.value = [
		{
			severity: "debug",
			category: "Client",
			resource: "main_race_2",
			timestamp: new Date(),
			message: "{ 'user': 'John', 'action': 'login' }",
		},
		{
			severity: "info",
			category: "Server",
			resource: "core_admin",
			timestamp: new Date(),
			message: "User logged in",
		},
		{
			severity: "warning",
			category: "Server",
			resource: "core_admin",
			timestamp: new Date(),
			message: "User tried deleting profile",
		},
		{
			severity: "error",
			category: "UI",
			resource: "main_race_2",
			timestamp: new Date(),
			message: "Error occurred in system",
		},
	]
})

let socket = new WebSocket("ws://localhost:3000/ws")
let socketOrder = new WebSocket("ws://localhost:3000/orderbookfeed")

socketOrder.onmessage = (event) => {
	console.log("Recieved from the server:", event.data)
}

socket.onmessage = (event) => {
	console.log("Recieved from the server:", event.data)
}

function sendWS() {
	socket.send("Hello from the client!")
}
</script>

<style scoped lang="scss">
.log-view {
	padding: 1.5rem;
}
</style>
