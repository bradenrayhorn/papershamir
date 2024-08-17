<script lang="ts">
  import QRCode from 'qrcode';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let passphrase: string;
  export let qrShares: string[];
  export let textShares: string[];

  let isShowingKey = false;
  let qrImages: string[] = [];
  let qrSize = 100;

  qrShares.forEach((share, i) => {
    QRCode.toDataURL(share, { errorCorrectionLevel: 'M' }).then((url: string) => {
      qrSize = Math.max(Math.round(url.length / 100), 150);
      qrImages[i] = url;
    });
  });
</script>

<div>
  {#if isShowingKey}
    <button on:click={() => (isShowingKey = false)}>Show Printout</button>

    <p>
      <b>Passphrase:</b> <span class="mono">{passphrase}</span>
    </p>

    <br />

    <p>Important! Write the passphrase down on each printed share.</p>
  {:else}
    <div class="menu hide-in-print">
      <button
        on:click={() => {
          dispatch('close');
        }}>Go Back</button
      >

      <button on:click={() => (isShowingKey = true)}>SHOW ENCRYPTION KEY</button>

      <p class="wrap">
        <b>Important!</b>
        Print this page, then write the encryption key on each printed share.
      </p>
    </div>

    <div
      class="code-list"
      style="grid-template-rows: repeat({textShares.length}, 1fr) 0 repeat({textShares.length}, 1fr)"
    >
      {#each textShares as share}
        <div class="share">
          {share}
        </div>
      {/each}

      <div class="pagebreak"></div>

      {#each qrImages as image}
        <div class="code">
          <img style="height: {qrSize}px" src={image} alt="qr code" />
        </div>
      {/each}
    </div>
  {/if}
</div>

<style>
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
    padding: 0.5rem 0;
  }

  .menu {
    margin: 1rem 0 1rem 0;
  }

  .wrap {
    margin: 1rem;
  }

  button {
    margin-left: 1rem;
  }

  .mono {
    white-space: pre;
    font-family: monospace;
  }

  @media print {
    .hide-in-print {
      display: none;
    }
  }
</style>
