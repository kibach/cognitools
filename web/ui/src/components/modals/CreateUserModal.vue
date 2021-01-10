<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <b-modal hide-footer v-bind:id="modalId">
    <template v-slot:modal-title> Create user </template>

    <b-form-group>
      <b-form-input placeholder="Username" type="text" v-model="username" id="input-Username" />
    </b-form-group>
    <b-form-group>
      <b-form-input placeholder="Password" type="password" v-model="password" />
    </b-form-group>

    <b-form-group v-for="attribute in requiredAttributes" v-bind:key="attribute.Name">
      <b-form-input v-bind:placeholder="attribute.Name" type="text" v-model="attributes[attribute.Name]" />
    </b-form-group>

    <b-button @click="() => register(username, password, getAttributeList())" block class="mt-3" variant="primary">
      Create
    </b-button>
  </b-modal>
</template>

<script>
export default {
  name: 'CreateUserModal',
  data() {
    return {
      username: '',
      password: '',
      attributes: {},
    };
  },
  props: {
    register: Function,
    modalId: String,
    requiredAttributes: Array,
  },
  methods: {
    getAttributeList() {
      return Object.entries(this.attributes).map(([Name, Value]) => ({
        Name,
        Value,
      }));
    },
  },
};
</script>
