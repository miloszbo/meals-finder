import axios from "axios";

const api = axios.create({
    baseURL: "http://localhost:8080",
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

export const verifyUser = () =>
  api.get('/verify') 

export const logoutUser = () =>
  api.get('/logout')

export const getTags = () =>
  api.get('/tags')

export const addRecipe = (formData) =>
  api.post('/recipes', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })

export const addUserTags = (data) =>
  api.post('/user/tags', data)

export const deleteUserTags = (tagName) =>
  api.delete(`/user/tags/${tagName}`)

export const displayUserTags = () =>
  api.get('/user/tags')

export default api