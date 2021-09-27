<template>
  <div>
    <div class="row">
      <div class="col-4">
        <div class="card">
          <div class="card-body">
            <v-btn
              :loading="false"
              :disabled="false"
              color="blue-grey"
              class="ma-2 white--text"
              @click="prepareAndGenerateReport"
            >
              Generate PDF
            </v-btn>
            <v-btn
              :loading="false"
              :disabled="false"
              color="blue-grey"
              class="ma-2 white--text"
              @click="toggleExpandCollapse"
            >
              Expand Log Detail
            </v-btn>

            <v-btn
              :loading="false"
              :disabled="false"
              color="blue-grey"
              class="ma-2 white--text"
              @click="getFilteredRports"
            >
              Perform Filter
            </v-btn>
          </div>
        </div>

        <div class="card">
          <!-- filter all? {{ filterParameteres.selectAllNodesAndApplyFilters }}
          <br />
          node ID {{ filterParameteres.node_id }} -->
          <div class="card-body">
            <v-checkbox
              v-model="filterParameteres.selectAllNodesAndApplyFilters"
              :label="`Select all`"
            ></v-checkbox>
            <v-text-field
              :disabled="filterParameteres.selectAllNodesAndApplyFilters"
              v-model="filterParameteres.node_id"
              :rules="rules"
            ></v-text-field>
          </div>
        </div>

        <div class="card">
          <div class="card-body">
            <!-- Log type {{ filterParameteres.log_location }} -->
            <v-radio-group v-model="filterParameteres.log_location" row>
              <v-radio label="Self" value="LOCATION_TYPE_SELF"></v-radio>
              <v-radio
                label="Incoming"
                value="INCOMING"
              ></v-radio> </v-radio-group
            ><v-radio-group v-model="filterParameteres.log_location" row>
              <v-radio label="Leaving" value="LEAVING"></v-radio>
              <v-radio label="Other" value=""></v-radio>
            </v-radio-group>
          </div>
        </div>

        <div class="card">
          <div class="card-body">
            <!-- operation status {{ filterParameteres.operation_status }} -->
            <v-radio-group v-model="filterParameteres.operation_status" row>
              <v-radio label="SUCCESS" value="SUCCESS"></v-radio>
              <v-radio label="FAILURE" value="FAILURE"></v-radio
            ></v-radio-group>
          </div>
        </div>

        <div class="card">
          <div class="card-body">
            <v-container fluid>
              <div class="row">
                <div class="col-6">
                  <v-checkbox
                    v-model="filterParameteres.metrics_to_show"
                    label="Latency"
                    value="LATENCY"
                  ></v-checkbox>
                  <v-checkbox
                    v-model="filterParameteres.metrics_to_show"
                    label="RTT"
                    value="RTT"
                  ></v-checkbox>
                </div>
                <div class="col-6">
                  <v-checkbox
                    v-model="filterParameteres.metrics_to_show"
                    label="SRT"
                    value="SRT"
                  ></v-checkbox>
                  <v-checkbox
                    v-model="filterParameteres.metrics_to_show"
                    label="Packet Loss"
                    value="PACKET_LOSS"
                  ></v-checkbox>
                </div>
              </div>
              <!-- <p>{{ filterParameteres.metrics_to_show }}</p> -->
            </v-container>
          </div>
        </div>

        <div class="card">
          <div class="card-body">
            <!-- start date
            {{ filterParameteres.date.start_date.split(" ").join("T") }}
            <br />
            end date
            {{ filterParameteres.date.end_date }} -->
            <custom-date-picker
              type="datetime"
              v-model="filterParameteres.date.start_date"
            ></custom-date-picker>
            <br />
            <custom-date-picker
              type="datetime"
              v-model="filterParameteres.date.end_date"
            ></custom-date-picker>
          </div>
        </div>
      </div>

      <div class="col-8" style="max-height:95vh ; overflow-y:scroll">
        <div class="row py-2 my-2">
          <div class="col-12 mt-4 " style="">
            <v-slider
              v-model="requestBody.limit"
              thumb-label="always"
              @change="
                () => {
                  if (expansionToggle) {
                    this.toggleExpandCollapse();
                  }

                  this.fetchReports({
                    limit: requestBody.limit.toString(),
                  });
                }
              "
            ></v-slider>
          </div>
        </div>

        <div class="row justify-content-center">
          <div class="col-11">
            <vue-html2pdf
              :show-layout="true"
              :float-layout="false"
              :enable-download="true"
              :preview-modal="false"
              :paginate-elements-by-height="1400"
              filename="report"
              :pdf-quality="2"
              :manual-pagination="false"
              pdf-format="a4"
              pdf-orientation="portrait"
              pdf-content-width="100%"
              @progress="() => {}"
              @hasStartedGeneration="() => {}"
              @hasGenerated="() => {}"
              @hasDownloaded="
                () => {
                  toggleCheckBox();
                }
              "
              ref="html2Pdf"
            >
              <section slot="pdf-content">
                <div class="row">
                  <div
                    v-for="(model, index) in report"
                    :key="index"
                    :style="'cursor:pointer'"
                    class="col-12"
                  >
                    <div class="card my-2">
                      <div class="card-body">
                        <div class="card-title">
                          <div class="row">
                            <div class="col-12" v-show="checkBoxToggle">
                              <v-checkbox
                                class="reportCheckCheckBox"
                                input-value=""
                                value
                              ></v-checkbox>
                            </div>
                            <div class="col-6"># {{ index + 1 }}</div>
                            <div class="col-6">
                              {{ model.created_date }}
                            </div>
                          </div>
                        </div>
                        <div class="row">
                          <div class="col-6">Node ID=> {{ model.node_id }}</div>
                          <div class="col-6">
                            Report Type => {{ model.type }}
                          </div>
                          <div class="row justify-content-center">
                            <!-- <div class="col-2">
                              <v-container>
                                <v-row>
                                  <v-slider
                                    v-model="jsonViewerDepthRange"
                                    step="1"
                                    vertical
                                    label="Range"
                                  ></v-slider>
                                </v-row>
                              </v-container>
                            </div> -->
                            <div class="col-10">
                              <v-expansion-panels flat light multiple>
                                <v-expansion-panel :key="index">
                                  <v-expansion-panel-header
                                    class="expansion-pannel-header"
                                  >
                                    Detailed Information
                                  </v-expansion-panel-header>
                                  <v-expansion-panel-content>
                                    <json-viewer
                                      :value="model"
                                      :expand-depth="jsonViewerDepthRange"
                                      copyable
                                      sort
                                      theme="themetheme"
                                    ></json-viewer>
                                  </v-expansion-panel-content>
                                </v-expansion-panel>
                              </v-expansion-panels>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                    <!-- <div v-for="(value, key) in awards[index]" :key="key"> -->
                  </div>
                </div>
              </section>
            </vue-html2pdf>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "report",
  componenets: {},
  data: () => {
    return {
      dialog: false,
      requestBody: {
        limit: "10",
      },
      filterParameteres: {
        selectAllNodesAndApplyFilters: true,
        node_id: "",
        log_type: "",
        log_location: "",
        metrics_to_show: [],
        operation_status: "",
        date: {
          start_date: "",
          end_date: "",
        },
      },
      checkBoxToggle: true,
      expansionToggle: false,
      jsonViewerDepthRange: 2,
      rules: [(value) => (value || "").length <= 20 || "Max 20 characters"],
    };
  },
  computed: {
    ...mapGetters(["getAllReports"]),
    report: function() {
      return this.getAllReports;
    },
  },
  methods: {
    ...mapActions(["fetchReports", "fetchFilteredReports"]),

    reloadAll() {
      console.log("the limit is now", this.requestBody.limit);
      this.fetchReports(this.limit);
    },
    /*
            Generate Report using refs and calling the
            refs function generatePdf()
        */

    prepareAndGenerateReport() {
      this.toggleCheckBox();

      this.generateReport();
    },

    generateReport() {
      this.$refs.html2Pdf.generatePdf();
    },

    toggleCheckBox() {
      this.checkBoxToggle = !this.checkBoxToggle;
    },

    toggleExpandCollapse() {
      this.expansionToggle = !this.expansionToggle;
      let expansionPannels = document.getElementsByClassName(
        "expansion-pannel-header"
      );
      expansionPannels.forEach((pannel) => {
        pannel.click();
      });
    },

    getFilteredRports() {
      let filterRequestBody = {
        limit: this.requestBody.limit.toString(),
        node_id: this.filterParameteres.node_id,
        operation_status: this.filterParameteres.operation_status,
        log_location: this.filterParameteres.log_location,
        start_date: `${this.filterParameteres.date.start_date
          .split(" ")
          .join("T")}:00.000000000`,
        end_date: `${this.filterParameteres.date.end_date
          .split(" ")
          .join("T")}:00.000000000`,
      };

      console.log(filterRequestBody);
      this.fetchFilteredReports(filterRequestBody);
    },
  },
  created() {
    this.fetchReports(this.requestBody);
  },
};
</script>

