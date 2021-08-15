<template>
  <div>
    <h1>Dise ist eine configuration</h1>
    <div>The data from the API is :- {{ config }}</div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "configuration",
  componenets: {},
  data: () => {
    return {};
  },
  computed: {
    ...mapGetters(["getConfig"]),

    config: function() {
      return this.getConfig;
    },
  },
  methods: {
    ...mapActions(["fetchConfig"]),
  },
  created() {
    this.fetchConfig();
  },
};
</script>

<style></style>
