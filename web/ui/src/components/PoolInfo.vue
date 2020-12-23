<template>
  <b-row class="info-pane">
    <b-col sm="4">
      <p><strong>Pool ID</strong>: {{ pool.Pool.Id }}</p>
      <p><strong>Pool name</strong>: {{ pool.Pool.Name }}</p>
      <p><strong>Pool ARN</strong>: {{ pool.Pool.Arn }}</p>
    </b-col>
    <b-col sm="4">
      <p><strong>Users can register themselves</strong>: {{ pool.Pool.AdminCreateUserConfig.AllowAdminCreateUserOnly ? 'No' : 'Yes' }}</p>
      <p><strong>Attribute used as a username</strong>: {{ pool.Pool.UsernameAttributes ? pool.Pool.UsernameAttributes.join(' or ') : 'username' }}</p>
      <p><strong>Username alias attributes</strong>: {{ pool.Pool.AliasAttributes ? pool.Pool.AliasAttributes.join(', ') : 'none' }}</p>
    </b-col>
    <b-col sm="4">
      <p><strong>Attributes to verify</strong>: {{ pool.Pool.AutoVerifiedAttributes.join(', ') }}</p>
      <p><strong>Created</strong>: {{ (new Date(pool.Pool.CreationDate)).toLocaleString() }}</p>
      <p><strong>Last modified</strong>: {{ (new Date(pool.Pool.LastModifiedDate)).toLocaleString() }}</p>
      <p><b-button v-b-modal="modalId">Create user</b-button></p>
    </b-col>
  </b-row>
</template>

<script>
  export default {
    name: 'PoolInfo',
    props: {
      modalId: String,
      pool: {
        type: Object,
        validator(value) {
          return !!value.Pool;
        },
      },
    },
  }
</script>

<style scoped>
  .info-pane {
    padding: 10px;
    background: #e9ecef;
    border-radius: 5px;
    margin-bottom: 20px;
    margin-left: auto;
    margin-right: auto;
  }
</style>
