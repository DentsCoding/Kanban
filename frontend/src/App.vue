

<template>
  <Kanban />
</template>

<style scoped>
/* App Container and Layout */
.app-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem 1rem;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  color: #1e293b;
  background-color: #f8fafc;
  min-height: 100vh;
}

.app-header {
  margin-bottom: 2rem;
  text-align: center;
}

.app-header h1 {
  font-size: 2rem;
  font-weight: 700;
  color: #0f172a;
  margin: 0;
}

/* Board Layout */
.board {
  display: flex;
  gap: 1.5rem;
  overflow-x: auto;
  padding-bottom: 1rem;
  align-items: flex-start;
}

/* Form and Dropdown Styling */
.task-form-section {
  margin-bottom: 2rem;
}

.task-form {
  display: flex;
  gap: 0.75rem;
  background: white;
  padding: 1rem;
  border-radius: 8px;
  border: 1px solid #e2e8f0;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.04);
}

.task-form input,
.priority-select {
  padding: 0.6rem 0.8rem;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
  font-size: 0.9rem;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
  background-color: #ffffff;
  color: #334155;
}

.task-form input {
  flex-grow: 1;
}

.task-form input:focus,
.priority-select:focus {
  border-color: #8e63f1; 
  box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.15); /* Soft indigo glow */
}

/* Fixed Dropdown Styling with foolproof arrow */
.priority-select {
  cursor: pointer;
  appearance: none;
  -webkit-appearance: none;
  -moz-appearance: none;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='16' height='16' viewBox='0 0 24 24' fill='none' stroke='%2364748b' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpolyline points='6 9 12 15 18 9'%3E%3C/polyline%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 0.75rem center;
  background-size: 1rem;
  padding-right: 2.5rem;
}

.task-form button {
  background: #8e63f1; 
  color: white;
  border: none;
  padding: 0.6rem 1.25rem;
  border-radius: 6px;
  font-weight: 600;
  font-size: 0.9rem;
  cursor: pointer;
  transition: background-color 0.2s;
}

.task-form button:hover {
  background: #5f2aa5;; 
}
</style>