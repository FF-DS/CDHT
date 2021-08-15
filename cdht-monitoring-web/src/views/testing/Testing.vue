<template>
  <div>
    <h1>Dise ist eine testing</h1>
    <div>The result from the API is :- {{ test }}</div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "testing",
  componenets: {},
  data: () => {
    return {};
  },
  computed: {
    ...mapGetters(["getTest"]),
    test: function() {
      return this.getTest;
    },
  },
  methods: {
    ...mapActions(["fetchTest"]),
  },
  created() {
    this.fetchTest;
  },
};
</script>

<style></style>
