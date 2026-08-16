<script setup>
import { computed, inject, nextTick, onMounted, onUnmounted, ref, watch } from 'vue'

const deleteTask = inject("deleteTask")
const changeTitle = inject("changeTitle")
const changeDescription = inject("changeDescription")

const props = defineProps({
    title: String,
    description: String,
    id: Number,
    priority: String,
    isDone: Boolean
})

const mousePosition = ref({x: 0, y: 0})
const showContextMenu = ref(false)
const isTitleEdited = ref(false)
const isDescEdited = ref(false)
const titleInputRef = ref(null)
const descInputRef = ref(null)
const titleEdit = ref(props.title)
const descEdit = ref(props.description)

watch(() => props.title, (newVal) => { titleEdit.value = newVal })
watch(() => props.description, (newVal) => { descEdit.value = newVal })

function onDragStart(event) {
    event.dataTransfer.setData("taskId", props.id)
    event.dataTransfer.effectAllowed = 'move'
}

function handleContextMenu(event) {
    event.preventDefault() 
    mousePosition.value = { x: event.clientX, y: event.clientY }
    showContextMenu.value = true
}

function closeContextMenu() {
    if(showContextMenu.value) showContextMenu.value = false
}

function handleDelete() {
    deleteTask(props.id)
    closeContextMenu()
}

function handleChangeTitle() {
    changeTitle(props.id, titleEdit.value)
    isTitleEdited.value = false
}

function cancelEditTitle() {
    titleEdit.value = props.title
    isTitleEdited.value = false
}

function handleChangeDescription() {
    changeDescription(props.id, descEdit.value)
    isDescEdited.value = false
}

function cancelEditDescription() {
    descEdit.value = props.description
    isDescEdited.value = false
}

async function editTitle() {
    if(!isTitleEdited.value) isTitleEdited.value = true
    await nextTick()
    titleInputRef.value?.focus()
}

async function editDescription() {
    if(!isDescEdited.value) isDescEdited.value = true
    await nextTick()
    descInputRef.value?.focus()
}

onMounted(() => window.addEventListener('click', closeContextMenu))
onUnmounted(() => window.removeEventListener('click', closeContextMenu))
</script>

<template>
    <div class="task-card" :class="{ 'completed-task': props.isDone }" :draggable="true" @dragstart="onDragStart" @contextmenu="handleContextMenu">

        <div class="task-header-row">
            <h3 v-if="!isTitleEdited" @dblclick="editTitle">{{ props.title }}</h3>
            <input 
                v-if="isTitleEdited" 
                ref="titleInputRef" 
                v-model="titleEdit" 
                type="text" 
                @mousedown.stop
                @focusout="handleChangeTitle" 
                @keyup.enter="handleChangeTitle" 
                @keyup.esc="cancelEditTitle">
            
            <span :class="['priority-badge', props.priority]">
                {{ props.priority }}
            </span>
        </div>

        <p v-if="!isDescEdited" @dblclick="editDescription">{{ props.description }}</p>
        <textarea 
            v-if="isDescEdited" 
            ref="descInputRef"
            v-model="descEdit" 
            spellcheck="false"
            rows="2"
            @mousedown.stop
            @focusout="handleChangeDescription"
            @keyup.enter.exact="handleChangeDescription"
            @keyup.esc="cancelEditDescription"
        ></textarea>
        
        <Teleport to="body">
            <div v-if="showContextMenu" class="context-menu" @click.stop :style="{top: `${mousePosition.y}px`, left: `${mousePosition.x}px`}">
                <button @click="handleDelete" class="menu-item delete">Delete Task</button>
            </div>
        </Teleport>
    </div>
</template>


<style scoped>

.task-card {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  padding: 0.85rem;
  box-shadow: 0 1px 3px rgba(0,0,0,0.04);
  cursor: grab;
  margin-bottom: 0.75rem;
  position: relative;
}

.task-card.completed-task {
  background: #f1f5f9;
  border-color: #cbd5e1;
  opacity: 0.75;
}

.task-card.completed-task h3 {
  text-decoration: line-through;
  color: #64748b;
}

.task-card.completed-task p {
  text-decoration: line-through;
  color: #94a3b8;
}

.task-card.completed-task .priority-badge {
  filter: grayscale(40%);
  opacity: 0.8;
}

.task-header-row {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 0.5rem;
  margin-bottom: 0.3rem;
}

.task-card h3 { 
  font-size: 1rem; 
  margin: 0; 
  color: #2d3748; 
  cursor: pointer;
  font-family: inherit;
}

.task-card p { 
  font-size: 0.875rem; 
  margin: 0; 
  color: #718096; 
  cursor: pointer;
  font-family: inherit;
  white-space: pre-wrap;
}

/* Inline Edit Inputs & Textareas */
.task-card input[type="text"],
.task-card textarea {
  font-family: inherit;
  font-size: 0.9rem;
  padding: 0.4rem 0.6rem;
  border: 1px solid #cbd5e1;
  border-radius: 4px;
  outline: none;
  background-color: #f8fafc;
  color: #334155;
  width: 100%;
  box-sizing: border-box;
  resize: none;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.task-card input[type="text"]:focus,
.task-card textarea:focus {
  border-color: #8e63f1;
  box-shadow: 0 0 0 2px rgba(142, 99, 241, 0.15);
  background-color: #ffffff;
}

/* Priority Badges */
.priority-badge {
  font-size: 0.65rem;
  font-weight: 700;
  text-transform: uppercase;
  padding: 2px 6px;
  border-radius: 4px;
  letter-spacing: 0.05em;
  flex-shrink: 0;
}

.priority-badge.low { background: #e0f2fe; color: #0369a1; }
.priority-badge.medium { background: #fef3c7; color: #b45309; }
.priority-badge.high { background: #fee2e2; color: #b91c1c; }

/* Context Menu */
.context-menu {
  position: fixed;
  background: white;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  z-index: 9999;
  padding: 4px;
  min-width: 140px;
}

.menu-item { 
  width: 100%; 
  text-align: left; 
  background: transparent; 
  border: none; 
  padding: 8px 12px; 
  border-radius: 4px; 
  font-size: 0.85rem; 
  font-weight: 500; 
  cursor: pointer; 
  font-family: inherit;
}

.menu-item.delete { color: rgb(224, 74, 74); }
.menu-item.delete:hover { background: #fee2e2; color: rgb(75, 22, 22); }

</style>