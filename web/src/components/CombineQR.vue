<script setup>
import { ref } from "vue";
import { QrcodeStream } from "vue-qrcode-reader";

const error = ref("");
const secret = ref("");
const key = ref("");

const scannedCodes = ref([]);

function build() {
  if (scannedCodes.value.length > 1) {
    secret.value = shamirCombineQR(scannedCodes.value, key.value);
  }
}

const onError = (err) => {
  error.value = `${err.name}: ${err.message}`;
};

const onDetect = (codes) => {
  const string = codes[0].rawValue;
  if (scannedCodes.value.includes(string)) {
    error.value = "Code already scanned!";
    return;
  }
  scannedCodes.value.push(string);

  build();

  error.value = "";
};
</script>

<template>
  <div class="wrap">
    <qrcode-stream @detect="onDetect" @error="onError" />

    <div>
      <p>Parts Scanned: {{ scannedCodes.length }}</p>
    </div>

    <div>
      <p>Encryption key:</p>
      <SecretInput v-model="key" @input="build()" />
    </div>

    <div>
      <p>Error:</p>
      <p class="red">{{ error }}</p>
    </div>

    <div>Secret:</div>
    <div class="content">{{ secret }}</div>
  </div>
</template>

<style>
.wrap {
  margin: 1rem;
}

.content {
  white-space: pre-wrap;
  font-family: monospace;
  border: 1px solid black;
  padding: 1rem;
}

video {
  max-width: 100%;
}
canvas {
  display: none;
}
</style>
