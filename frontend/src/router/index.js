import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '@/pages/Login.vue'
import HomePage from '@/pages/Home.vue'
import RegisterPage from '@/pages/Register.vue'
import NavPage from '@/pages/Nav.vue'
import Profile from '@/pages/Profile.vue'
import Browser from '@/pages/Browser-test.vue'
import Recipe from '@/pages/Recipe.vue'
import RecipeMg from '@/pages/RecipeMg.vue'
import TestPage from '@/pages/TestPage.vue'

const routes = [
  //{ path: '/', name: 'Root', component: HomePage },
  { path: '/', name: 'HomePage', component: HomePage},
  { path: '/nav', name: 'Nav', component: NavPage },
  { path: '/recipemg', name:'RecipeMG', component: RecipeMg},
  { path: '/home', name: 'Home', component: HomePage },
  { path: '/login', name: 'Login', component: LoginPage },
  { path: '/register', name: 'Register', component: RegisterPage },
  { path: '/profile', name: 'Profile', component: Profile, meta: { requiresAuth: true } }, // 🔒
  { path: '/browser', name: 'Browser', component: Browser, meta: { requiresAuth: true } }, // 🔒
  { path: '/re/:id', name: 'Recipe', component: Recipe, meta: { requiresAuth: true } }, // 🔹 zmieniona nazwa
  { path: '/register', name: 'Register', component: RegisterPage },
  { path: '/test', component: TestPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

/*
router.beforeEach(async (to, from, next) => {
  if (to.meta.requiresAuth) {
    try {
      const resp = await verifyUser()
      if (resp.status == 401) {
        throw new Error("unauthorized")
      }
      next()
    } catch {
      next('/login')
    }
  } else {
    next()
  }
})
*/
export default router