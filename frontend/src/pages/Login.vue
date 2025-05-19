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

        <form class="w-full flex flex-col items-center">
          <div class="form-control mb-4 w-full">
            <input type="email" placeholder="Email" class="input input-bordered w-full" />
          </div>

          <div class="form-control mb-4 w-full">
            <input type="password" placeholder="Password" class="input input-bordered w-full" />
          </div>

          <div class="form-control mb-4 w-full flex justify-center">
          <router-link to="/profile" class="btn btn-primary w-full text-center">Log in</router-link>         
           </div>

          <div class="text-center text-sm">
            <span>Don't have an account?</span>
            <router-link to="/register" class="link link-primary ml-1">Sign up</router-link>
          </div>
          <!--
          <div class="text-center text-sm mt-2">
            <a class="link link-hover link-secondary font-semibold">Forgot password?</a>
          </div>
          -->
        </form>
      </div>
    </div>
  </div>

</template>

<script>
import html2canvas from 'html2canvas'

export default {
  name: 'LoginPage',
  methods: {
    captureImage() {
      html2canvas(this.$refs.captureTarget).then(canvas => {
        const link = document.createElement('a')
        link.download = 'log-dish.jpg'
        link.href = canvas.toDataURL('image/jpeg', 0.95)
        link.click()
      })
    }
  }
}
</script>

<style>
</style>