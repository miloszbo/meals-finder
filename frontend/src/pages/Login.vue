<template>
  <div class="flex min-h-screen">
    <!-- Left panel with image -->
    <div class="hidden lg:block w-1/2 h-screen">
      <img
        src="/images/login_dish.png"
        alt="Delicious food"
        class="object-cover w-full h-full"
      />
    </div>
    <!-- Right panel with login form -->
    <div class="flex flex-col justify-center items-center w-full lg:w-1/2 h-screen bg-[#111921]">
      <div class="w-full max-w-md p-8 flex flex-col items-center justify-center">
        <div class="text-4xl font-bold mb-8 text-white">MealsFinder</div>
        <form class="w-full flex flex-col items-center" @submit.prevent="loginUser">
          <div class="form-control mb-4 w-full">
            <input
              type="text"
              placeholder="Email"
              class="input input-bordered w-full bg-[#232a34] text-white border border-[#3d4451] focus:outline-none focus:ring-2 focus:ring-blue-500"
              v-model="form.login"
              required
            />
          </div>
          <div class="form-control mb-4 w-full">
            <input
              type="password"
              placeholder="Password"
              class="input input-bordered w-full bg-[#232a34] text-white border border-[#3d4451] focus:outline-none focus:ring-2 focus:ring-blue-500"
              v-model="form.password"
              required
            />
          </div>
          <div class="form-control mb-4 w-full">
            <button type="submit" class="btn w-full bg-[#466aff] hover:bg-[#3d5ed1] text-white font-bold">
              Log in
            </button>
          </div>
          <div class="text-center text-sm text-white w-full">
            <div>
              <span>Don't have an account?</span>
              <router-link to="/register" class="ml-1 text-[#466aff] hover:underline">Sign up</router-link>
            </div>
            <div class="mt-2 font-bold text-xs hover:underline cursor-pointer">
              Forgot password?
            </div>
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
