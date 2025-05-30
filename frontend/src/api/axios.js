import axios from "axios";

const api = axios.create({
    baseURL: "http://localhost:8080",
    headers: {
        'Content-Type':'application/json'
    },
    withCredentials:true,
})

export const getAllRecipes = (params) =>
  api.get('/browser', { params })

export const getRecipeById = (id) =>
  api.get(`/re/${id}`)

export const verifyUser = () =>
  api.get('/user/verify') 

export const logoutUser = () =>
  api.post('/user/logout')

export default api