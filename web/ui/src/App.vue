<template>
  <div id="app">
    <b-container class="top-level-container">
      <router-view />
    </b-container>
  </div>
</template>

<script>
import { BContainer } from 'bootstrap-vue';

export default {
  name: 'App',
  components: {
    'b-container': BContainer,
  },
};
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.top-level-container {
  margin-top: 40px;
  margin-bottom: 20px;
}
</style>
