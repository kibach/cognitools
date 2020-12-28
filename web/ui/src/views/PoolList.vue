<template>
  <div>
    <b-breadcrumb :items="breadcrumbs" />
    <pool-table v-bind:user-pools="userPools" v-if="!isLoading" />
    <b-spinner variant="primary" label="Loading" v-if="isLoading" />
  </div>
</template>

<script>
import APIClient from '../lib/client';
import {BSpinner, BBreadcrumb} from 'bootstrap-vue';
import PoolTable from "../components/PoolTable";
import ToastsErrors from "../mixins/ToastsErrors";

export default {
  name: 'PoolList',
  components: {
    'b-breadcrumb': BBreadcrumb,
    'b-spinner': BSpinner,
    'pool-table': PoolTable,
  },
  mixins: [ToastsErrors],
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
        const poolList = await APIClient
          .get('/pools');
        this.userPools = poolList.data;
      } catch (error) {
        this.errorToast(error);
      } finally {
        this.isLoading = false;
      }
    }
  },
}
</script>
