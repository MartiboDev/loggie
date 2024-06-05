<template>
	<main class="log-view">
		<LogTable :logs="logs" />
	</main>
</template>

<script setup>
import setupWebsocket from "@/util/setupWebsocket"
import { ref, onMounted } from "vue"

const logs = ref([])

onMounted(() => {
	fetch("http://localhost:8080/")
		.then((response) => response.json())
		.then((data) => {
			logs.value = data ?? []
		})
		.catch((error) => {
			console.error("Error fetching logs:", error)
		})
	// logs.value = [
	// 	{
	// 		severity: "debug",
	// 		category: "Client",
	// 		resource: "main_race_2",
	// 		timestamp: new Date(),
	// 		message: "{ 'user': 'John', 'action': 'login' }",
	// 	},
	// 	{
	// 		severity: "info",
	// 		category: "Server",
	// 		resource: "core_admin",
	// 		timestamp: new Date(),
	// 		message: "User logged in",
	// 	},
	// 	{
	// 		severity: "warning",
	// 		category: "Server",
	// 		resource: "core_admin",
	// 		timestamp: new Date(),
	// 		message: "User tried deleting profile",
	// 	},
	// 	{
	// 		severity: "error",
	// 		category: "UI",
	// 		resource: "main_race_2",
	// 		timestamp: new Date(),
	// 		message: "Error occurred in system",
	// 	},
	// ]
})

const ws = setupWebsocket()

ws.onmessage = (event) => {
	console.log("Received message from server:", event.data)

	const log = JSON.parse(event.data)

	logs.value.push(log)
}
</script>

<style scoped lang="scss">
.log-view {
	padding: 1.5rem;
}
</style>
