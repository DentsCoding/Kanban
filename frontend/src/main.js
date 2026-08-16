import { compile, createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import './style.css'
import App from './App.vue'
import LoginPage from '@/views/LoginPageView.vue/'
import Kanban from '@/views/KanbanBoardView.vue' // move components into board


const routes = [
    { path: '/login', component: LoginPage },
    { path: '/board', component: Kanban }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

const app = createApp(App)
app.use(router)
app.mount('#app')
