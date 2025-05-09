import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '@/pages/Login.vue'
import HomePage from '@/pages/Home.vue'

const routes = [
  { path: '/', name: 'Home', component: HomePage },
  { path: '/login', name: 'Login', component: HomePage },
  { path: '/login', name: 'Login', component: LoginPage },
  { path: '/login', name: 'Login', component: RegisterPage },

]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router