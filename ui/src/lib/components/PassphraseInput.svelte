<script lang="ts">
  export let value = '';

  let error = '';
</script>

<input
  autocomplete="off"
  autocorrect="off"
  autocapitalize="none"
  spellcheck="false"
  {value}
  on:input={(e) => {
    value = e.currentTarget.value;

    let newError = '';
    for (const c of value) {
      if (!/^[ awxhekn123456789]*$/i.test(c)) {
        newError += c.toUpperCase();
      }
    }
    error = newError;
  }}
/>

{#if error}
  Invalid character in encryption key:
  <span class="mono">
    {error}
  </span>
{/if}

<style>
  .mono {
    font-family: monospace;
  }
</style>
