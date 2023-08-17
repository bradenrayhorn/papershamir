<script setup>
import { ref } from "vue";

const error = ref("");
const inputSecret = ref("");
const inputKey = ref("");
const secret = ref("");

function reconstruct() {
  const result = shamirCombine(inputSecret.value, inputKey.value);
  secret.value = result;
}
</script>

<template>
  <div class="wrap">
    <label>
      <b>Input secret</b>
      <p>Please enter a blank newline between each secret.</p>
      <textarea v-model="inputSecret"> </textarea>
    </label>

    <div>
      <b>Encryption key:</b>
      <SecretInput v-model="inputKey" />
    </div>

    <p>
      <button @click="reconstruct()">Reconstruct</button>
    </p>

    <div>
      <p>Error:</p>
      <p class="red">{{ error }}</p>
    </div>

    <div>Secret:</div>
    <div class="content">{{ secret }}</div>
  </div>
</template>

<style scoped>
.wrap {
  margin: 1rem;
}

.content {
  white-space: pre-wrap;
  font-family: monospace;
  border: 1px solid black;
  padding: 1rem;
}

textarea {
  min-width: 100%;
  min-height: 25vh;
}
</style>
