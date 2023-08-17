import "./assets/main.css";

import { createApp } from "vue";
import App from "./App.vue";
import PageHeader from "./components/PageHeader.vue";
import SecretInput from "./components/SecretInput.vue";

const app = createApp(App);

app.component("PageHeader", PageHeader);
app.component("SecretInput", SecretInput);

app.mount("#app");
