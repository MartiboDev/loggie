<template>
  <span>{{ category }}</span>
  <span>{{ timestamp }}</span>
  <span>{{ message }}</span>
</template>

<script setup lang="ts">
import { ref } from "vue";

const category = ref("system");
const timestamp = ref("2021-10-01 12:00:00");
const message = ref("System started");
</script>

<style scoped>
.v-card {
  display: flex;
  align-items: center;
  gap: 8px;
}
</style>
