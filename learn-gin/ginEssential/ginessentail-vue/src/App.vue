<template>
  <div id="app">
    <navbar />
    <b-container>
      <router-view />
    </b-container>
  </div>
</template>

<script>
import Navbar from "./views/layout/Navbar.vue";

export default {
  components: { Navbar },
  data() {
    return {};
  },
};
</script>


<style lang="scss">
// #app {
//   font-family: Avenir, Helvetica, Arial, sans-serif;
//   -webkit-font-smoothing: antialiased;
//   -moz-osx-font-smoothing: grayscale;
//   text-align: center;
//   color: #2c3e50;
// }

// #nav {
//   padding: 30px;

//   a {
//     font-weight: bold;
//     color: #2c3e50;

//     &.router-link-exact-active {
//       color: #42b983;
//     }
//   }
// }
</style>
