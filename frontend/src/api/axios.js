import axios from "axios";

const api = axios.create({
    baseURL: "https://mealsfinder-backend-3ghs7wg7-euc.a.run.app",
    headers: {
        'Content-Type':'application/json'
    },
    withCredentials:true,
    paramsSerializer: function (params) {
    const searchParams = new URLSearchParams()
    for (const key in params) {
      const value = params[key]
      if (Array.isArray(value)) {
        value.forEach(val => searchParams.append(key, val)) // No key + []
      } else {
        searchParams.append(key, value)
      }
    }
    return searchParams.toString()
  }
})

export const getAllRecipes = (params) =>
  api.get('/browser', { params })

export const getRecipeById = (id) =>
  api.get(`/re/${id}`)

export async function submitRecipeRating({ stars, recipe_id }) {
  return await api.post(`/re`, { stars, recipe_id })
}

export async function getRatings(id) {
  return await api.get(`/ret/${id}`)
}

export async function getReview(id) {
  return await api.get(`/rev/${id}`)
}

export const verifyUser = () =>
  api.get('/verify') 

export const logoutUser = () =>
  api.get('/logout')

export const getTags = () =>
  api.get('/tags')

export const addRecipe = (data) =>
  api.post('/recipe', data)

export const addUserTags = (data) =>
  api.post('/user/tags', data)

export const deleteUserTags = (tagName) =>
  api.delete(`/user/tags/${tagName}`)

export const displayUserTags = () =>
  api.get('/user/tags')

export default api