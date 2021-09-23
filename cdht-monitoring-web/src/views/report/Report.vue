<template>
  <div>
    <div class="row">
      <div class="col-4">kotet</div>
      <div class="col-8">
        <input
          type="number"
          class="form-control "
          v-model="requestBody.limit"
        />
        <button
          type="submit"
          @click="
            () => {
              this.fetchReports({
                limit: requestBody.limit,
              });
            }
          "
          class="btn float-left btn-outline-secondary py-2 px-3"
          style="border:none"
        >
          <span class="button-text">
            Reload
            <i class="fal fa-times-circle px-1"></i>
          </span>
        </button>
        <div
          v-for="(model, index) in report"
          :key="index"
          :style="'cursor:pointer'"
        >
          This is model # {{ index }}
          <br />
          created_date => {{ model.created_date }}
          <br />
          report_type => {{ model.type }}
          <br />
          node_id => {{ model.node_id }}

          <br />
          <br />
          <!-- <div v-for="(value, key) in awards[index]" :key="key"> -->
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
      limit: 10,
      requestBody: {
        limit: 10,
      },
    };
  },
  computed: {
    ...mapGetters(["getAllReports"]),
    report: function() {
      return this.getAllReports;
    },
  },
  methods: {
    ...mapActions(["fetchReports"]),

    reloadAll() {
      console.log("the limit is now", this.requestBody.limit);
      this.fetchReports(this.limit);
    },
  },
  created() {
    this.fetchReports(this.requestBody);
  },
};
</script>

<style></style>
