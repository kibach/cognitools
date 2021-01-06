<template>
  <b-button variant="outline-secondary" v-bind:size="size" v-clipboard="valueToCopy">
    <b-icon-clipboard></b-icon-clipboard>
  </b-button>
</template>

<script>
import { BIconClipboard } from 'bootstrap-vue';

export default {
  name: 'CopyToClipboardButton',
  props: {
    valueToCopy: String,
    size: {
      type: String,
      default: 'sm',
    },
  },
  components: {
    'b-icon-clipboard': BIconClipboard,
  },
};
</script>
