<script setup>
import Task from '@/components/kanban/Task.vue';
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

function onDrop(event) {
    emit("task-dropped", props.statusKey)
}
</script>

<template>
    <section class="column" @dragover.prevent @drop="onDrop">
        <div class="column-header">
            <h2>{{ props.columnTitle }}</h2>
            <span class="task-count">{{ columnTasks.length }}</span>
        </div>
        
        <div class="task-list">
            <TransitionGroup name="task-list">
                <Task 
                    v-for="task in columnTasks" 
                    :key="task.id" 
                    :title="task.title" 
                    :description="task.description" 
                    :id="task.id"
                    :priority="task.priority"
                    :isDone="statusKey === 'done'"
                />
            </TransitionGroup>

            <div v-if="columnTasks.length === 0" class="empty-drop-zone">
                <span>Drop tasks here</span>
            </div>
        </div>
    </section>
</template>

<style scoped>

.column {
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  min-height: 400px;
  width: 300px; 
  flex-shrink: 0;
}

.task-list {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  flex-grow: 1;
  position: relative;
  width: 100%; 
}

.empty-drop-zone {
  border: 2px dashed #cbd5e1;
  border-radius: 6px;
  padding: 2rem 1rem;
  text-align: center;
  color: #94a3b8;
  font-size: 0.85rem;
  background: rgba(248, 250, 252, 0.5);
  width: 100%;
  box-sizing: border-box; 
}

.column-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.column-header h2 {
  font-size: 1rem;
  font-weight: 600;
  color: #334155;
  margin: 0;
}

.task-count {
  background: #e2e8f0;
  color: #475569;
  font-size: 0.75rem;
  font-weight: 600;
  padding: 2px 8px;
  border-radius: 999px;
}

/* Transition Group Animation Rules */
.task-list-move,
.task-list-enter-active,
.task-list-leave-active {
  transition: all 0.3s ease;
}

.task-list-enter-from,
.task-list-leave-to {
  opacity: 0;
  transform: translateY(15px);
}

.task-list-leave-active {
  position: absolute;
  width: 100%;
}

</style>