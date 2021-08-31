<template>
  <v-app>
    <v-app-bar
      app
      absolute dense color="#141414"
      dark
    >
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="https://getbukkit.org/assets/img/getbukkitwhite.svg"
          transition="scale-transition"
          width="150"
        />
      </div>
    </v-app-bar>

    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  data: () => ({
    //
  }),
};
</script>

<style>
::-webkit-scrollbar {
    display: none;
}
</style>
