<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div>
    <b-breadcrumb :items="breadcrumbs" />

    <user-info
        v-bind:user="user"
        v-bind:is-signup-confirming="isSignupConfirming"
        v-bind:confirm-signup="confirmSignup"
        modal-id="bv-password-change"
        v-if="!isLoading"
    />

    <user-attribute-editor-table
      v-bind:user="user"
      v-bind:pool-id="poolId"
      v-bind:is-attributes-updating="isAttributesUpdating"
      v-bind:update-attributes="updateAttributes"
      v-if="!isLoading"
    />

    <change-password-modal
      new-password=""
      v-bind:change-password="changePassword"
      modal-id="bv-password-change"
    />

    <b-spinner variant="primary" label="Loading" v-if="isLoading" />
  </div>
</template>

<script>
  import {BBreadcrumb, BSpinner} from "bootstrap-vue";
  import axios from "axios";
  import UserInfo from "../components/UserInfo";
  import ChangePasswordModal from "../components/modals/ChangePasswordModal";
  import UserAttributeEditorTable from "../components/UserAttributeEditorTable";
  import ToastsErrors from "../mixins/ToastsErrors";

  export default {
    name: 'UserView',
    components: {
      'b-breadcrumb': BBreadcrumb,
      'b-spinner': BSpinner,
      'user-info': UserInfo,
      'change-password-modal': ChangePasswordModal,
      'user-attribute-editor-table': UserAttributeEditorTable,
    },
    mixins: [ToastsErrors],
    data() {
      return {
        isLoading: false,
        isAttributesUpdating: false,
        isSignupConfirming: false,
        newPassword: '',
        poolId: this.$route.params.poolId,
        username: this.$route.params.username,
        user: {
          User: {
            Username: '',
            UserAttributes: [],
            UserStatus: '',
            Enabled: false,
            UserCreateDate: '',
            UserLastModifiedDate: '',
          },
          MFAPreferences: {
            EnabledOptions: [],
            PreferedOption: '',
          }
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
          {
            text: `User ${this.$route.params.username}`,
            to: { name: 'view-user', poolId: this.$route.params.poolId, username: this.$route.params.username },
          },
        ],
      };
    },
    mounted() {
      this.loadUser();
    },
    methods: {
      async loadUser() {
        try {
          this.isLoading = true;
          const userData = await axios
            .get(`http://localhost:8819/api/pools/${this.poolId}/users/${this.username}`);
          this.user = userData.data;
        } catch (error) {
          this.errorToast(error);
        } finally {
          this.isLoading = false;
        }
      },
      async updateAttributes() {
        try {
          this.isAttributesUpdating = true;
          const updatableAttributes = this.user.User.Attributes.filter(attr => attr.Name !== 'sub');
          await axios
            .post(`http://localhost:8819/api/pools/${this.poolId}/users/${this.username}/attributes`, updatableAttributes);
        } catch (error) {
          this.errorToast(error);
        } finally {
          this.isAttributesUpdating = false;
        }
      },
      async changePassword(newPassword) {
        try {
          const passwordUpdateRequest = { NewPassword: newPassword };
          await axios
            .post(`http://localhost:8819/api/pools/${this.poolId}/users/${this.username}/change_password`, passwordUpdateRequest);
        } catch (error) {
          this.errorToast(error);
        } finally {
          this.newPassword = '';
          this.$bvModal.hide('bv-password-change');
        }
      },
      async confirmSignup() {
        try {
          this.isSignupConfirming = true;
          await axios
            .post(`http://localhost:8819/api/pools/${this.poolId}/users/${this.username}/confirm_signup`);
          this.user.User.UserStatus = 'CONFIRMED';
        } catch (error) {
          this.errorToast(error);
        } finally {
          this.isSignupConfirming = false;
        }
      }
    }
  }
</script>
