<template>
	<div class="header-layout">
		<i class="icon pi pi-search"></i>
		<span class="title">Loggie</span>
	</div>
</template>

<script setup lang="ts"></script>

<style scoped>
.header-layout {
	display: flex;
	align-items: center;
	gap: 0.5rem;

	background-color: var(--layout-background-color);
	border-bottom: 1px solid var(--border-color);
	padding: 1rem;
}

.icon {
	font-size: 1.25rem;
}

.title {
	font-size: 1.5rem;
	font-weight: bold;
}
</style>
