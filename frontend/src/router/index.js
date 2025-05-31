import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '@/pages/Login.vue'
import HomePage from '@/pages/Home.vue'
import RegisterPage from '@/pages/Register.vue'
import NavPage from '@/pages/Nav.vue'
import Profile from '@/pages/Profile.vue'
import BrowserTest from '@/pages/Browser-test.vue'
import Recipe from '@/pages/Recipe.vue'
import { verifyUser } from '@/api/axios' 

const routes = [
  { path: '/', name: 'Root', component: HomePage },
  { path: '/nav', name: 'Nav', component: NavPage },
  { path: '/home', name: 'Home', component: HomePage },
  { path: '/login', name: 'Login', component: LoginPage },
  { path: '/register', name: 'Register', component: RegisterPage },
  { path: '/profile', name: 'Profile', component: Profile, meta: { requiresAuth: true } }, // 🔒
  { path: '/browser', name: 'CategoryBrowser', component: BrowserTest, meta: { requiresAuth: true } }, // 🔒
  { path: '/re/:id', name: 'Recipe', component: Recipe } // 🔹 zmieniona nazwa
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach(async (to, from, next) => {
  if (to.meta.requiresAuth) {
    try {
      await verifyUser()
      next()
    } catch {
      next('/login')
    }
  } else {
    next()
  }
})

export default router