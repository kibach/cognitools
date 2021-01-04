<template>
  <b-button @click="click" v-bind:variant="variant">
    <b-spinner small class="button-spinner" variant="primary" label="Processing" v-show="isSpinning" />
    <slot />
  </b-button>
</template>

<script>
export default {
  name: 'SpinningButton',
  props: {
    isSpinning: Boolean,
    click: Function,
    variant: String,
  },
};
</script>

<style scoped>
.button-spinner {
  margin-right: 10px;
}
</style>
