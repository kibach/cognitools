<template>
  <b-modal v-bind:id="modalId" hide-footer>
    <template v-slot:modal-title> Change password </template>
    <b-form-input v-model="newPassword" placeholder="New password" type="password" />
    <b-button class="mt-3" variant="primary" block @click="() => changePassword(newPassword)">Change</b-button>
  </b-modal>
</template>

<script>
export default {
  name: 'ChangePasswordModal',
  props: {
    newPassword: String,
    changePassword: Function,
    modalId: String,
  },
};
</script>
