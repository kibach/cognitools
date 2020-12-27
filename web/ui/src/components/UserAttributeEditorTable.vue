<template>
  <div>
    <b-table striped hover :items="user.User.Attributes" :fields="tableColumns">
      <template v-slot:cell(Value)="row" v-bind:poolId="poolId">
        <b-form-input v-if="row.item.Name !== 'sub'" v-model="row.item.Value" />
        <span v-if="row.item.Name === 'sub'">{{ row.item.Value }}</span>
      </template>
    </b-table>
    <b-row class="info-pane">
      <b-col sm="4" offset-sm="8" class="text-right">
        <spinning-button v-bind:is-spinning="isAttributesUpdating" v-bind:click="updateAttributes" variant="primary">Update attributes</spinning-button>
      </b-col>
    </b-row>
  </div>
</template>

<script>
  import {BFormInput, BTable} from "bootstrap-vue";
  import SpinningButton from "./primitive/SpinningButton";

  export default {
    name: 'UserAttributeEditorTable',
    props: {
      user: {
        type: Object,
        validator(value) {
          return !!value.User;
        }
      },
      poolId: String,
      isAttributesUpdating: Boolean,
      updateAttributes: Function,
    },
    components: {
      'b-table': BTable,
      'b-form-input': BFormInput,
      'spinning-button': SpinningButton,
    },
    data() {
      return {
        tableColumns: [
          {
            key: 'Name',
            label: 'Name',
          }, {
            key: 'Value',
            label: 'Value',
          },
        ],
      };
    }
  }
</script>
