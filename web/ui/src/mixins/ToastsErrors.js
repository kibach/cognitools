export default {
  methods: {
    errorToast(error) {
      this.$bvToast.toast(error.response ? error.response.data.Message : error.message, {
        title: 'Error',
        autoHideDelay: 10000,
        variant: 'danger',
      });
    },
  },
};
