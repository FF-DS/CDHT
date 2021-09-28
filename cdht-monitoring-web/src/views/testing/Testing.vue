<template>
  <v-container fluid class="my-5">
    <v-layout row>
      <v-flex xs12 sm5>
        <v-layout row class="mx-2">
          <v-card elevation="2" outlined shaped tile>
            <v-divider></v-divider>
            <div class="text-darken-1 mb-2 h3">
              Run Test tools on a node
            </div>
            <v-divider></v-divider>
            <v-card elevation="3" tile class="my-3">
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
            <v-card elevation="3" tile class="my-3">
              <v-card-title>DNS lookup</v-card-title>
              <v-text-field
                label="Node ID"
                v-model="lookupNodeId"
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
                  @click="runDNSLookup"
                >
                  Run DNS lookup
                  <v-icon right dark>
                    mdi-subdirectory-arrow-right
                  </v-icon>
                </v-btn></v-card-actions
              >
            </v-card>
            <v-card elevation="3" tile class="my-3">
              <v-card-title>Hop Count Tool</v-card-title>
              <v-text-field
                label="Start Node ID"
                v-model="hopNodeId1"
                :rules="pingNodeIdRules"
                placeholder="2"
                class="mx-2 my-0"
                outlined
              ></v-text-field>
              <v-text-field
                label="Destination Node ID"
                v-model="hopNodeId2"
                :rules="pingNodeIdRules"
                placeholder="2"
                class="mx-2 my-0"
                outlined
              ></v-text-field>

              <v-card-actions
                ><v-btn
                  color="success"
                  elevation="2"
                  class="mr-4"
                  @click="runHopCount"
                >
                  Run Hop Count
                  <v-icon right dark>
                    mdi-subdirectory-arrow-right
                  </v-icon>
                </v-btn></v-card-actions
              >
            </v-card>
          </v-card>
        </v-layout>
      </v-flex>
      <v-flex xs12 sm7>
        <v-btn
          :loading="false"
          :disabled="false"
          color="blue-grey"
          class="ma-2 white--text"
          @click="fetchResultsForActiveOperation"
        >
          Refresh Results
        </v-btn>

        {{ activeOperationResults }}
        {{ activeOperationId }}
        <v-card class="mx-0" elevation="2">
          <div class="">
            <v-layout row class="mx-2" fluid>
              <v-container class="ma-0 pa-0">
                <v-divider></v-divider>
                <div class="text-darken-1 mb-2 h4">
                  Response log from the node
                </div>
                <v-divider></v-divider>
                <div
                  class=""
                  v-for="(report, index) in logReports"
                  v-bind:key="index"
                >
                  <v-card
                    elevation="2"
                    v-if="report.Type == 'COMMAND_TYPE_PING'"
                  >
                    <v-layout class="my-1 pa-3">
                      <v-flex xs6>
                        <span>Command executor Node Info</span>

                        <v-card-text>
                          <v-list-item two-line>
                            <v-list-item-content>
                              <v-list-item-subtitle>
                                <span class="text--gray">Log Type :</span>
                                {{ report.Type }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node ID :</span>
                                {{ report.NodeId }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node Address :</span>
                                {{ report.NodeAddress }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray"
                                  >Operation Status :</span
                                >
                                {{ report.OperationStatus }}
                              </v-list-item-subtitle>
                            </v-list-item-content>
                          </v-list-item>
                        </v-card-text>
                      </v-flex>
                      <v-divider vertical></v-divider>
                      <v-flex xs5>
                        <span>Remote Node Info</span>

                        <v-card-text>
                          <v-list-item two-line>
                            <v-list-item-content>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node Status :</span>
                                {{ report.Body.rtt }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node id :</span>
                                {{ report.Body.Node_id }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node Address :</span>
                                {{ report.Body.Node_Address }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">RTT :</span>
                                {{ report.Body.rtt }}
                              </v-list-item-subtitle>
                            </v-list-item-content>
                          </v-list-item>
                        </v-card-text>
                      </v-flex>
                    </v-layout>
                  </v-card>

                  <v-card
                    elevation="2"
                    v-if="report.Type == 'COMMAND_TYPE_HOP_COUNT'"
                  >
                    <v-layout class="my-1 pa-3">
                      <v-flex xs6>
                        <span>Command executor Node Info</span>

                        <v-card-text>
                          <v-list-item two-line>
                            <v-list-item-content>
                              <v-list-item-subtitle>
                                <span class="text--gray">Log Type :</span>
                                {{ report.Type }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node ID :</span>
                                {{ report.NodeId }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node Address :</span>
                                {{ report.NodeAddress }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray"
                                  >Operation Status :</span
                                >
                                {{ report.OperationStatus }}
                              </v-list-item-subtitle>
                            </v-list-item-content>
                          </v-list-item>
                        </v-card-text>
                      </v-flex>
                      <v-divider vertical></v-divider>
                      <v-flex xs5>
                        <span>Remote Node Info</span>

                        <v-card-text>
                          <v-list-item two-line>
                            <v-list-item-content>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node Status :</span>
                                {{ report.Body.rtt }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">Hop length :</span>
                                {{ report.Body.length }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">RTT :</span>
                                {{ report.Body.rtt }}
                              </v-list-item-subtitle>
                            </v-list-item-content>
                          </v-list-item>
                        </v-card-text>
                      </v-flex>
                    </v-layout>
                  </v-card>

                  <v-card
                    elevation="2"
                    v-if="report.Type == 'COMMAND_TYPE_LOOK_UP'"
                  >
                    <v-layout class="my-1 pa-3">
                      <v-flex xs6>
                        <span>Command executor Node Info</span>

                        <v-card-text>
                          <v-list-item two-line>
                            <v-list-item-content>
                              <v-list-item-subtitle>
                                <span class="text--gray">Log Type :</span>
                                {{ report.Type }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node ID :</span>
                                {{ report.NodeId }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node Address :</span>
                                {{ report.NodeAddress }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray"
                                  >Operation Status :</span
                                >
                                {{ report.OperationStatus }}
                              </v-list-item-subtitle>
                            </v-list-item-content>
                          </v-list-item>
                        </v-card-text>
                      </v-flex>
                      <v-divider vertical></v-divider>
                      <v-flex xs5>
                        <span>Remote Node Info</span>

                        <v-card-text>
                          <v-list-item two-line>
                            <v-list-item-content>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node M :</span>
                                {{ report.Body.NodeM }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node id :</span>
                                {{ report.Body.NodeId }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">Node Address :</span>
                                {{ report.Body.NodeAddress }}
                              </v-list-item-subtitle>
                              <v-list-item-subtitle>
                                <span class="text--gray">RTT :</span>
                                {{ report.Body.rtt }}
                              </v-list-item-subtitle>
                            </v-list-item-content>
                          </v-list-item>
                        </v-card-text>
                      </v-flex>
                    </v-layout>
                  </v-card>
                </div>
              </v-container>
            </v-layout>
          </div>
        </v-card>
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
      lookupNodeId: null,
      hopNodeId1: null,
      hopNodeId2: null,
      logReports: [
        {
          Type: "COMMAND_TYPE_PING",
          NodeId: "NodeId",
          NodeAddress: "NodeAddress",
          OperationStatus: "OperationStatus",
          Body: {
            rtt: "rtt",
            Node_Status: "Node_Status",
            Node_id: "Node_id",
            Node_Address: "Node_Address",
          },
        },
        {
          Type: "COMMAND_TYPE_HOP_COUNT",
          NodeId: "NodeId",
          NodeAddress: "NodeAddress",
          OperationStatus: "OperationStatus",
          Body: {
            hops: "rtt",
            rtt: "Node_Status",
            length: "length",
          },
        },
        {
          Type: "COMMAND_TYPE_LOOK_UP",
          NodeId: "NodeId",
          NodeAddress: "NodeAddress",
          OperationStatus: "OperationStatus",
          Body: {
            NodeId: "NodeId",
            NodeAddress: "NodeAddress",
            NodeM: "NodeM",
            rtt: "rtt",
          },
        },
        {
          Type: "COMMAND_TYPE_LOOK_UP",
          NodeId: "NodeId",
          NodeAddress: "NodeAddress",
          OperationStatus: "OperationStatus",
          Body: {
            NodeId: "NodeId",
            NodeAddress: "NodeAddress",
            NodeM: "NodeM",
            rtt: "rtt",
          },
        },
        {
          Type: "COMMAND_TYPE_LOOK_UP",
          NodeId: "NodeId",
          NodeAddress: "NodeAddress",
          OperationStatus: "OperationStatus",
          Body: {
            NodeId: "NodeId",
            NodeAddress: "NodeAddress",
            NodeM: "NodeM",
            rtt: "rtt",
          },
        },
      ],
      pingNodeIdRules: [
        (v) => !isNaN(parseFloat(v)) || "It has to be a Number",
      ],
    };
  },
  computed: {
    ...mapGetters(["getActiveOperationId", "getActiveOperationResults"]),
    test: function() {
      return this.getTest;
    },
    activeOperationId: function() {
      return this.getActiveOperationId;
    },
    activeOperationResults: function() {
      return this.getActiveOperationResults;
    },
  },
  methods: {
    ...mapActions([
      "sendPingRequest",
      "sendDNSRequest",
      "sendHopCountRequest",
      "getCurrentActiveOperationReport",
    ]),
    runPing() {
      this.sendPingRequest({
        body: {
          node_id: this.pingNodeId,
        },
      });
    },
    runDNSLookup() {
      this.sendDNSRequest({
        body: {
          node_id: this.lookupNodeId,
        },
      });
    },
    runHopCount() {
      this.sendHopCountRequest({
        body: {
          node_id_1: this.hopNodeId1,
          node_id_2: this.hopNodeId2,
        },
      });
    },
    fetchResultsForActiveOperation() {
      this.getCurrentActiveOperationReport({
        operation_id: this.getActiveOperationId,
      });
    },
  },
  created() {
    this.fetchTest;
  },
};
</script>

<style></style>
