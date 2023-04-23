import { createApp } from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import store from "./store";

// Vuetify
import "vuetify/styles";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";

// Styles
import "@/styles/index.scss";

const vuetify = createVuetify({
  components,
  directives,
  defaults: {
    global: {
      ripple: true,
      elevation: 0,
      density: "compact",
    },
    a: {
      textDecoration: "none",
    },
  },
});

const app = createApp(App);
app.use(store);
app.use(router);
app.use(vuetify);
app.mount("#app");
