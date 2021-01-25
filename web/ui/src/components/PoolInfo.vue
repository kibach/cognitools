<template>
  <b-row class="info-pane">
    <b-col sm="8" class="pool-header">
      <h3>
        {{ pool.Pool.Name }}
        <span
          v-b-tooltip
          title="Users can register themselves"
          v-if="!pool.Pool.AdminCreateUserConfig.AllowAdminCreateUserOnly"
        >
          <b-icon-person-check />
        </span>
        <span
          v-b-tooltip
          title="Users cannot register themselves"
          v-if="pool.Pool.AdminCreateUserConfig.AllowAdminCreateUserOnly"
        >
          <b-icon-person-x />
        </span>
      </h3>
      <h6>
        <strong>ID</strong>: {{ pool.Pool.Id }}
        <copy-to-clipboard-button class="size-xs" :valueToCopy="pool.Pool.Id" />
      </h6>
      <h6>
        <strong>ARN</strong>: {{ pool.Pool.Arn }}
        <copy-to-clipboard-button class="size-xs" :valueToCopy="pool.Pool.Arn" />
      </h6>
      <h6>
        <strong>Created</strong>:
        {{ new Date(pool.Pool.CreationDate).toLocaleString() }}
        /
        <strong>Last modified</strong>:
        {{ new Date(pool.Pool.LastModifiedDate).toLocaleString() }}
      </h6>
    </b-col>
    <b-col sm="4" class="pool-header">
      <p><b-button v-b-modal="modalId">Create user</b-button></p>
      <h6>
        <strong>Attributes to verify</strong>:
        {{ pool.Pool.AutoVerifiedAttributes ? pool.Pool.AutoVerifiedAttributes.join(', ') : 'none' }}
      </h6>
      <h6>
        <strong>Attribute used as a username</strong>:
        {{ pool.Pool.UsernameAttributes ? pool.Pool.UsernameAttributes.join(' or ') : 'username' }}
      </h6>
      <h6>
        <strong>Username alias attributes</strong>:
        {{ pool.Pool.AliasAttributes ? pool.Pool.AliasAttributes.join(', ') : 'none' }}
      </h6>
    </b-col>
  </b-row>
</template>

<script>
import CopyToClipboardButton from './primitive/CopyToClipboardButton.vue';
import { BIconPersonCheck, BIconPersonX } from 'bootstrap-vue';

export default {
  name: 'PoolInfo',
  components: {
    'copy-to-clipboard-button': CopyToClipboardButton,
    'b-icon-person-check': BIconPersonCheck,
    'b-icon-person-x': BIconPersonX,
  },
  props: {
    modalId: String,
    pool: {
      type: Object,
      validator(value) {
        return !!value.Pool;
      },
    },
  },
};
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

.pool-header {
  text-align: left;
}

.size-xs {
  padding: 0.2rem 0.4rem;
  font-size: 0.625rem;
  line-height: 1.2;
  border-radius: 0.2rem;
  vertical-align: bottom;
}
</style>
