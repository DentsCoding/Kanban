// routes
import { createRouter, createWebHistory } from 'vue-router'
import LandingView from '@/views/LandingView.vue'
import LoginPageView from '@/views/LoginPageView.vue'
import SignUpView from '@/views/SignUpView.vue'
import AppLayout from '@/views/AppLayout.vue'
import KanbanBoardView from '@/views/KanbanBoardView.vue'
import ProjectsView from '@/views/ProjectsView.vue'
const routes = [
    // Public Routes
    {
        path: "/",
        name: "landing",
        component: LandingView
    },
    {
        path: "/login",
        name: "login",
        component: LoginPageView
    },
    {
        path: "/signup",
        name: "signup",
        component: SignUpView
    },

    // Private authenticated routes
    {
        path: "/app",
        component: AppLayout,
        meta: { requiresAuth: true },
        children: [
            {
                path: "",
                redirect: { name: "board" }
            },
            {
                path: "board",
                name: "board",
                component: KanbanBoardView
            }, // TODO: Add paths for projects, tasks and statistics
            {
                path: "projects",
                name: "projects",
                component: ProjectsView
            }
        ]
    }
]

// router
const router = createRouter({
    history: createWebHistory(),
    routes
})


// before
router.beforeEach(to => {
    const token = localStorage.getItem("token")
    const isLoggedIn = !!token // TODO: Implement token validation
    const isAuthRequired = to.matched.some(p => p.meta.requiresAuth)
    if (!isLoggedIn && isAuthRequired) {

        return { name: "login" }
    }

    if ((to.name == "login" || to.name == "signup" || to.name == "landing") && isLoggedIn) {
        return { name: "board" }
    }
})


export default router