<template>
  <div>
    <b-table striped hover :items="users.Users" :fields="tableColumns" no-local-sorting :busy.sync="isBusy">
      <template v-slot:cell(UserCreateDate)="data">
        {{ new Date(data.value).toLocaleString() }}
      </template>
      <template v-slot:cell(UserLastModifiedDate)="data">
        {{ new Date(data.value).toLocaleString() }}
      </template>
      <template v-slot:cell(Enabled)="data">
        {{ data.value ? 'Yes' : 'No' }}
      </template>
      <template v-slot:cell(actions)="row" v-bind:poolId="poolId">
        <b-button
          size="sm"
          :to="{
            name: 'user-view',
            params: { poolId: poolId, username: row.item.Username },
          }"
          class="mr-1"
        >
          View
        </b-button>
      </template>
    </b-table>
    <spinning-button block variant="outline-primary" :click="onLoadMore" :is-spinning="isBusy" v-if="showLoadMore"
      >Load more</spinning-button
    >
  </div>
</template>

<script>
import SpinningButton from './primitive/SpinningButton';

export default {
  name: 'PoolUsersTable',
  components: {
    'spinning-button': SpinningButton,
  },
  props: {
    users: {
      type: Object,
      validator(value) {
        return Array.isArray(value.Users);
      },
    },
    poolId: String,
    onLoadMore: Function,
    showLoadMore: Boolean,
    isBusy: Boolean,
  },
  data() {
    return {
      sortBy: 'Username',
      sortDesc: false,
      tableColumns: [
        {
          key: 'Username',
          label: 'Username',
          sortable: true,
        },
        {
          key: 'UserStatus',
          label: 'Status',
        },
        {
          key: 'Enabled',
          label: 'Is enabled',
        },
        {
          key: 'UserCreateDate',
          label: 'Created',
        },
        {
          key: 'UserLastModifiedDate',
          label: 'Last modified',
        },
        {
          key: 'actions',
          label: 'Actions',
        },
      ],
    };
  },
};
</script>
