<template>
	<main class="log-view">
		<DataTable :value="logs">
			<Column field="level" header="Level" style="width: 5%">
				<template #body="{ data }">
					<div class="cell-center">
						<i
							class="pi"
							:class="logLevels[data.level].icon"
							:style="{ color: `var(--${logLevels[data.level].color})` }"
						></i>
						<span>{{
							data.level.charAt(0).toUpperCase() + data.level.slice(1)
						}}</span>
					</div>
				</template>
			</Column>
			<Column field="category" header="Category" style="width: 10%"></Column>
			<Column field="timestamp" header="Timestamp" style="width: 15%"></Column>
			<Column field="message" header="Message" style="width: 70%"></Column>
		</DataTable>
	</main>
</template>

<script setup>
import { ref, onMounted } from "vue"

onMounted(() => {
	logs.value = [
		{
			level: "debug",
			category: "User",
			timestamp: new Date().toISOString(),
			message: "{ 'user': 'John', 'action': 'login' }",
		},
		{
			level: "info",
			category: "User",
			timestamp: new Date().toISOString(),
			message: "User logged in",
		},
		{
			level: "warning",
			category: "User",
			timestamp: new Date().toISOString(),
			message: "User tried deleting profile",
		},
		{
			level: "error",
			category: "System",
			timestamp: new Date().toISOString(),
			message: "Error occurred in system",
		},
	]
})

const logs = ref([])

const logLevels = {
	debug: { icon: "pi-question-circle", color: "gray-500" },
	info: { icon: "pi-info-circle", color: "cyan-500" },
	warning: { icon: "pi-exclamation-triangle", color: "orange-500" },
	error: { icon: "pi-exclamation-circle", color: "red-500" },
}
</script>

<style scoped lang="scss">
.log-view {
	padding: 1.5rem;
}

.pi {
	font-size: 1.1rem;
}

.cell-center {
	display: flex;
	align-items: center;
	gap: 0.5rem;
}
</style>
