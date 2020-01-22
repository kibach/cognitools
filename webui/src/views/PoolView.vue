<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div>
    <b-breadcrumb :items="breadcrumbs" />
    <pool-info v-bind:pool="pool" v-if="!isLoading" />
    <pool-users-table v-bind:users="users" v-bind:pool-id="poolId" v-if="!isLoading" />
    <b-spinner variant="primary" label="Loading" v-if="isLoading" />
  </div>
</template>

<style scoped>

</style>

<script>
  import {BBreadcrumb, BSpinner} from "bootstrap-vue";
  import axios from "axios";
  import PoolInfo from "../components/PoolInfo";
  import PoolUsersTable from "../components/PoolUsersTable";
  import ToastsErrors from "../mixins/ToastsErrors";

  export default {
    name: 'PoolView',
    components: {
      'b-breadcrumb': BBreadcrumb,
      'b-spinner': BSpinner,
      'pool-info': PoolInfo,
      'pool-users-table': PoolUsersTable,
    },
    mixins: [ToastsErrors],
    data() {
      return {
        isLoading: false,
        poolId: this.$route.params.poolId,
        pool: {
          Pool: {
            AdminCreateUserConfig: {},
            AllowAdminCreateUserOnly: false,
            AutoVerifiedAttributes: [],
            UsernameAttributes: [],
            AliasAttributes: [],
            CreationDate: '',
            LastModifiedDate: '',
            Id: '',
            Arn: '',
            Name: '',
          }
        },
        users: {
          Users: [],
        },
        breadcrumbs: [
          {
            text: 'User pools',
            to: { name: 'user-pools' },
          },
          {
            text: `Pool ${this.$route.params.poolId}`,
            to: { name: 'user-pool', poolId: this.$route.params.poolId },
          },
        ],
      };
    },
    async mounted() {
      try {
        this.isLoading = true;
        const poolData = await axios
          .get(`http://localhost:8819/api/pools/${this.poolId}`);
        this.pool = poolData.data;

        const userList = await axios
          .get(`http://localhost:8819/api/pools/${this.poolId}/users`);
        this.users = userList.data;
      } catch (error) {
        this.errorToast(error);
      } finally {
        this.isLoading = false;
      }
    }
  }
</script>
