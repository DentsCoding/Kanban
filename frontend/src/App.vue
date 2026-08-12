<script setup>
  import { ref, provide, watch } from 'vue' 
  import Column from './components/Column.vue';

  const taskIndex = ref(3)
  const newTaskTitle = ref("")
  const newTaskDesc = ref("")
  const dropAreas = ref([{title: "To Do", statusKey: "todo"}, {title: "In Progress", statusKey: "in-progress"}, {title: "Done", statusKey: "done"}])
  const tasks = ref(JSON.parse(localStorage.getItem("kanban_tasks")))

function addNewTask() {
  if (newTaskTitle.value.trim() === "") return
  tasks.value.push({
    id: taskIndex.value, 
    title: newTaskTitle.value, 
    description: newTaskDesc.value, 
    status: "todo"})

  newTaskTitle.value = ""
  newTaskDesc.value = ""
  taskIndex.value++
}

function handleTaskDrop(newStatus) {
  const taskId = event.dataTransfer.getData('taskId')
  const task = tasks.value.find(t => t.id == taskId)
  if(task) {
    task.status = newStatus
  }

}


function deleteTask(taskId) {
  tasks.value = tasks.value.filter(t => t.id !== taskId)
}

function changeTitle(taskId, newTitle)
{
  const activeTask = tasks.value.find(t => t.id === taskId)
  activeTask.title = newTitle
}

function changeDescription(taskId, newDescription)
{
  const activeTask = tasks.value.find(t => t.id === taskId)
  activeTask.description = newDescription
}

  provide('deleteTask', deleteTask)
  provide('changeTitle', changeTitle)
  provide('changeDescription', changeDescription)

watch(tasks, (newTasks) => {
  localStorage.setItem("kanban_tasks", JSON.stringify(newTasks))
}, {deep: true})
</script>

<template>
  <div class="app-container">
    <header class="app-header">
      <h1>Kanban Board</h1>
    </header>

    <!-- Task Adding Form -->
    <section class="task-form-section">
      <form @submit.prevent="addNewTask" class="task-form">
        <input 
          v-model="newTaskTitle" 
          type="text" 
          placeholder="Task title..." 
          required 
        />
        <input 
          v-model="newTaskDesc" 
          type="text" 
          placeholder="Task description..." 
        />
        <button type="submit">Add Task</button>
      </form>
    </section>

    <!-- Columns Area -->
    <div class="board">
      <Column 
        v-for="area in dropAreas" 
        :columnTitle="area.title"
        :key="area.statusKey" 
        :statusKey="area.statusKey" 
        :tasks="tasks" 
        @task-dropped="handleTaskDrop"
      />
    </div>
  </div>
</template>

<style scoped>
.task-form-section {
  background: #ffffff;
  padding: 1rem 1.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
  margin-bottom: 2rem;
}

.task-form {
  display: flex;
  gap: 10px;
}

.task-form input {
  flex: 1;
  padding: 0.6rem 1rem;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  font-size: 0.95rem;
}

.task-form input:focus {
  outline: none;
  border-color: #42a1b8;
}

.task-form button {
  background-color: #42a1b8;
  color: white;
  border: none;
  padding: 0.6rem 1.2rem;
  font-weight: 600;
  border-radius: 4px;
  cursor: pointer;
  transition: background 0.2s;
}

.task-form button:hover {
  background-color: #35495e; /* Vue Dark Blue */
}

.board {
  display: flex;
  gap: 1.5rem;
  align-items: flex-start;
  overflow-x: auto;
}
</style>