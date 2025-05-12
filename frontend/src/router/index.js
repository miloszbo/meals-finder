import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '@/pages/Login.vue'
import HomePage from '@/pages/Home.vue'
import RegisterPage from '@/pages/Register.vue'
import NavPage from '@/pages/Nav.vue'

const routes = [
  { path: '/', name: 'Nav', component: NavPage },
  { path: '/', name: 'Home', component: HomePage },
  { path: '/login', name: 'Login', component: LoginPage },
  { path: '/register', name: 'Register', component: RegisterPage }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router