<script setup>
import { ref } from "vue";
import CombineQR from "./CombineQR.vue";
import CombineText from "./CombineText.vue";

const Showing = Object.freeze({ Options: "Options", QR: "QR", Text: "Text" });
const showing = ref(Showing.Options);
</script>

<template>
  <div>
    <PageHeader :subheader="`Reconstruct Secret - ${showing}`">
      <button @click="$emit('goBack')">Go Back</button>
    </PageHeader>

    <div v-if="showing === Showing.Options" class="selector-buttons">
      <button @click="showing = Showing.QR">qr</button>
      <button @click="showing = Showing.Text">text</button>
    </div>

    <CombineQR v-if="showing === Showing.QR" />
    <CombineText v-if="showing === Showing.Text" />
  </div>
</template>

<style scoped>
.selector-buttons {
  display: flex;
  gap: 1rem;
  margin-left: 1rem;
}
</style>
