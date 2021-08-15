<template>
  <div>
    <h1>Dise ist eine monitoring</h1>
    <div>The data from the API is :- {{ monitor }}</div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "monitoring",
  componenets: {},
  data: () => {
    return {};
  },
  computed: {
    ...mapGetters(["getMonitor"]),
    monitor: function() {
      return this.getMonitor;
    },
  },
  methods: {
    ...mapActions(["fetchMonitor"]),
  },
  created() {
    this.fetchMonitor();
  },
};
</script>

<style></style>
