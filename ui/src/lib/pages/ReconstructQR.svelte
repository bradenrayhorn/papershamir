<script lang="ts">
  import { Html5QrcodeScanner, Html5QrcodeSupportedFormats } from 'html5-qrcode';
  import { onMount } from 'svelte';
  import PassphraseInput from '../components/PassphraseInput.svelte';

  let scannedCodes: string[] = [];
  let error = '';
  let secret = '';
  let passphrase = '';

  function constructSecret() {
    if (scannedCodes.length > 1) {
      secret = shamirCombineQR(scannedCodes, passphrase);
    }
  }

  onMount(() => {
    const scanner = new Html5QrcodeScanner(
      'qr-reader',
      { fps: 10, formatsToSupport: [Html5QrcodeSupportedFormats.QR_CODE] },
      false,
    );

    scanner.render(
      (value) => {
        if (scannedCodes.includes(value)) {
          return;
        }
        scannedCodes = [...scannedCodes, value];
        error = '';

        constructSecret();
      },
      () => {},
    );
  });
</script>

<h2>Reconstruct via QR Code</h2>

<div>
  <p>
    First, enter encryption key:
    <PassphraseInput bind:value={passphrase} />
  </p>
  <p><button on:click={constructSecret}>Build secret</button></p>

  <p>
    <b>Codes scanned: {scannedCodes.length}</b>
  </p>

  <div>
    <pre>{secret}</pre>
  </div>

  <p>
    {error}
  </p>
</div>

<div id="qr-reader"></div>
