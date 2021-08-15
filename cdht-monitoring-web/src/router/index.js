import Vue from "vue";
import VueRouter from "vue-router";

import App from "@/App.vue";
import Configuration from "@/views/configuration/Configuration.vue";
import Monitoring from "@/views/monitoring/Monitoring.vue";
import Report from "@/views/report/Report.vue";
import Testing from "@/views/testing/Testing.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/home",
    name: "Home",
    component: App,
  },
  {
    path: "/configuration",
    name: "Configuration",
    component: Configuration,
  },
  {
    path: "/monitoring",
    name: "Monitoring",
    component: Monitoring,
  },
  {
    path: "/report",
    name: "Report",
    component: Report,
  },
  {
    path: "/testing",
    name: "Testing",
    component: Testing,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
