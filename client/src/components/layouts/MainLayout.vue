<template>
	<div class="main-layout">
		<div class="main-navbar-layout">
			<Navbar />
		</div>
		<div class="main-sidebar-layout">
			<Sidebar />
		</div>
		<div class="main-content-layout">
			<PageTitle />
			<slot></slot>
		</div>
	</div>
</template>

<script setup lang="ts"></script>

<style scoped>
.main-layout {
	display: grid;
	grid-template-columns: 240px 1fr;
	grid-template-rows: auto 1fr;
	grid-template-areas:
		"navbar navbar"
		"sidebar content";

	height: 100vh;
}

.main-navbar-layout {
	grid-area: navbar;
}

.main-sidebar-layout {
	grid-area: sidebar;
}

.main-content-layout {
	grid-area: content;
}
</style>
