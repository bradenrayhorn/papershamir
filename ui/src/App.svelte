<script lang="ts">
  import Router from 'svelte-spa-router';
  import Home from './lib/pages/Home.svelte';
  import Split from './lib/pages/Split.svelte';
  import ReconstructQR from './lib/pages/ReconstructQR.svelte';
  import ReconstructText from './lib/pages/ReconstructText.svelte';

  const routes = {
    '/': Home,
    '/split': Split,
    '/reconstruct-qr': ReconstructQR,
    '/reconstruct-text': ReconstructText,
  };
</script>

<Router {routes} />
