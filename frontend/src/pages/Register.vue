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
          <!-- Etap 1 -->
          <template v-if="step === 1">
            <div class="form-control mb-4 w-full">
              <input type="text" placeholder="Username" class="input input-bordered w-full" v-model="form.username" />
            </div>

            <div class="form-control mb-4 w-full">
              <input type="email" placeholder="Email" class="input input-bordered w-full" v-model="form.email" />
            </div>

            <div class="form-control mb-4 w-full">
              <input type="password" placeholder="Password" class="input input-bordered w-full" v-model="form.password" />
            </div>

            <div class="form-control mb-4 w-full">
              <input type="password" placeholder="Confirm Password" class="input input-bordered w-full" v-model="form.confirmPasswordpassword" />
            </div>

            <div class="form-control mb-4 w-full flex justify-center">
              <button type="button" class="btn btn-primary w-full" @click="step = 2">Next</button>
            </div>
          </template>

          <!-- Etap 2 -->
          <template v-else>
            <div class="form-control mb-4 w-full">
              <input type="text" placeholder="Phone Number" class="input input-bordered w-full" v-model="form.phoneNumber" />
            </div>

            <div class="form-control mb-4 w-full">
              <input type="text" placeholder="Age" class="input input-bordered w-full" v-model="age" />
            </div>

            <div class="form-control mb-4 w-full">
              <input type="text" placeholder="Sex" class="input input-bordered w-full" v-model="sex" />
            </div>

            <div class="form-control mb-4 w-full">
              <button type="submit" class="btn btn-success w-full">Sign up</button>
            </div>
          </template>

          

          <div class="text-center text-sm">
            <span>Have an account?</span>
            <router-link to="/Login" class="link link-primary ml-1">Sign in</router-link>
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
export default {
  name: 'RegisterPage',
  data() {
    return {
      step: 1,
      form: {
        username: '',
        email: '',
        password: '',
        confirmPassword: '',
        sex: '',
        age: '',
        phoneNumber: ''
      }
    }
  }
}
</script>