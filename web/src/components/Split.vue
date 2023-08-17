<script setup>
import { computed, ref } from "vue";
import PrintPage from "./PrintPage.vue";

const thresholdInput = ref("3");
const sharesInput = ref("5");
const secretInput = ref("");
const error = ref("");

const printText = ref("");
const printQR = ref("");
const printPassphrase = ref("");

function splitSecret() {
  const result = shamirSplit(
    secretInput.value.trim(),
    +sharesInput.value,
    +thresholdInput.value,
  );

  if (result instanceof Error) {
    error.value = result;
    return;
  }

  error.value = "";
  printText.value = result.text;
  printQR.value = result.qr;
  printPassphrase.value = result.passphrase;
}

function clear() {
  printText.value = "";
  printQR.value = "";
  printPassphrase.value = "";
  error.value = "";
}

const showingPrint = computed(
  () => !!printText.value && !!printQR.value && !!printPassphrase.value,
);
</script>

<template>
  <div>
    <div v-if="!showingPrint">
      <PageHeader subheader="Split New Secret">
        <button @click="$emit('goBack')">Go Back</button>
      </PageHeader>

      <div class="action">
        <label>
          <b>Threshold</b>
          <input type="number" v-model="thresholdInput" />
        </label>

        <label>
          <b>Shares</b>
          <input type="number" v-model="sharesInput" />
        </label>

        <label>
          <b>Secret:</b>
          <textarea v-model="secretInput"></textarea>
        </label>

        <div class="action-buttons">
          <button @click="splitSecret">Split secret</button>
        </div>

        <p v-if="!!error" class="red">
          <b>{{ error }}</b>
        </p>
      </div>
    </div>

    <PrintPage
      v-if="showingPrint"
      @back-clicked="clear()"
      :text="printText"
      :qr="printQR"
      :passphrase="printPassphrase"
    />
  </div>
</template>

<style scoped>
.action {
  display: flex;
  flex-direction: column;
  margin: 1rem;
  gap: 1rem;
}
.action-buttons {
  display: flex;
  margin-top: 2rem;
  gap: 1rem;
}

textarea,
label {
  display: block;
}

textarea {
  min-width: 100%;
  min-height: 25vh;
}
</style>
