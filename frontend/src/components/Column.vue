<script setup>
import Task from './Task.vue';
import { computed } from 'vue'

const props = defineProps({
    columnTitle: String,
    statusKey: String,
    tasks: Array
})

const emit = defineEmits(['task-dropped'])

const columnTasks = computed(() => {
    return props.tasks.filter(t => t.status === props.statusKey)
})


function onDrop() {
    emit("task-dropped", props.statusKey)
}
</script>

<template>
    <section class="column" @dragover.prevent @drop="onDrop">
        <h2>{{ props.columnTitle }}</h2>
        
        <div class="task-list">
            <Task 
                v-for="task in columnTasks" 
                :key="task.id" 
                :title="task.title" 
                :description="task.description" 
                :id="task.id"
            />
            <p v-if="columnTasks.length === 0" class="empty-text">No tasks here</p>
        </div>
    </section>
</template>

<!-- SCOPED CSS: Applies ONLY to Column.vue -->
<style scoped>
.column {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  width: 320px;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  max-height: 70vh; /* Limits height so it becomes scrollable */
  box-shadow: 0 2px 4px rgba(0,0,0,0.02);
}

.column h2 {
  font-size: 1.1rem;
  margin: 0;
  padding: 1rem;
  border-bottom: 1px solid #edf2f7;
  color: #35495e;
  background: #fafbfc;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
}

.task-list {
  padding: 1rem;
  overflow-y: auto; /* Makes the column scrollable when cards overflow */
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  flex: 1;
}

.empty-text {
  color: #a0aec0;
  font-size: 0.9rem;
  text-align: center;
  margin: 1rem 0;
}
</style>