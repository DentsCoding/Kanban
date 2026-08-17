<template>
  <div class="app-container">
    <header class="app-header">
      <h1 @click="$emit('navigate', 'welcome')" class="brand-link">KanbanBoard</h1>
    </header>

    <main class="auth-main">
      <div class="auth-card">
        <div class="auth-header">
          <h2>Create an account</h2>
          <p>Get started with your free Kanban board.</p>
        </div>

        <form @submit.prevent="handleSignup" class="auth-form">
          <div class="input-group">
            <label for="username">Username</label>
            <input 
              id="username"
              v-model="username" 
              type="text" 
              placeholder="Choose a username" 
              required 
            />
          </div>

          <div class="input-group">
            <label for="email">Email</label>
            <input 
              id="email"
              v-model="email" 
              type="email" 
              placeholder="you@example.com" 
              required 
            />
          </div>

          <div class="input-group">
            <label for="password">Password</label>
            <input 
              id="password"
              v-model="password" 
              type="password" 
              placeholder="Create a password" 
              required 
            />
          </div>

          <button type="submit" class="auth-submit-btn">Sign Up</button>
        </form>

        <div class="auth-footer">
          <p>Already have an account? <a href="#" @click.prevent="$emit('navigate', 'login')">Sign in</a></p>
        </div>
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

.app-header {
  margin-bottom: 2rem;
  text-align: center;
}

.brand-link {
  font-size: 1.5rem;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
  cursor: pointer;
  display: inline-block;
}

.auth-main {
  flex-grow: 1;
  display: flex;
  align-items: center;
  justify-content: center;
}

.auth-card {
  background: white;
  padding: 2.5rem;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05), 0 2px 4px -1px rgba(0, 0, 0, 0.03);
  width: 100%;
  max-width: 420px;
}

.auth-header {
  margin-bottom: 1.75rem;
  text-align: center;
}

.auth-header h2 {
  font-size: 1.5rem;
  font-weight: 700;
  color: #0f172a;
  margin: 0 0 0.5rem 0;
}

.auth-header p {
  font-size: 0.9rem;
  color: #64748b;
  margin: 0;
}

.auth-form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.input-group label {
  font-size: 0.85rem;
  font-weight: 600;
  color: #334155;
}

.input-group input {
  padding: 0.6rem 0.8rem;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
  font-size: 0.9rem;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
  background-color: #ffffff;
  color: #334155;
  font-family: inherit;
}

.input-group input:focus {
  border-color: #8e63f1; 
  box-shadow: 0 0 0 3px rgba(142, 99, 241, 0.15);
}

.auth-submit-btn {
  background: #8e63f1; 
  color: white;
  border: none;
  padding: 0.75rem;
  border-radius: 6px;
  font-weight: 600;
  font-size: 0.95rem;
  cursor: pointer;
  transition: background-color 0.2s;
  margin-top: 0.5rem;
}

.auth-submit-btn:hover {
  background: #5f2aa5; 
}

.auth-footer {
  margin-top: 1.5rem;
  text-align: center;
  font-size: 0.875rem;
  color: #64748b;
}

.auth-footer a {
  color: #8e63f1;
  text-decoration: none;
  font-weight: 600;
}

.auth-footer a:hover {
  text-decoration: underline;
}
</style>