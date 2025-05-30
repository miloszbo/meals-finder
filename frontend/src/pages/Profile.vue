<template>
  <div class="bg-base-200 text-white h-[1080px] overflow-hidden flex flex-col">
    <div class="h-[80px] shrink-0">
      <Navbar />
    </div>

    <div class="flex-grow flex h-[calc(1080px-300px)]">
      <!-- LeftBar -->
      <aside class="w-64 bg-base-100 border-r border-base-content p-4 flex flex-col items-start">
        <div class="w-full space-y-1 mt-2">
          <button
            v-for="item in sidebarItems"
            :key="item.name"
            @click="tab = item.tab"
            :class="[
              'w-full flex items-center gap-2 px-4 py-2 rounded-full text-left transition',
              tab === item.tab ? 'bg-primary text-white' : 'hover:bg-base-300'
            ]"
          >
            <component :is="item.icon" class="w-5 h-5" />
            <span class="font-medium">{{ item.name }}</span>
          </button>
        </div>
      </aside>

      <main class="flex-1 p-6 space-y-6 overflow-hidden">
        <!-- Main -->
        <div v-if="tab === 'profile'" class="flex flex-col items-center justify-center h-full">
          <div class="avatar mb-4">
            <div class="w-36 rounded-full ring ring-primary ring-offset-base-100 ring-offset-2">
              <img src="https://i.pravatar.cc/300" alt="User avatar" />
            </div>
          </div>
          <h2 class="text-2xl font-bold mb-2">Username</h2>
          <button class="btn btn-primary btn-sm">Edit profile</button>
        </div>

        <div v-else-if="tab === 'settings'" class="max-w-lg space-y-4 overflow-y-auto">
          <h2 class="text-xl font-bold">Dane osobowe</h2>
          <label class="form-control">
            <span class="label-text text-white">Imię i nazwisko</span>
            <input type="text" class="input input-bordered w-full" placeholder="Wprowadź imię" />
          </label>
          <label class="form-control">
            <span class="label-text text-white">Data urodzenia</span>
            <input type="date" class="input input-bordered w-full" />
          </label>
          <button class="btn btn-success mt-4">Zapisz</button>
        </div>

        <!-- Pozostałe sekcje... -->
        <div v-else-if="tab === 'privacy'" class="space-y-2">
          <h2 class="text-xl font-bold">Dane i prywatność</h2>
          <p class="text-gray-400">Tutaj możesz zarządzać danymi powiązanymi z Twoim kontem.</p>
        </div>
        <div v-else-if="tab === 'security'" class="space-y-2">
          <h2 class="text-xl font-bold">Bezpieczeństwo</h2>
          <p class="text-gray-400">Ustaw hasło, uwierzytelnianie 2FA i inne zabezpieczenia konta.</p>
        </div>
        <div v-else-if="tab === 'sharing'" class="space-y-2">
          <h2 class="text-xl font-bold">Osoby i udostępnianie</h2>
          <p class="text-gray-400">Zarządzaj użytkownikami, którzy mogą przeglądać Twoje dane lub przepisy.</p>
        </div>
        <div v-else-if="tab === 'payments'" class="max-w-2xl space-y-4">
          <h2 class="text-xl font-bold">Płatności i subskrypcje</h2>
          <table class="table w-full text-white">
            <thead>
              <tr>
                <th>Plan</th>
                <th>Cena</th>
                <th>Funkcje</th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>Free</td>
                <td>0 zł</td>
                <td>Ograniczony dostęp do przepisów</td>
                <td><button class="btn btn-sm btn-disabled">Obecny</button></td>
              </tr>
              <tr>
                <td>Pro</td>
                <td>19,99 zł/mies.</td>
                <td>Nieograniczone przepisy, zapisane dania</td>
                <td><button class="btn btn-sm btn-primary">Wybierz</button></td>
              </tr>
            </tbody>
          </table>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import Navbar from '@/components/Navbar.vue'
import { ref } from 'vue'
import { User, Settings, Lock, Users, CreditCard } from 'lucide-vue-next'



const sidebarItems = [
  { name: 'Strona główna', tab: 'profile', icon: User },
  { name: 'Dane osobowe', tab: 'settings', icon: Settings },
  { name: 'Dane i prywatność', tab: 'privacy', icon: Settings },
  { name: 'Bezpieczeństwo', tab: 'security', icon: Lock },
  { name: 'Osoby i udostępnianie', tab: 'sharing', icon: Users },
  { name: 'Płatności i subskrypcje', tab: 'payments', icon: CreditCard }
]


import { useRoute } from 'vue-router'
import { onMounted } from 'vue'

const route = useRoute()
const tab = ref('profile')

onMounted(() => {
  const tabQuery = route.query.tab
  if (tabQuery && typeof tabQuery === 'string') {
    tab.value = tabQuery
  }
})

</script>