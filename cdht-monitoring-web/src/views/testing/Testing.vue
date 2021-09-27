<template>
  <div class="container">
    <div class="row">
      <div class="col left">

        <div class="card">
          <div class="card-body">
            <div class="col">
              <div class="row">
                <div class="card">
                  <div class="card-body">Run test tools on a node</div>
                </div>
              </div>
              <div class="row">
                <div class="col">
                  <button
                    class="btn btn-primary"
                    type="submit"
                  >Run ping tool</button>
                </div>
                <div class="col">
                  <div class="mb-3">
                    <input
                      type="text"
                      class="form-control"
                      placeholder="node id"
                    >
                  </div>
                </div>
              </div>
              <div class="row">
                <div class="card">
                  <div class="card-body">
                    <div class="col">
                      Node Status: Alive
                    </div>
                    <div class="col">
                      Node address: 193.45.0.89
                    </div>
                  </div>
                </div>
              </div>
              <div class="row">
                <div class="card">
                  <div class="card-body">
                    <div class="col">
                      <button
                        class="btn btn-primary"
                        type="submit"
                      >Run DNS lookup</button>
                    </div>
                    <div class="col">
                      <div class="mb-3">
                        <input
                          type="text"
                          class="form-control"
                          placeholder="node id"
                        >
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="row">
                <div class="card">
                  <div class="card-body">
                    <div class="col">
                      Node Status: Alive
                    </div>
                    <div class="col">
                      Node address: 193.45.0.89
                    </div>
                  </div>
                </div>
              </div>
              <div class="row">
                <div class="card">
                  <div class="card-body">
                    <div class="col">
                      <button
                        class="btn btn-primary"
                        type="submit"
                      >Run hop count</button>
                    </div>
                    <div class="col">
                      <div class="row">
                        <div class="mb-3">
                          <input
                            type="text"
                            class="form-control"
                            placeholder="start node id"
                          >
                        </div>
                      </div>
                      <div class="row">
                        <div class="mb-3">
                          <input
                            type="text"
                            class="form-control"
                            placeholder="destination node id"
                          >
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="row">
                <div class="card">
                  <div class="card-body">
                    <div class="col">
                      Node Status: Alive
                    </div>
                    <div class="col">
                      Node address: 193.45.0.89
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="col">
        <div class="row">
          response log from the tool...
        </div>
        <div class="row">
          <div class="card scroll">
            <div class="card-body">

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
  name: "testing",
  componenets: {},
  data: () => {
    return {};
  },
  computed: {
    ...mapGetters(["getTest"]),
    test: function () {
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

<style scoped>
.left {
  width: 33%;
}
.scroll {
  max-height: 100%;
  overflow-y: auto;
}
</style>
