<template>
  <div>
    <b-breadcrumb :items="breadcrumbs" />
    <pool-table v-bind:user-pools="userPools" v-if="!isLoading" />
    <b-spinner variant="primary" label="Loading" v-if="isLoading" />
  </div>
</template>

<script>
import axios from 'axios';
import {BSpinner, BBreadcrumb} from 'bootstrap-vue';
import PoolTable from "../components/PoolTable";

export default {
  name: 'PoolList',
  components: {
    'b-breadcrumb': BBreadcrumb,
    'b-spinner': BSpinner,
    'pool-table': PoolTable,
  },
  data() {
    return {
      isLoading: false,
      userPools: {
        Pools: [],
      },
      breadcrumbs: [
        {
          text: 'User pools',
          to: { name: 'user-pools' },
        },
      ],
    };
  },
  mounted() {
    this.loadPools();
  },
  methods: {
    async loadPools() {
      try {
        this.isLoading = true;
        const poolList = await axios
          .get('http://localhost:8819/api/pools');
        this.userPools = poolList.data;
      } catch (error) {
        this.$bvToast.toast(error.response ? error.response.data.Message : error.message, {
          title: 'Error',
          autoHideDelay: 5000,
          variant: 'danger',
        });
      } finally {
        this.isLoading = false;
      }
    }
  },
}
</script>
