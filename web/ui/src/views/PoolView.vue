<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div>
    <b-breadcrumb :items="breadcrumbs" />
    <pool-info v-bind:pool="pool" modal-id="bv-create-user" v-if="!isLoading" />
    <pool-users-table v-bind:users="users" v-bind:pool-id="poolId" v-if="!isLoading" />
    <b-spinner variant="primary" label="Loading" v-if="isLoading" />

    <create-user-modal
      v-bind:register="createUser"
      v-bind:required-attributes="pool.Pool.SchemaAttributes.filter(attr => attr.Required && attr.Name !== 'sub')"
      modal-id="bv-create-user"
    />
  </div>
</template>

<style scoped></style>

<script>
import { BBreadcrumb, BSpinner } from 'bootstrap-vue';
import APIClient from '../lib/client';
import PoolInfo from '../components/PoolInfo';
import PoolUsersTable from '../components/PoolUsersTable';
import ToastsErrors from '../mixins/ToastsErrors';
import CreateUserModal from '../components/modals/CreateUserModal';

export default {
  name: 'PoolView',
  components: {
    'b-breadcrumb': BBreadcrumb,
    'b-spinner': BSpinner,
    'pool-info': PoolInfo,
    'pool-users-table': PoolUsersTable,
    'create-user-modal': CreateUserModal,
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
          SchemaAttributes: [],
          CreationDate: '',
          LastModifiedDate: '',
          Id: '',
          Arn: '',
          Name: '',
        },
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
  methods: {
    async createUser(username, password, attributes) {
      try {
        const user = await APIClient.post(`/pools/${this.poolId}/users`, {
          Username: username,
          Password: password,
          Attributes: attributes,
        });

        return this.$router.push({
          name: 'user-view',
          poolId: this.$route.params.poolId,
          username: user.data.User.Username,
        });
      } catch (error) {
        this.errorToast(error);
      }
    },
  },
  async mounted() {
    try {
      this.isLoading = true;
      const poolData = await APIClient.get(`/pools/${this.poolId}`);
      this.pool = poolData.data;

      const userList = await APIClient.get(`/pools/${this.poolId}/users`);
      this.users = userList.data;
    } catch (error) {
      this.errorToast(error);
    } finally {
      this.isLoading = false;
    }
  },
};
</script>
