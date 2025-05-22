import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '@/pages/Login.vue'
import HomePage from '@/pages/Home.vue'
import RegisterPage from '@/pages/Register.vue'
import NavPage from '@/pages/Nav.vue'
import CategoryBrowser from '@/pages/CategoryBrowser.vue'
import Profile from '@/pages/Profile.vue'

const routes = [
   { path: '/', name: 'Root', component: HomePage },
  { path: '/nav', name: 'Nav', component: NavPage },
  { path: '/home', name: 'Home', component: HomePage },
  { path: '/login', name: 'Login', component: LoginPage },
  { path: '/register', name: 'Register', component: RegisterPage },
  { path: '/profile', name: 'Profile', component: Profile },
  { path: '/browser', name: 'CategoryBrowser', component: CategoryBrowser }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router