<template>
  <div class="flex h-screen">
    <!-- Left panel -->
    <div class="w-1/2 hidden lg:block">
      <img
        src="//assets/login_register dish.png"
        alt="Delicious food"
        class="object-cover w-full h-full"
      />
    </div>

    <!-- Right panel -->
    <div class="w-full lg:w-1/2 flex items-center justify-center bg-base-200">
      <div class="w-full max-w-md p-8 flex flex-col items-center justify-center">

        <router-link to="/home" class="text-4xl font-bold mb-8 hover:text-primary transition-colors duration-200">
          MealsFinder
        </router-link>

        <form class="w-full flex flex-col items-center" @submit.prevent="loginUser">
          <div class="form-control mb-4 w-full">
            <input
              type="text"
              placeholder="Login"
              class="input input-bordered w-full"
              v-model="form.login"
              required
            />
          </div>

          <div class="form-control mb-4 w-full">
            <input
              type="password"
              placeholder="Password"
              class="input input-bordered w-full"
              v-model="form.password"
              required
            />
          </div>

          <div class="form-control mb-4 w-full">
            <button type="submit" class="btn btn-primary w-full">Log in</button>
          </div>

          <div class="text-center text-sm">
            <span>Don't have an account?</span>
            <router-link to="/register" class="link link-primary ml-1">Sign up</router-link>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import api from '@/api/axios'

export default {
  name: 'LoginPage',
  data() {
    return {
      form: {
        login: '',
        password: ''
      }
    }
  },
  methods: {
    async loginUser() {
      try {
        const res = await api.post('/user/login', {
          login: this.form.login,
          password: this.form.password
        })

      
        if (res.status === 200) {
          const token = res.data.token
          if (token) {
            localStorage.setItem('token', token)
          }

          alert('Zalogowano pomyślnie!')
          this.$router.push('/profile')
        }
      } catch (err) {
        console.error('Błąd logowania:', err)
        alert('Nieprawidłowy login lub hasło.')
      }
    }
  }
}
</script>