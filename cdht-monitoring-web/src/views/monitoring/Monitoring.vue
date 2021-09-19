<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <div class="row graph">
          <div class="card">

          </div>
        </div>
        <div class="row graph">
          <div class="card"></div>
        </div>
      </div>
      <div class="col">
        <div class="card">
          <div class="card-body">

            <div class="row">
              <div class="col-sm-10">
                <div class="mb-3">
                  <input
                    type="text"
                    class="form-control"
                    id="formGroupExampleInput"
                    placeholder="end date"
                  >
                </div>
              </div>
              <div class="col">
                <button
                  class="btn btn-primary"
                  type="submit"
                >ok</button>
              </div>
            </div>
            <div
              class="row"
              id="lable1"
            >
              Real-time node statistics
            </div>
            <div class="row">
              <div
                class="card"
                id="detail"
              >
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
    </div>
    <div>The data from the API is :- {{ monitor }}</div>
  </div>

</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "monitoring",
  componenets: {},
  data: () => {
    return {};
  },
  computed: {
    ...mapGetters(["getMonitor"]),
    monitor: function () {
      return this.getMonitor;
    },
  },
  methods: {
    ...mapActions(["fetchMonitor"]),
  },
  created() {
    this.fetchMonitor();
  },
};
</script>

<style>
.container {
}
.graph {
  height: 400px;
  margin-bottom: 2%;
}
#lable1 {
  position: inherit;
  color: blueviolet;
  left: 30px;
}
#detail {
  position: inherit;
}
</style>
