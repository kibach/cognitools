<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div>
    <b-breadcrumb :items="breadcrumbs" />
    <pool-info v-bind:pool="pool" modal-id="bv-create-user" v-if="!isHeaderLoading" />
    <b-spinner variant="primary" label="Loading" v-if="isHeaderLoading" />
    <search-bar :on-search-change="searchQueryChange" />
    <pool-users-table
      v-bind:users="users"
      v-bind:pool-id="poolId"
      :on-load-more="() => loadUserList(users.PaginationToken)"
      :show-load-more="!isUserListExhausted"
      :is-busy="isUserListLoading"
    />

    <create-user-modal
      v-bind:register="createUser"
      v-bind:required-attributes="pool.Pool.SchemaAttributes.filter(attr => attr.Name !== 'sub')"
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
import SearchBar from '../components/SearchBar.vue';

export default {
  name: 'PoolView',
  components: {
    'b-breadcrumb': BBreadcrumb,
    'b-spinner': BSpinner,
    'pool-info': PoolInfo,
    'pool-users-table': PoolUsersTable,
    'create-user-modal': CreateUserModal,
    'search-bar': SearchBar,
  },
  mixins: [ToastsErrors],
  data() {
    return {
      isUserListLoading: false,
      isHeaderLoading: false,
      isUserListExhausted: false,
      searchField: '',
      searchQuery: '',
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
        PaginationToken: '',
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
    async loadPoolInfo() {
      try {
        this.isHeaderLoading = true;
        const poolData = await APIClient.get(`/pools/${this.poolId}`);
        this.pool = poolData.data;
      } catch (error) {
        this.errorToast(error);
      } finally {
        this.isHeaderLoading = false;
      }
    },
    async loadUserList(paginationToken = '') {
      try {
        if (this.isUserListExhausted) {
          return;
        }

        this.isUserListLoading = true;
        const after = encodeURIComponent(paginationToken) || '';
        const filterName = encodeURIComponent(this.searchField) || '';
        const filterValue = encodeURIComponent(this.searchQuery) || '';
        const userList = await APIClient.get(
          `/pools/${this.poolId}/users?after=${after}&filterName=${filterName}&filterValue=${filterValue}`
        );
        this.users.Users = this.users.Users.concat(userList.data.Users);
        this.users.PaginationToken = userList.data.PaginationToken;

        this.isUserListExhausted = !this.users.PaginationToken;
      } catch (error) {
        this.errorToast(error);
      } finally {
        this.isUserListLoading = false;
      }
    },
    async searchQueryChange(searchField, searchQuery) {
      const isRedundantChange = this.searchField !== searchField && !searchQuery;

      this.searchField = searchField;
      this.searchQuery = searchQuery;

      if (!isRedundantChange) {
        this.isUserListExhausted = false;
        this.users.Users = [];
        this.loadUserList();
      }
    },
  },
  async mounted() {
    this.loadPoolInfo();
    this.loadUserList();
  },
};
</script>
