<template>
  <b-table striped hover :items="users.Users" :fields="tableColumns" :sort-by.sync="sortBy"
           :sort-desc.sync="sortDesc">
    <template v-slot:cell(UserCreateDate)="data">
      {{ (new Date(data.value)).toLocaleString() }}
    </template>
    <template v-slot:cell(UserLastModifiedDate)="data">
      {{ (new Date(data.value)).toLocaleString() }}
    </template>
    <template v-slot:cell(Enabled)="data">
      {{ data.value ? 'Yes' : 'No' }}
    </template>
    <template v-slot:cell(actions)="row" v-bind:poolId="poolId">
      <b-button size="sm" :to="{ name: 'user-view', params: { poolId: poolId, username: row.item.Username } }" class="mr-1">
        View
      </b-button>
    </template>
  </b-table>
</template>

<script>
  export default {
    name: 'PoolUsersTable',
    props: {
      users: {
        type: Object,
        validator(value) {
          return Array.isArray(value.Users);
        },
      },
      poolId: String,
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
          }, {
            key: 'UserStatus',
            label: 'Status',
          }, {
            key: 'Enabled',
            label: 'Is enabled',
          }, {
            key: 'UserCreateDate',
            label: 'Created',
          }, {
            key: 'UserLastModifiedDate',
            label: 'Last modified',
          }, {
            key: 'actions',
            label: 'Actions',
          }
        ],
      };
    },
  }
</script>
