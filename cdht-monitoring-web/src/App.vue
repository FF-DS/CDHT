<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <v-btn href="http://localhost:8080" target="_blank" text>
        <span class="mr-2">CDHT</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
    </v-app-bar>

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
