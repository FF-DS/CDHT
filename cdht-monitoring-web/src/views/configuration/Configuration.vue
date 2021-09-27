<template>

  <div class="container">
    <div class="row">

      <!-- first column -->

      <div
        class="col"
        id="left"
      >
        <div class="row">
          Change jump spacing
        </div>
        <div class="row">
          <div class="col">
            <div class="mb-3">
              <input
                type="text"
                class="form-control"
                placeholder="jump space values"
              >
            </div>
          </div>
          <div class="col">
            <button
              class="btn btn-primary"
              type="submit"
            >save changes</button>
          </div>
        </div>
        <div class="row">
          <div
            class="card"
            id="scroll"
          >
            <div class="card-body">

            </div>

          </div>
        </div>
      </div>

      <!-- second column -->

      <div class="col">
        <div class="row">Node level replica counts</div>
        <div class="row">
          <div class="col">
            <div class="mb-3">
              <input
                type="text"
                class="form-control"
                placeholder="node id"
              >
            </div>
          </div>
          <div class="col">
            <button
              class="btn btn-primary"
              type="submit"
            >get replicas</button>
          </div>
        </div>
      </div>
    </div>
    <div>The data from the API is :- {{ config }}</div>
  </div>

</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "configuration",
  componenets: {},
  data: () => {
    return {};
  },
  computed: {
    ...mapGetters(["getConfig"]),

    config: function () {
      return this.getConfig;
    },
  },
  methods: {
    ...mapActions(["fetchConfig"]),
  },
  created() {
    this.fetchConfig();
  },
};
</script>


<style scoped>
#scroll {
  max-height: 100px;
  overflow-y: auto;
}
#left {
  width: 40px;
}
</style>