<style lang="scss">
.themetheme {
  background: #fff;
  white-space: nowrap;
  color: #525252;
  font-size: 14px;
  font-family: Consolas, Menlo, Courier, monospace;

  .jv-ellipsis {
    color: #999;
    background-color: #eee;
    display: inline-block;
    line-height: 0.9;
    font-size: 0.9em;
    padding: 0px 4px 2px 4px;
    border-radius: 3px;
    vertical-align: 2px;
    cursor: pointer;
    user-select: none;
  }
  .jv-button {
    color: #49b3ff;
  }
  .jv-key {
    color: #111111;
  }
  .jv-item {
    &.jv-array {
      color: #111111;
    }
    &.jv-boolean {
      color: #fc1e70;
    }
    &.jv-function {
      color: #067bca;
    }
    &.jv-number {
      color: #fc1e70;
    }
    &.jv-number-float {
      color: #fc1e70;
    }
    &.jv-number-integer {
      color: #fc1e70;
    }
    &.jv-object {
      color: #111111;
    }
    &.jv-undefined {
      color: #e08331;
    }
    &.jv-string {
      color: #42b983;
      word-break: break-word;
      white-space: normal;
    }
  }
  .jv-code {
    .jv-toggle {
      &:before {
        padding: 0px 2px;
        border-radius: 2px;
      }
      &:hover {
        &:before {
          background: #eee;
        }
      }
    }
  }
}
</style>
