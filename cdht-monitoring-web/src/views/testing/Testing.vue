<template>
  <v-container fluid class="my-5">
    <v-layout row>
      <v-flex xs12 sm6>
        <v-layout row class="mx-2">
          <v-card elevation="2" outlined shaped tile>
            <v-divider></v-divider>
            <div class="text-darken-1 mb-2 h3">
              Run Test tools on a node
            </div>
            <v-divider></v-divider>
            <v-card elevation="3" tile>
              <v-card-title>Ping Tool</v-card-title>
              <v-text-field
                label="Node ID"
                v-model="pingNodeId"
                :rules="pingNodeIdRules"
                placeholder="2"
                class="mx-2"
                outlined
              ></v-text-field>

              <v-card-actions
                ><v-btn
                  color="success"
                  elevation="2"
                  class="mr-4"
                  @click="runPing"
                >
                  Run Ping
                  <v-icon right dark>
                    mdi-subdirectory-arrow-right
                  </v-icon>
                </v-btn></v-card-actions
              >
            </v-card>
            <v-card class="my-2" elevation="3" tile>
              <v-card-title>Card title</v-card-title>
              <v-card-subtitle>Card sub title</v-card-subtitle>
              <v-card-actions>Action 1</v-card-actions>
            </v-card>
            <v-card class="my-2" elevation="3" tile>
              <v-card-title>Card title</v-card-title>
              <v-card-subtitle>Card sub title</v-card-subtitle>
              <v-card-actions>Action 1</v-card-actions>
            </v-card>
          </v-card>
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
    return {
      pingNodeId: null,
      pingNodeIdRules: [
        (v) => !isNaN(parseFloat(v)) || "It has to be a Number",
      ],
    };
  },
  computed: {
    ...mapGetters(["getTest"]),
    test: function() {
      return this.getTest;
    },
  },
  methods: {
    ...mapActions(["fetchTest"]),
    runPing() {},
  },
  created() {
    this.fetchTest;
  },
};
</script>

<style></style>
