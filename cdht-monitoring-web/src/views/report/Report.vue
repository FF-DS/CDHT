<template>
  <div class="container">
    <div class="row">

      <div class="first_col">
        <div class="col">
          <div class="row">
            <div class="card">
              <div class="card-body">
                <div class="col">
                  <select
                    class="form-select form-select-sm"
                    aria-label=".form-select-sm example"
                    id="file_type"
                  >
                    <option selected>PDF</option>
                    <option value="1">DOCX</option>
                    <option value="2">XSL</option>
                  </select>
                </div>
                <div class="col">
                  <button
                    class="btn btn-primary"
                    type="submit"
                  >Generate</button>
                </div>
              </div>
            </div>
          </div>
          <div>
            <div class="row">
              <div class="card">
                <div class="card-body">
                  <div class="col">
                    <div class="mb-3">
                      <input
                        type="text"
                        class="form-control"
                        id="formGroupExampleInput"
                        placeholder="Node ID"
                      >
                    </div>
                  </div>
                  <div class="col">
                    <input
                      type="checkbox"
                      id="pick_node"
                      name="all_nodes"
                      value="pick_node"
                    >
                    <label for="packet_loss">All</label>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card-body">
              <p>Message Type</p>
              <div class="row">
                <div class="col">
                  <input
                    type="radio"
                    id="all_types"
                    name="drone"
                    value="all_types"
                    checked
                  >
                  <label for="all_types">All</label>
                </div>
                <div class="col">
                  <input
                    type="radio"
                    id="leaving"
                    name="drone"
                    value="leaving"
                  >
                  <label for="leaving">Leaving</label>
                </div>
              </div>
              <div class="row">
                <div class="col">
                  <input
                    type="radio"
                    id="corrupted"
                    name="drone"
                    value="corrupted"
                  >
                  <label for="corrupted">Corrupted</label>
                </div>
                <div class="col">
                  <input
                    type="radio"
                    id="incoming"
                    name="drone"
                    value="incoming"
                  >
                  <label for="incoming">Incoming</label>
                </div>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card-body">
              <p>
                Metric Picker
              </p>
              <div class="row">
                <div class="col">
                  <input
                    type="checkbox"
                    id="rtt"
                    name="metric"
                    value="rtt"
                    checked
                  >
                  <label for="rtt">RTT</label>
                </div>
                <div class="col">
                  <input
                    type="checkbox"
                    id="latency"
                    name="metric"
                    value="latency"
                  >
                  <label for="leaving">Latency</label>
                </div>
              </div>
              <div class="row">
                <div class="col">
                  <input
                    type="checkbox"
                    id="packet_loss"
                    name="metric"
                    value="packet_loss"
                  >
                  <label for="packet_loss">Packet loss</label>
                </div>
                <div class="col">
                  <input
                    type="checkbox"
                    id="srt"
                    name="metric"
                    value="srt"
                  >
                  <label for="srt">SRT</label>
                </div>
              </div>
            </div>
          </div>
          <div>
            <div class="card">
              <div class="card-body">
                <p>
                  Date Picker
                </p>
                <div class="row">
                  <div class="mb-3">
                    <input
                      type="text"
                      class="form-control"
                      id="formGroupExampleInput"
                      placeholder="start date"
                    >
                  </div>
                </div>
                <div class="row">
                  <div class="mb-3">
                    <input
                      type="text"
                      class="form-control"
                      id="formGroupExampleInput"
                      placeholder="end date"
                    >
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="col">
        <div class="row">
          <div class="card">
            <div class="card-body">
              <div class="row">
                <div class="col">node id</div>
                <div class="col">
                  <div class="form-check form-switch">
                    <input
                      class="form-check-input"
                      type="checkbox"
                      id="flexSwitchCheckDefault"
                    >
                    <label
                      class="form-check-label"
                      for="flexSwitchCheckDefault"
                    >Include in report</label>
                  </div>
                </div>
              </div>
              <div class="row">
                <div class="col">message type</div>
                <div class="col">RTT</div>
              </div>
              <div class="row">
                <div class="col">timestamp</div>
              </div>
            </div>
          </div>

        </div>
      </div>
    </div>
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
    report: function () {
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

<style>
.container {
  position: absolute;
  left: 0%;
}
.first_col {
  width: 30%;
  margin-right: 10%;
}
.row {
  margin-bottom: 20px;
}
.card {
  margin-bottom: 20px;
  padding: 20px;
}
</style>
