<script setup>
import { ref, defineProps, onMounted } from "vue";

const props = defineProps(["modelValue"]);

const r = ref();
onMounted(() => {
  r.value.setAttribute("autocomplete", "off");
  r.value.setAttribute("autocorrect", "off");
  r.value.setAttribute("autocapitalize", "none");
  r.value.setAttribute("spellcheck", "false");
});
</script>

<template>
  <input
    ref="r"
    :value="modelValue"
    @input="$emit('update:modelValue', $event.target.value)"
  />
</template>

<style scoped>
input {
  text-transform: uppercase;
  width: 100%;
}
</style>
