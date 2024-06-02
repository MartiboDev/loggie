<template>
	<DataTable :value="logs" tableStyle="min-width: 50rem">
		<Column field="severity" header="Severity"></Column>
		<Column field="category" header="Category"></Column>
		<Column field="resource" header="Resource"></Column>
		<Column field="timestamp" header="Timestamp"></Column>
		<Column field="message" header="Message"></Column>
	</DataTable>
</template>

<script setup lang="ts">
import { computed, type PropType } from "vue"

interface Log {
	id: number
	severity: string
	category: string
	resource: string
	timestamp: string
	message: string
}

const props = defineProps({
	logs: {
		type: Object as PropType<Log[]>,
		required: true,
	},
})

const logSeverity = {
	debug: { icon: "pi-question-circle", color: "gray-500" },
	info: { icon: "pi-info-circle", color: "cyan-500" },
	warning: { icon: "pi-exclamation-triangle", color: "orange-500" },
	error: { icon: "pi-exclamation-circle", color: "red-500" },
}
</script>

<style scoped>
table,
th,
td {
	width: 100%;
	border: 1px solid var(--border-color);
	border-collapse: collapse;
}

th {
	text-align: left;
}

th,
td {
	padding: 8px;
	text-align: left;
}
</style>
