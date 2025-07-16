import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Recommendations from '../views/Recommendations.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
    path: '/',
    redirect: '/home'
    },
    {
      path: '/home',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/recommendations',
      name: 'recommendations',
      component: Recommendations,
    },
  ],
})

export default router
