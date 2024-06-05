<template>
	<DataTable :value="logs" tableStyle="min-width: 50rem" showGridlines>
		<Column field="severity" header="Severity" style="width: 10%"></Column>
		<Column field="category" header="Category" style="width: 10%"></Column>
		<Column field="resource" header="Resource" style="width: 15%"></Column>
		<Column field="timestamp" header="Timestamp" style="width: 15%"></Column>
		<Column field="message" header="Message" style="width: 50%">
			<template #body="{ data }">
				<template v-if="canBeObject(data.message)">
					<pre>{{ formatMessage(data.message) }}</pre>
				</template>
				<template v-else>{{ data.message }}</template>
			</template>
		</Column>
	</DataTable>
</template>

<script setup lang="ts">
import type { Log } from "@/models/Log"
import { type PropType } from "vue"

defineProps({
	logs: {
		type: Object as PropType<Log[]>,
		required: true,
	},
})

function formatMessage(message: string) {
	try {
		const jsonString = message.replace(/'/g, '"')
		const obj = JSON.parse(jsonString)
		return JSON.stringify(obj, null, 2)
	} catch (e) {
		return message
	}
}

function canBeObject(message: string) {
	try {
		const jsonString = message.replace(/'/g, '"')
		JSON.parse(jsonString)
		return true
	} catch (e) {
		return false
	}
}

// const logSeverity = {
// 	debug: { icon: "pi-question-circle", color: "gray-500" },
// 	info: { icon: "pi-info-circle", color: "cyan-500" },
// 	warning: { icon: "pi-exclamation-triangle", color: "orange-500" },
// 	error: { icon: "pi-exclamation-circle", color: "red-500" },
// }
</script>

<style scoped>
pre {
	margin: 0;
}
</style>
