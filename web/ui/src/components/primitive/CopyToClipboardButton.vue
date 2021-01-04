<template>
  <b-button variant="outline-secondary" size="sm" v-clipboard="valueToCopy">
    <b-icon-clipboard></b-icon-clipboard>
  </b-button>
</template>

<script>
import { BIconClipboard } from 'bootstrap-vue';

export default {
  name: 'CopyToClipboardButton',
  props: {
    valueToCopy: String,
  },
  components: {
    'b-icon-clipboard': BIconClipboard,
  },
};
</script>
