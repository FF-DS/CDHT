<template>
  <div>
    <!-- <div>The data from the API is :- {{ monitor }}</div> -->

    <div class="row mt-3">
      <div class="col-7" style="max-height:95vh ; overflow-y:scroll">
        <div class="row justify-content-center">
          <div class="col-12">
            <div id="chart">
              <apexchart
                width="100%"
                type="area"
                :options="chartOptions"
                :series="series"
              ></apexchart>
            </div>
          </div>
          <div class="col-12 align-self-center">
            <div id="pie-chart" class="">
              <apexchart
                type="donut"
                width="55%"
                :options="pieChartOptions"
                :series="pieSeries"
              ></apexchart>
            </div>
          </div>
        </div>
      </div>
      <div class="col-5" style="max-height:95vh ; overflow-y:scroll">
        <div class="col-12">
          <p>Node ID</p>
          <v-text-field
            v-model="filterParameters.node_id"
            :rules="rules"
          ></v-text-field>
          <p>
            <v-btn
              :loading="false"
              :disabled="false"
              color="blue-grey"
              class="ma-2 white--text"
              @click="refresh"
            >
              Refresh Stats
            </v-btn>
          </p>
        </div>
        <div v-if="monitorStats != null">
          <div
            v-for="(model, index) in monitorStats"
            :key="index"
            :style="'cursor:pointer'"
            class="col-12"
          >
            <div class="card">
              <div class="card-body">
                <div class="row">
                  <div class="col-6 my-2">
                    <p>Node Address : {{ model.node_address }}</p>
                    <p>Log Location : {{ model.log_location }}</p>
                    <p>Operation Status : {{ model.operation_status }}</p>
                  </div>
                  <div class="col-6">
                    <div class="card">
                      <div class="card-body">
                        <div class="card-title">
                          Stats
                        </div>
                        <json-viewer
                          :value="model.log_body"
                          :expand-depth="0"
                          copyable
                          sort
                          theme="themetheme"
                        ></json-viewer>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div v-else>
          <div class="card">
            <div class="card-body">
              There is no data to show
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "monitoring",
  componenets: {},
  data: () => {
    return {
      filterParameters: {
        node_id: "2",
        limit: "10",
      },
      pieSeries: [22, 78],
      pieChartOptions: {
        title: {
          text: "Failure to Success ratio",
          align: "left",
        },
        labels: ["Packet Loss", "Packet Delivered"],
        chart: {
          type: "donut",
        },
        responsive: [
          {
            breakpoint: 480,
            options: {
              chart: {
                width: 200,
              },
              legend: {
                position: "bottom",
              },
            },
          },
        ],
      },
    };
  },
  computed: {
    ...mapGetters(["getMonitoringStats"]),

    monitorStats: function() {
      return this.getMonitoringStats;
    },
    rttValues: function() {
      let rttVals = [];
      if (this.monitorStats != undefined || this.monitorStats != null) {
        this.monitorStats.forEach((report) => {
          report.log_body.forEach((logBodyEntry) => {
            if (logBodyEntry.Key == "rtt") {
              rttVals.push(Number(logBodyEntry.Value));
            }
          });
        });
      }
      return rttVals;
    },
    srtValues: function() {
      let srtVals = [];
      if (this.monitorStats != undefined || this.monitorStats != null) {
        this.monitorStats.forEach((report) => {
          report.log_body.forEach((logBodyEntry) => {
            if (logBodyEntry.Key == "srt") {
              srtVals.push(parseFloat(logBodyEntry.Value));
            }
          });
        });
      }
      return srtVals;
    },
    latencyValues: function() {
      let latencyVals = [];
      if (this.monitorStats != undefined || this.monitorStats != null) {
        this.monitorStats.forEach((report) => {
          report.log_body.forEach((logBodyEntry) => {
            if (logBodyEntry.Key == "latency") {
              latencyVals.push(parseFloat(logBodyEntry.Value));
            }
          });
        });
      }
      return latencyVals;
    },
    timeValues: function() {
      let timeVals = [];

      if (this.monitorStats != undefined || this.monitorStats != null) {
        this.monitorStats.forEach((report) => {
          timeVals.push(parseFloat(report.created_date.split("T")[1]));
        });
      }
      return timeVals;
    },

    chartOptions: function() {
      return {
        chart: {
          height: 350,
          type: "line",
          dropShadow: {
            enabled: true,
            color: "#000",
            top: 18,
            left: 7,
            blur: 10,
            opacity: 0.2,
          },
          toolbar: { show: true },
        },
        colors: ["#77B6EA", "#545454", "#5454cc"],
        dataLabels: { enabled: true },
        stroke: { curve: "smooth" },
        title: {
          text: "RTT, SRT and Latency distribution over time",
          align: "left",
        },
        grid: {
          borderColor: "#e7e7e7",
          row: {
            colors: ["#f3f3f3", "transparent"], // takes an array which will be repeated on columns opacity: 0.5
          },
        },
        markers: { size: 1 },
        xaxis: {
          categories: ["a", "b", "c", "d", "e", "g"],
          title: { text: "Time" },
        },
        yaxis: { title: { text: "Value" }, min: 0.1, max: 2 },
        legend: {
          position: "top",
          horizontalAlign: "right",
          floating: true,
          offsetY: -10,
          offsetX: -5,
        },
      };
    },

    series: function() {
      return [
        {
          name: "rtt",
          data: [0.1, 0.2, 1.3, 1.6, 0.9, 1.75],
        },
        {
          name: "srt",
          data: [1.1, 0.25, 0.3, 1.6, 0.25, 0.75],
        },
        {
          name: "latency",
          data: [0.1, 0.2, 0.3, 0.6, 1.9, 0.35],
        },
      ];
    },
  },
  methods: {
    ...mapActions(["fetchMonitoringStats"]),
    refresh() {
      this.fetchMonitoringStats({
        limit: this.filterParameters.limit,
        node_id: this.filterParameters.node_id,
      });
    },
  },
  created() {
    this.fetchMonitoringStats({
      limit: this.filterParameters.limit,
      node_id: this.filterParameters.node_id,
    });
  },
};
</script>

<style></style>
