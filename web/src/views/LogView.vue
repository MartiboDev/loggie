<template>
	<main class="log-view">
		<LogTable :logs="logs" />
	</main>
</template>

<script setup lang="ts">
import setupWebsocket from "@/util/setupWebsocket"
import { ref, onMounted, type Ref } from "vue"
import type { Log } from "@/models/Log"

const logs: Ref<Log[]> = ref([])

onMounted(() => {
	fetch("http://localhost:8080/")
		.then((response) => response.json())
		.then((data) => {
			logs.value = data ?? []
		})
		.catch((error) => {
			console.error("Error fetching logs:", error)
		})
})

const ws = setupWebsocket()

ws.onmessage = (event) => {
	console.log("Received message from server:", event.data)

	const log: Log = JSON.parse(event.data)

	logs.value.push(log)
}
</script>

<style scoped lang="scss">
.log-view {
	padding: 1.5rem;
}
</style>
