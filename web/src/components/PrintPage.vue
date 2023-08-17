<script setup>
import QRCode from "qrcode";
import { reactive, ref } from "vue";

const props = defineProps(["text", "qr", "passphrase"]);

const showingKey = ref(false);
const codes = ref([]);
const qrSize = ref("48");

props.text.forEach((text, i) => {
  const string = props.qr[i];
  QRCode.toDataURL(string, { errorCorrectionLevel: "Q" }).then((url) => {
    qrSize.value = Math.round(url.length / 250);
    if (qrSize.value < 48) {
      qrSize.value = 48;
    }
    codes.value[i] = {
      image: url,
      code: text,
    };
  });
});

function showKey() {
  showingKey.value = true;
}
</script>

<template>
  <div>
    <div v-if="!showingKey">
      <div class="menu hide-in-print">
        <button @click="$emit('backClicked')">Go Back</button>

        <button @click="showKey()">SHOW ENCRYPTION KEY</button>

        <p class="wrap">
          <b>Important!</b>
          Print this page, then write the encryption key on each printed share.
        </p>
      </div>

      <div
        class="code-list"
        :style="{
          gridTemplateRows: `repeat(${codes.length}, 1fr) 0 repeat(${codes.length}, 1fr)`,
        }"
      >
        <div v-for="code in codes" class="share">
          {{ code.code }}
        </div>

        <div class="pagebreak"></div>

        <div v-for="code in codes" class="code">
          <img :style="{ height: `${qrSize}px` }" :src="code.image" />
        </div>
      </div>
    </div>
    <div v-if="showingKey" class="passphrase hide-in-print">
      <button @click="showingKey = false">Show Printout</button>

      <p>
        <b>Passphrase:</b> <span class="mono">{{ props.passphrase }}</span>
      </p>

      <br />

      <p>Important! Write the passphrase down on each printed share.</p>
    </div>
  </div>
</template>

<style scoped>
.share {
  white-space: pre;
  font-family: monospace;
  font-size: 0.5rem;

  width: fit-content;
  padding: 0.5rem 0;
}

.pagebreak {
  break-before: page;
  page-break-before: always;
  height: 0;
}

.code-list {
  display: grid;
  grid-auto-rows: 1fr;
}

.code-list > div {
  display: flex;
  align-items: center;
}

.code {
  text-align: right;
  justify-content: flex-end;
}

.menu {
  margin: 1rem 0 1rem 0;
}

.wrap {
  margin: 1rem;
}

.passphrase {
  margin: 1rem;
}
.passphrase > button {
  margin: 0 0 1rem 0;
}

button {
  margin-left: 1rem;
}

.mono {
  white-space: pre;
  font-family: monospace;
}
</style>
