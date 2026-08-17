<script setup>
    import { useRouter } from 'vue-router'
    const router = useRouter()

    router.push({name: 'landing'})
</script>
<template>
  <div class="app-layout-wrapper">
    <!-- Top App Bar (Fixed across the very top, never moves) -->
    <header class="top-navbar">
      <div class="navbar-left">
        <button class="burger-btn" @click="toggleSidebar" aria-label="Toggle Navigation">
          <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="3" y1="12" x2="21" y2="12"></line>
            <line x1="3" y1="6" x2="21" y2="6"></line>
            <line x1="3" y1="18" x2="21" y2="18"></line>
          </svg>
        </button>
        <h2 class="current-section-title">{{ currentSectionName }}</h2>
      </div>

      <div class="navbar-right">
        <button class="logout-btn" @click="handleLogout">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path>
            <polyline points="16 17 21 12 16 7"></polyline>
            <line x1="21" y1="12" x2="9" y2="12"></line>
          </svg>
          Sign Out
        </button>
      </div>
    </header>

    <!-- App Body Container (Sits below the fixed navbar) -->
    <div class="app-body">
      <!-- Slide-out Sidebar Panel (Sits underneath top-navbar) -->
      <aside :class="['sidebar', { 'sidebar-open': isSidebarOpen }]">
        <nav class="sidebar-nav">
          <RouterLink :to="{ name: 'board' }" class="nav-item">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="7" height="9"></rect><rect x="14" y="3" width="7" height="5"></rect><rect x="14" y="12" width="7" height="9"></rect><rect x="3" y="16" width="7" height="5"></rect></svg>
            Board
          </RouterLink>

          <RouterLink :to="{ name: 'projects' }" class="nav-item">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path></svg>
            Projects (Soon)
          </RouterLink>
          <a href="#" class="nav-item disabled" @click.prevent>
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="20" x2="18" y2="10"></line><line x1="12" y1="20" x2="12" y2="4"></line><line x1="6" y1="20" x2="6" y2="14"></line></svg>
            Statistics (Soon)
          </a>
        </nav>
      </aside>

      <!-- Main Content Area (Smoothly adjusts margin when sidebar opens, pushing content inside) -->
      <main :class="['main-content-area', { 'sidebar-expanded': isSidebarOpen }]">
        <div class="page-view-container">
          <router-view></router-view>
        </div>
      </main>
    </div>
  </div>
</template>

<script>
export default {
  name: "AppLayout",
  data() {
    return {
      isSidebarOpen: false 
    };
  },
  computed: {
    currentSectionName() {
      const routeName = this.$route.name;
      if (!routeName) return "Dashboard";
      return routeName.charAt(0).toUpperCase() + routeName.slice(1);
    }
  },
  methods: {
    toggleSidebar() {
      this.isSidebarOpen = !this.isSidebarOpen;
    },
    handleLogout() {
      localStorage.removeItem("token");
      this.$router.push({ name: "landing" });
    }
  }
};
</script>

<style scoped>
.app-layout-wrapper {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  background-color: #f8fafc;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  color: #1e293b;
  position: relative;
  overflow-x: hidden;
}

/* Top Navbar (Fixed across the very top) */
.top-navbar {
  height: 70px;
  background-color: #ffffff;
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 2rem;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 100; /* Stays above everything */
}

.navbar-left {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.burger-btn {
  background: transparent;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
  padding: 0.4rem;
  cursor: pointer;
  color: #475569;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s;
}

.burger-btn:hover {
  background-color: #f1f5f9;
  color: #0f172a;
}

.current-section-title {
  font-size: 1.1rem;
  font-weight: 600;
  color: #1e293b;
  margin: 0;
}

.logout-btn {
  background: transparent;
  border: 1px solid #cbd5e1;
  color: #475569;
  padding: 0.5rem 0.85rem;
  border-radius: 6px;
  font-size: 0.85rem;
  font-weight: 600;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  transition: all 0.2s;
  font-family: inherit;
}

.logout-btn:hover {
  background-color: #fee2e2;
  border-color: #fca5a5;
  color: #b91c1c;
}

/* App Body Container */
.app-body {
  display: flex;
  margin-top: 70px; /* Offset content down by the fixed navbar height */
  flex-grow: 1;
  position: relative;
}

/* Sidebar Panel (Starts BELOW the top navbar, pushes content when open) */
.sidebar {
  width: 260px;
  background-color: #ffffff;
  border-right: 1px solid #e2e8f0;
  display: flex;
  flex-direction: column;
  position: fixed;
  top: 70px; /* Starts right below the top navbar */
  bottom: 0;
  left: -260px; /* Hidden completely off-screen by default */
  z-index: 90; /* Sits below top-navbar (100) */
  transition: left 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.sidebar.sidebar-open {
  left: 0; /* Slides into view underneath the fixed top bar */
}

.sidebar-nav {
  padding: 1.5rem 1rem;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem 1rem;
  border-radius: 6px;
  color: #475569;
  text-decoration: none;
  font-size: 0.9rem;
  font-weight: 500;
  transition: all 0.2s;
}

.nav-item:hover {
  background-color: #f1f5f9;
  color: #0f172a;
}

.nav-item.router-link-active {
  background-color: #f3e8ff; 
  color: #8e63f1; 
  font-weight: 600;
}

.nav-item.disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* Main Content Area (Pushes content right smoothly) */
.main-content-area {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  min-width: 0;
  margin-left: 0; /* Full width by default */
  transition: margin-left 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.main-content-area.sidebar-expanded {
  margin-left: 260px; /* Pushes content inside router-view aside */
}

.page-view-container {
  flex-grow: 1;
  padding: 2rem;
}

/* Mobile adjustments */
@media (max-width: 768px) {
  .main-content-area.sidebar-expanded {
    margin-left: 0; /* Fallback to overlay behavior on small mobile screens */
  }
}
</style>