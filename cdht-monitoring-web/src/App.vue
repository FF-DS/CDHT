<template>
  <v-app>
    <div id="nav">
      <nav class="container navbar navbar-expand-lg navbar-light bg-light">
        <a class="navbar-brand" href="#">CDHT</a>
        <button
          class="navbar-toggler"
          type="button"
          data-toggle="collapse"
          data-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>

        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <div class="navbar-nav ml-auto">
            <router-link to="/report" class="nav-link">Report</router-link>
            <router-link to="/monitoring" class="nav-item nav-link"
              >Monitoring</router-link
            >
            <router-link to="/testing" class="nav-link">Testing</router-link>
          </div>
        </div>
      </nav>
    </div>

    <v-main> <router-view /></v-main>
  </v-app>
</template>

<script>
export default {
  name: "App",

  components: {},

  data: () => ({
    //
  }),

  computed: {},

  methods: {},

  created() {},
};
</script>
