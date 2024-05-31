import { createRouter, createWebHistory } from 'vue-router'
import MainView from '../views/MainView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: MainView
    }
  ]
})

export default router
