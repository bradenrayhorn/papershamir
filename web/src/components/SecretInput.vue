<script setup>
import { ref, onMounted } from "vue";
defineEmits(["input", "update:modelValue"]);

const props = defineProps(["modelValue"]);

const error = ref("");

const r = ref();
onMounted(() => {
  r.value.setAttribute("autocomplete", "off");
  r.value.setAttribute("autocorrect", "off");
  r.value.setAttribute("autocapitalize", "none");
  r.value.setAttribute("spellcheck", "false");
});

function onInputChange(event) {
  const input = event.target.value;

  error.value = "";

  for (const c of input) {
    if (!/^[ awxhekn123456789]*$/i.test(c)) {
      error.value += c.toUpperCase();
    }
  }
}
</script>

<template>
  <input
    ref="r"
    :value="modelValue"
    @input="
      $emit('update:modelValue', $event.target.value);
      $emit('input', $event.target.value);
      onInputChange($event);
    "
  />
  <span v-if="error" class="red">
    Invalid character in encryption key:
    <span class="mono">
      {{ error }}
    </span>
  </span>
</template>

<style scoped>
input {
  text-transform: uppercase;
  width: 100%;
}

.mono {
  font-family: monospace;
}
</style>
