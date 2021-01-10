<template>
  <b-row class="info-pane">
    <b-col sm="8" class="user-header">
      <h3>
        {{ displayName }}
        <span v-b-tooltip title="User is disabled" v-if="!user.User.Enabled"><b-icon-x-circle-fill /></span>
      </h3>
      <h6><strong>Username</strong>: {{ user.User.Username }}</h6>
      <h6>
        <strong>Status</strong>:
        <span class="user-status" v-b-tooltip :title="statusDescription">{{ user.User.UserStatus }}</span>
      </h6>
      <h6>
        <strong>MFA via</strong>:
        {{
          user.MFAPreferences.EnabledOptions
            ? user.MFAPreferences.EnabledOptions.map(op => op.AttributeName).join(', ')
            : 'Disabled'
        }}
      </h6>
      <h6>
        <strong>Created</strong>:
        {{ new Date(user.User.UserCreateDate).toLocaleString() }}
        /
        <strong>Last modified</strong>:
        {{ new Date(user.User.UserLastModifiedDate).toLocaleString() }}
      </h6>
    </b-col>
    <b-col sm="4">
      <b-button-group vertical>
        <b-button v-b-modal="modalId">Change password</b-button>
        <spinning-button
          v-bind:is-spinning="isSignupConfirming"
          v-bind:click="confirmSignup"
          v-if="user.User.UserStatus === 'UNCONFIRMED'">
          Confirm
        </spinning-button>
      </b-button-group>
    </b-col>
  </b-row>
</template>

<script>
import { BButton, BCol, BRow, BButtonGroup, BIconXCircleFill } from 'bootstrap-vue';
import SpinningButton from './primitive/SpinningButton';
import { getAttribute } from '../lib/utils/attributes';
import { STATUS_DESCRIPTIONS } from '../lib/utils/userStatus';

export default {
  name: 'UserInfo',
  props: {
    modalId: String,
    isSignupConfirming: Boolean,
    confirmSignup: Function,
    user: {
      type: Object,
      validator(value) {
        return !!value.User;
      },
    },
  },
  computed: {
    displayName() {
      if (getAttribute(this.user.User.Attributes, 'name') && getAttribute(this.user.User.Attributes, 'family_name')) {
        return `${getAttribute(this.user.User.Attributes, 'name')} ${getAttribute(
          this.user.User.Attributes,
          'family_name'
        )}`;
      }

      if (getAttribute(this.user.User.Attributes, 'email')) {
        return getAttribute(this.user.User.Attributes, 'email');
      }

      if (getAttribute(this.user.User.Attributes, 'phone_number')) {
        return getAttribute(this.user.User.Attributes, 'phone_number');
      }

      return this.user.User.Username;
    },
    statusDescription() {
      return STATUS_DESCRIPTIONS[this.user.User.UserStatus] || 'Unknown';
    },
  },
  components: {
    'b-row': BRow,
    'b-col': BCol,
    'b-button': BButton,
    'b-icon-x-circle-fill': BIconXCircleFill,
    'b-button-group': BButtonGroup,
    'spinning-button': SpinningButton,
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

.user-header {
  text-align: left;
}

.user-status {
  text-decoration: underline dotted;
  text-decoration-thickness: from-font;
  cursor: pointer;
}
</style>
