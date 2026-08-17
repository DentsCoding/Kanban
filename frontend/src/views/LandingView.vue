<script setup>
    import { useRouter } from 'vue-router';

    const router = useRouter()
    function signInHandler() {
        localStorage.setItem("token", "exists")

        router.push({ name: 'board'})
    } 
</script>
<template>
  <div class="app-container">
    <!-- Top Navbar -->
    <header class="navbar">
      <div class="navbar-brand">
        <h1>KanbanBoard</h1>
      </div>
      <nav class="navbar-links">
        <button class="nav-btn secondary" @click="signInHandler">Sign In</button>
        <button class="nav-btn primary" @click="$emit('navigate', 'signup')">Sign Up</button>
      </nav>
    </header>

    <!-- Hero Section -->
    <main class="hero-section">
      <div class="hero-content">
        <h2>Organize your workflow with absolute ease</h2>
        <p>
          A minimalist Kanban board designed to keep you focused, track your daily tasks, 
          and streamline your productivity without the clutter.
        </p>
        <button class="cta-button" @click="emit('navigate', 'signup')">Get Started</button>
      </div>
    </main>
  </div>
</template>

<style scoped>
.app-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem 1rem;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  color: #1e293b;
  background-color: #f8fafc;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

/* Navbar */
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 0;
  margin-bottom: 4rem;
  border-bottom: 1px solid #e2e8f0;
}

.navbar-brand h1 {
  font-size: 1.5rem;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
}

.navbar-links {
  display: flex;
  gap: 0.75rem;
}

.nav-btn {
  padding: 0.5rem 1rem;
  border-radius: 6px;
  font-weight: 600;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.2s;
  font-family: inherit;
}

.nav-btn.secondary {
  background: transparent;
  color: #475569;
  border: 1px solid #cbd5e1;
}

.nav-btn.secondary:hover {
  background: #f1f5f9;
  color: #0f172a;
}

.nav-btn.primary {
  background: #8e63f1;
  color: white;
  border: none;
}

.nav-btn.primary:hover {
  background: #5f2aa5;
}

/* Hero Section */
.hero-section {
  flex-grow: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  padding: 2rem 0;
}

.hero-content {
  max-width: 600px;
}

.hero-content h2 {
  font-size: 2.75rem;
  font-weight: 800;
  color: #0f172a;
  margin: 0 0 1.25rem 0;
  line-height: 1.2;
}

.hero-content p {
  font-size: 1.1rem;
  color: #64748b;
  margin: 0 0 2rem 0;
  line-height: 1.6;
}

.cta-button {
  background: #8e63f1;
  color: white;
  border: none;
  padding: 0.85rem 2rem;
  border-radius: 8px;
  font-weight: 600;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.2s, transform 0.1s;
  box-shadow: 0 4px 12px rgba(142, 99, 241, 0.25);
}

.cta-button:hover {
  background: #5f2aa5;
}

.cta-button:active {
  transform: translateY(1px);
}
</style>