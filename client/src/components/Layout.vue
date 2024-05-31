<template>
  <div class="layout-container">
    <div class="header-container">
      <h4>Loggie Log Viewer</h4>
    </div>
    <div
      v-if="showSidebar"
      class="sidebar-container"
      :class="{ hidden: !showSidebar }"
    >
      <span class="link-item">Logs</span>
      <v-btn block @click="hideSidebar()">Hide sidebar</v-btn>
    </div>
    <div class="main-container" :class="{ 'full-width': !showSidebar }">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

const showSidebar = ref(true);

const hideSidebar = () => {
  showSidebar.value = false;
};
</script>

<style scoped lang="scss">
.layout-container {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  grid-template-rows: auto 1fr;

  height: 100vh;
}

.header-container {
  grid-column: 1 / 6;
  grid-row: 1;
  background-color: #333;
  padding: 8px;
}

.sidebar-container {
  grid-column: 1;
  grid-row: 2;
  background-color: #444;
  padding: 8px;

  &.hidden {
    display: none;
  }
}

.main-container {
  grid-column: 2 / 6;
  grid-row: 2;
  background-color: #555;
  padding: 8px;

  &.full-width {
    grid-column: 1 / 6;
  }
}
</style>
