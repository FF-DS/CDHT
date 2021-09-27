<template>
  <div>
    <!-- <div>The data from the API is :- {{ monitor }}</div> -->

    <div class="row">
      <div class="col-7">
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
      <div class="col-5">
        <div
          v-for="(model, index) in monitorStats"
          :key="index"
          :style="'cursor:pointer'"
          class="col-12"
        >
          <div class="card">
            <div class="card-body">
              {{ model }}
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
      pieSeries: [44, 55],
      pieChartOptions: {
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

      chartOptions: {
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
          categories: [
            "a",
            "b",
            "c",
            "d",
            "e",
            "f",
            "g",
            "h",
            "e",
            "f",
            "g",
            "h",
          ],
          title: { text: "Time" },
        },
        yaxis: { title: { text: "Value" }, min: 0.1, max: 4 },
        legend: {
          position: "top",
          horizontalAlign: "right",
          floating: true,
          offsetY: -10,
          offsetX: -5,
        },
      },
      series: [
        {
          name: "rtt",
          data: [
            0.28,
            1.29,
            0.33,
            1.36,
            0.32,
            0.35,
            1.33,
            0.33,
            0.36,
            0.32,
            0.32,
            3.3,
          ],
        },
        {
          name: "srt",
          data: [
            0.12,
            0.11,
            0.14,
            0.18,
            0.0,
            1.3,
            3.3,
            3.6,
            0.32,
            0.35,
            0.13,
            1.3,
          ],
        },
        {
          name: "latency",
          data: [
            0.12,
            0.11,
            0.2,
            1.8,
            1.7,
            1.3,
            1.3,
            1.4,
            0.18,
            0.17,
            0.13,
            1.3,
          ],
        },
      ],
    };
  },
  computed: {
    ...mapGetters(["getMonitoringStats"]),
    monitorStats: function() {
      return this.getMonitoringStats;
    },
  },
  methods: {
    ...mapActions(["fetchMonitoringStats"]),
  },
  created() {
    this.fetchMonitoringStats({ node_id: "2" });
  },
};
</script>

<style></style>
