<template>
  <b-row class="info-pane">
    <b-col sm="4">
      <p><strong>Username</strong>: {{ user.User.Username }}</p>
      <p><strong>Confirmation status</strong>: {{ user.User.UserStatus }}</p>
      <p><strong>Is enabled</strong>: {{ user.User.Enabled ? 'Yes' : 'No' }}</p>
    </b-col>
    <b-col sm="4">
      <p><b-button v-b-modal="modalId">Change password</b-button></p>
      <p>
        <spinning-button v-bind:is-spinning="isSignupConfirming" v-bind:click="confirmSignup" v-if="user.User.UserStatus === 'UNCONFIRMED'">Confirm</spinning-button>
      </p>
    </b-col>
    <b-col sm="4">
      <p><strong>MFA via</strong>: {{ user.MFAPreferences.EnabledOptions ? user.MFAPreferences.EnabledOptions.map(op => op.AttributeName).join(', ') : 'Disabled' }}</p>
      <p><strong>Created</strong>: {{ (new Date(user.User.UserCreateDate)).toLocaleString() }}</p>
      <p><strong>Last modified</strong>: {{ (new Date(user.User.UserLastModifiedDate)).toLocaleString() }}</p>
    </b-col>
  </b-row>
</template>

<script>
  import {BButton, BCol, BRow} from "bootstrap-vue";
  import SpinningButton from "./primitive/SpinningButton";

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
        }
      }
    },
    components: {
      'b-row': BRow,
      'b-col': BCol,
      'b-button': BButton,
      'spinning-button': SpinningButton,
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
