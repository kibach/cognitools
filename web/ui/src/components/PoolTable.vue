<template>
  <b-table striped hover :items="userPools.Pools" :fields="tableColumns">
    <template v-slot:cell(CreationDate)="data">
      {{ new Date(data.value).toLocaleString() }}
    </template>
    <template v-slot:cell(LastModifiedDate)="data">
      {{ new Date(data.value).toLocaleString() }}
    </template>
    <template v-slot:cell(actions)="row">
      <b-button size="sm" :to="{ name: 'user-pool', params: { poolId: row.item.Id } }" class="mr-1"> View </b-button>
    </template>
  </b-table>
</template>

<script>
export default {
  name: 'PoolTable',
  props: {
    userPools: {
      type: Object,
      validator(value) {
        return !!value.Pools;
      },
    },
  },
  data() {
    return {
      tableColumns: [
        {
          key: 'Id',
          label: 'Pool ID',
        },
        {
          key: 'Name',
          label: 'Pool name',
        },
        {
          key: 'CreationDate',
          label: 'Created',
        },
        {
          key: 'LastModifiedDate',
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

<style scoped></style>
