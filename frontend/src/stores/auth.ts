import { defineStore } from 'pinia'
import axios from 'axios'

interface RegisterData {
  email: string
  password: string
  firstName: string
  lastName: string
  username: string
}

interface LoginData {
  email: string
  password: string
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as any,
    token: localStorage.getItem('token'),
    isLoading: false,
    error: null as string | null
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    currentUser: (state) => state.user
  },

  actions: {
    async register(data: RegisterData) {
      try {
        this.isLoading = true
        const response = await axios.post('/api/auth/register', data)
        this.token = response.data.token
        localStorage.setItem('token', this.token)
        await this.fetchUser()
      } catch (error: any) {
        this.error = error.response?.data?.message || 'Ошибка регистрации'
        throw error
      } finally {
        this.isLoading = false
      }
    },

    async login(data: LoginData) {
      try {
        this.isLoading = true
        const response = await axios.post('/api/auth/login', data)
        this.token = response.data.token
        localStorage.setItem('token', this.token)
        await this.fetchUser()
      } catch (error: any) {
        this.error = error.response?.data?.message || 'Ошибка входа'
        throw error
      } finally {
        this.isLoading = false
      }
    },

    async fetchUser() {
      try {
        const response = await axios.get('/api/auth/me', {
          headers: {
            Authorization: `Bearer ${this.token}`
          }
        })
        this.user = response.data
      } catch (error) {
        this.logout()
        throw error
      }
    },

    logout() {
      this.user = null
      this.token = null
      localStorage.removeItem('token')
    }
  }
}) 