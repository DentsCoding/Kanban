<script setup>
import { inject, nextTick, onMounted, onUnmounted, ref, watch } from 'vue'

const deleteTask = inject("deleteTask")
const changeTitle = inject("changeTitle")
const changeDescription = inject("changeDescription")

const props = defineProps({
    title: String,
    description: String,
    id: Number
})

const mousePosition = ref({x: 0, y: 0})
const showContextMenu = ref(false)
const isTitleEdited = ref(false)
const isDescEdited = ref(false)
const titleInputRef = ref(null)
const descInputRef = ref(null)
const titleEdit = ref(props.title)
const descEdit = ref(props.description)
const isDraggable = ref(true)

function onDragStart(event) {
    event.dataTransfer.setData("taskId", props.id)
    event.dataTransfer.effectAllowed = 'move'
}

function handleContextMenu(event) {
    event.preventDefault() 

    mousePosition.value = {
        x: event.clientX,
        y: event.clientY
    }

    showContextMenu.value = true
    
}

function closeContextMenu() {
    if(showContextMenu.value) showContextMenu.value = false
}

function handleDelete()
{
    deleteTask(props.id)
    closeContextMenu()
}


function handleChangeTitle() {
    changeTitle(props.id, titleEdit.value)
    isTitleEdited.value = false
    isDraggable.value = true
}

function handleChangeDescription() {
    changeDescription(props.id, descEdit.value)
    isDescEdited.value = false
    isDraggable.value = true
}

function cancelEditTitle() {
    titleEdit.value = props.title
    isTitleEdited.value = false
    isDraggable = true
}

function cancelEditDescription() {
    descEdit.value = props.description
    isDescEdited.value = false
    isDraggable = true
}

async function editTitle() {
    if(!isTitleEdited.value) isTitleEdited.value = true
    isDraggable.value = false

    await nextTick()
    titleInputRef.value?.focus()
}

async function editDescription() {
    if(!isDescEdited.value) isDescEdited.value = true
    isDraggable.value = false

    await nextTick()
    descInputRef.value?.focus()
}

watch(() => props.title, (newVal) => {titleEdit.value = newVal})
watch(() => props.description, (newVal) => {descEdit.value = newVal})

onMounted(() => window.addEventListener('click', closeContextMenu))
onUnmounted(() => window.removeEventListener('click', closeContextMenu))
</script>

<template>
    <div class="task-card" :draggable="isDraggable" @dragstart="onDragStart" @contextmenu="handleContextMenu">

        <h3 v-if="!isTitleEdited" @dblclick="editTitle">{{ props.title }}</h3>
        <input 
            v-if="isTitleEdited" 
            v-model="titleEdit" 
            ref="titleInputRef" 
            type="text" 
            @focusout="handleChangeTitle"
            @keyup.enter="handleChangeTitle"
            @keyup.esc="cancelEditTitle"
            >

        <p v-if="!isDescEdited" @dblclick="editDescription">{{ props.description }}</p>
        <input 
            v-if="isDescEdited" 
            v-model="descEdit" 
            ref="descInputRef" 
            type="text" 
            @focusout="handleChangeDescription" 
            @keyup.enter="handleChangeDescription"
            @keyup.esc="cancelEditDescription"
            >
        
        <Teleport to="body">
            <div 
                v-if="showContextMenu"
                class="context-menu"
                @click.stop 
                :style="{top: `${mousePosition.y}px`, left: `${mousePosition.x}px`}"
            >
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
}

.task-card h3 { font-size: 1rem; margin: 0 0 0.3rem 0; color: #2d3748; }
.task-card p { font-size: 0.875rem; margin: 0; color: #718096; }

/* Floating Context Menu Styles */
.context-menu {
  position: fixed; /* Fixed ensures it positions relative to the screen, not the card */
  background: white;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -4px rgba(0, 0, 0, 0.1);
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
}

.menu-item.delete {
    color: rgb(255, 83, 83)
}

.menu-item.delete:hover {
  background: #fee2e2;
  color: rgb(95, 11, 11)
}
</style>