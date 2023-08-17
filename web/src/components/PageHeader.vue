<script setup>
const props = defineProps(["subheader"]);
</script>

<template>
  <div class="header">
    <h1>papershamir</h1>
    <h3 v-if="props.subheader?.length > 0">{{ props.subheader }}</h3>
    <div class="options">
      <slot />
    </div>
  </div>
</template>

<style scoped>
.header {
  margin: 1rem;
}
.options {
  display: flex;
  gap: 1rem;
  margin-top: 0.5rem;
}
</style>
