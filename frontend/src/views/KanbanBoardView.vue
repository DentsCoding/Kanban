<script setup>
  import { ref, provide, watch } from 'vue' 
  import Column from './components/Column.vue';
  import Kanban from './views/KanbanBoardView.vue/index.js';

  const newTaskTitle = ref("")
  const newTaskDesc = ref("")
  const newTaskPriority = ref("medium")
  const dropAreas = ref([{title: "To Do", statusKey: "todo"}, {title: "In Progress", statusKey: "in-progress"}, {title: "Done", statusKey: "done"}])
  const tasks = ref(JSON.parse(localStorage.getItem("kanban_tasks")))

  const maxId = tasks.value.length > 0 ? Math.max(...tasks.value.map(t => t.id)) : 0
  const taskIndex = ref(maxId + 1)

  function addNewTask() {
    if (newTaskTitle.value.trim() === "") return
    tasks.value.push({
      id: taskIndex.value, 
      title: newTaskTitle.value, 
      description: newTaskDesc.value, 
      priority: newTaskPriority.value,
      status: "todo",
    })

    newTaskTitle.value = ""
    newTaskDesc.value = ""
    newTaskPriority.value = "medium"
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

          <select v-model="newTaskPriority" class="priority-select">
            <option value="low">Low</option>
            <option value="medium">Medium</option>
            <option value="high">High</option>
          </select>

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