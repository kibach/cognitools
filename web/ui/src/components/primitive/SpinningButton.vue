<template>
  <span>
    <b-spinner small class="button-spinner" variant="primary" label="Processing" v-show="isSpinning" />
    <b-button @click="click" v-bind:variant="variant">
      <slot />
    </b-button>
  </span>
</template>

<script>
  export default {
    name: 'SpinningButton',
    props: {
      isSpinning: Boolean,
      click: Function,
      variant: String,
    }
  }
</script>

<style scoped>
  .button-spinner {
    margin-right: 10px;
  }
</style>