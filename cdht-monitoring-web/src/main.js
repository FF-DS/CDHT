import Vue from "vue";
import App from "./App.vue";
import store from "./store";
import router from "@/router";
import vuetify from "./plugins/vuetify";

import JSONView from "vue-json-component";
Vue.use(JSONView);

import TreeView from "vue-json-tree-view";
Vue.use(TreeView);

import JsonViewer from "vue-json-viewer";

// Import JsonViewer as a Vue.js plugin
Vue.use(JsonViewer);

import VueHtml2pdf from "vue-html2pdf";
Vue.use(VueHtml2pdf);

import VueDatetimePickerJs from "vue-date-time-picker-js";
Vue.use(VueDatetimePickerJs, {
  name: "custom-date-picker",
  props: {
    inputFormat: "YYYY-MM-DD HH:mm",
    format: "YYYY-MM-DD HH:mm",
    editable: false,
    inputClass: "form-control my-custom-class-name",
    placeholder: "Please select a date",
    altFormat: "YYYY-MM-DD HH:mm",
    color: "#00acc1",
    autoSubmit: false,
    //...
    //... And whatever you want to set as default
    //...
  },
});

Vue.config.productionTip = false;

new Vue({
  vuetify,
  store,
  router,
  render: (h) => h(App),
}).$mount("#app");
