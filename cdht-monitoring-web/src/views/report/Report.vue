<template>
  <div>
    <h1>Dise ist eine report</h1>
    <div>The result from the API is :- {{ report }}</div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "report",
  componenets: {},
  data: () => {
    return {};
  },
  computed: {
    ...mapGetters(["getReport"]),
    report: function() {
      return this.getReport;
    },
  },
  methods: {
    ...mapActions(["fetchReports"]),
  },
  created() {
    this.fetchReports();
  },
};
</script>

<style></style>
