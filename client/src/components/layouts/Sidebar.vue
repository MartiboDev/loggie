<template>
	<div class="sidebar-layout">
		<RouterLink v-for="page of pages" :to="page.path" class="nav-button">
			<Button
				:label="page.name"
				:icon="`pi ${page.icon}`"
				:severity="activePage === page.path ? 'primary' : 'secondary'"
				text
			/>
		</RouterLink>
	</div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from "vue"
import { RouterLink, useRoute } from "vue-router"

const activePage = ref<string>("")

const setActivePage = (page: string) => {
	activePage.value = page
}

const pages = [
	{
		name: "Logs",
		icon: "pi pi-book",
		path: "/",
	},
]

const route = useRoute()

onMounted(() => {
	setActivePage(route.path)
})

watch(
	() => route.path,
	(newPath) => {
		setActivePage(newPath)
	}
)
</script>

<style scoped lang="scss">
.sidebar-layout {
	display: flex;
	flex-direction: column;
	gap: 0.25rem;

	background-color: var(--layout-background-color);
	border-right: 1px solid var(--border-color);
	height: 100%;
	padding: 0.5rem 1rem;
}

.nav-button {
	display: block;

	button {
		width: 100%;
	}
}
</style>
