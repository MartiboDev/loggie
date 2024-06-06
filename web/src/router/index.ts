import { createRouter, createWebHistory } from "vue-router"
import LogView from "../views/LogView.vue"

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: "/",
			name: "Log Explorer",
			component: LogView,
		},
	],
})

export default router
