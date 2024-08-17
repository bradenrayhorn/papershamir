<script lang="ts">
  import PrintPreview from '../components/PrintPreview.svelte';

  let threshold = '3';
  let shares = '5';
  let secret = '';
  let error = '';

  let passphrase = '';
  let qrShares: string[] = [];
  let textShares: string[] = [];

  $: isShowingPreview = !!passphrase && qrShares.length > 0 && textShares.length > 0;

  function splitSecret() {
    const result = shamirSplit(secret.trim(), +shares, +threshold);

    if (result instanceof Error) {
      error = result?.message ?? 'Unknown error';
      return;
    }

    error = '';
    textShares = result.text;
    qrShares = result.qr;
    passphrase = result.passphrase;
  }
</script>

{#if isShowingPreview}
  <PrintPreview
    {qrShares}
    {textShares}
    {passphrase}
    on:close={() => {
      passphrase = '';
      qrShares = [];
      textShares = [];
    }}
  />
{:else}
  <h2>Split new secret</h2>
  <a href="#/">Cancel</a>

  <div class="action">
    <label>
      <b>Threshold</b>
      <input type="number" bind:value={threshold} />
    </label>

    <label>
      <b>Shares</b>
      <input type="number" bind:value={shares} />
    </label>

    <label>
      <b>Secret:</b>
      <textarea bind:value={secret}></textarea>
    </label>

    <div class="action-buttons">
      <button on:click={splitSecret}>Split secret</button>
    </div>

    {#if error}
      <p class="red">
        <b>{error}</b>
      </p>
    {/if}
  </div>
{/if}

<style>
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
