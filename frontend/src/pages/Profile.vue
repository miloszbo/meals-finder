<template>
  <div v-if="isAuthorized">
    <div class="flex min-h-screen bg-base-200 text-white">
      <!-- Sidebar -->
      <aside class="w-64 bg-base-100 border-r border-base-content p-4">
        <ul class="menu">
          <li><h2 class="menu-title">User Settings</h2></li>
          <li><a>Payments</a></li>
          <li><a>Saved Recipes</a></li>
        </ul>
      </aside>

      <!-- Main Profile Area -->
      <main class="flex-1 p-8 flex flex-col items-center bg-base-200">
        <!-- Avatar -->
        <div class="avatar mb-4">
          <div class="w-36 rounded-full ring ring-primary ring-offset-base-100 ring-offset-2">
            <img src="https://i.pravatar.cc/300" alt="User avatar" />
          </div>
        </div>

        <!-- Username -->
        <h2 class="text-2xl font-bold mb-2">Username</h2>

        <!-- Edit Button -->
        <button class="btn btn-primary btn-sm mb-6">Edit profile</button>

        <!-- Editable Fields -->
        <div class="w-full max-w-sm space-y-4">
          <label class="form-control w-full">
            <div class="label">
              <span class="label-text text-white">Name</span>
            </div>
            <input type="text" class="input input-bordered w-full" placeholder="Enter your name" />
          </label>

          <label class="form-control w-full">
            <div class="label">
              <span class="label-text text-white">Bio</span>
            </div>
            <textarea class="textarea textarea-bordered w-full" placeholder="Add a bio"></textarea>
          </label>
        </div>
      </main>
    </div>
  </div>
  <div v-else>
    You are unauthorized
  </div>
</template>

<script>
import api from '@/api/axios'

export default {
  name: 'ProfilePage',
  data() {
    return {
      isAuthorized: false,
    };
  },
  created(){
    this.checkAuthorization();
  },
  methods: {
    async checkAuthorization() {
      try {
        const response = await api.get('/profile')
        this.isAuthorized = response.status !== 401
        console.log(response.data)
      } catch (error) {
        console.error('Authorization check failed:', error);
        this.Authorization = false;
      }
    }
  }
};
</script>