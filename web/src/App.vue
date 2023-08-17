<script setup>
import { ref } from "vue";
import Home from "./components/Home.vue";
import Split from "./components/Split.vue";
import Combine from "./components/Combine.vue";

const Showing = Object.freeze({ Home: 1, Combine: 2, Split: 3 });
const showing = ref(Showing.Home);
</script>

<template>
  <Home
    v-if="showing === Showing.Home"
    @split-clicked="showing = Showing.Split"
    @reconstruct-clicked="showing = Showing.Combine"
  />

  <Split v-if="showing === Showing.Split" @go-back="showing = Showing.Home" />
  <Combine
    v-if="showing === Showing.Combine"
    @go-back="showing = Showing.Home"
  />
</template>
