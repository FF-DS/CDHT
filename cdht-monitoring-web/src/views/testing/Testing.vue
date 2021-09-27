<template>
  <v-container fluid class="my-5">
    <v-layout row>
      <v-flex xs12 sm6>
        <v-layout row class="mx-2">
          <v-divider></v-divider>
        </v-layout>
      </v-flex>
      <v-flex xs12 sm6>
        <v-layout row class="mx-2">
          <v-divider></v-divider>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
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
