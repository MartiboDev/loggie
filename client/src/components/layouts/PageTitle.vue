<template>
	<div class="page-title">
		<span class="title">{{ route.name }}</span>
	</div>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router"

const route = useRoute()
</script>

<style scoped>
.page-title {
	border-bottom: 1px solid var(--border-color);
	padding: 1rem 1.5rem;
}

.title {
	font-size: 1.1rem;
	font-weight: 700;
}
</style>
