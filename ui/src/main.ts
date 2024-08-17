import './main.css';
import App from './App.svelte';
import './wasm/wasm_exec.js';
import mainWASM from './wasm/main.wasm?url';

const go = new Go();
WebAssembly.instantiateStreaming(fetch(mainWASM), go.importObject).then((result) => {
  go.run(result.instance);
});

const app = new App({
  target: document.getElementById('app')!,
});

export default app;
