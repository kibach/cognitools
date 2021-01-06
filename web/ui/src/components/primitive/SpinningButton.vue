<template>
  <b-button @click="click" v-bind:variant="variant" v-bind:block="block">
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
    block: Boolean,
  },
};
</script>

<style scoped>
.button-spinner {
  margin-right: 10px;
  vertical-align: middle;
}
</style>
