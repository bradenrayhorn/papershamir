<template>
  <PageHeader>
    <button @click="$emit('splitClicked')">Split new secret</button>
    <button @click="$emit('reconstructClicked')">
      Reconstruct existing secret
    </button>
  </PageHeader>
</template>
