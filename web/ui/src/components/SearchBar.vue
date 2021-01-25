<template>
  <div class="search-bar">
    <b-form @submit.stop.prevent>
      <b-input-group>
        <b-form-select
          v-model="selectedSearchField"
          :options="searchFields"
          style="max-width: 25%"
          @change="handleSearchFieldChange"
        ></b-form-select>
        <b-form-input
          v-model="searchQuery"
          placeholder="Search for..."
          @update="handleSearchQueryUpdate"
          debounce="400"
        ></b-form-input>
      </b-input-group>
    </b-form>
  </div>
</template>

<script>
import { BForm, BFormSelect, BFormInput, BInputGroup } from 'bootstrap-vue';
import { detectSearchQueryType } from '../lib/utils/searchDetect';

const detectedQueryTypeToField = detectedType =>
  ({ uuid: 'sub', phone: 'phone_number', email: 'email' }[detectedType] || 'username');

export default {
  name: 'SearchBar',
  components: {
    'b-form': BForm,
    'b-form-select': BFormSelect,
    'b-form-input': BFormInput,
    'b-input-group': BInputGroup,
  },
  data() {
    return {
      searchFields: [
        { value: 'auto', text: 'Auto-detect' },
        { value: 'username', text: 'Username' },
        { value: 'email', text: 'Email' },
        { value: 'phone_number', text: 'Phone number' },
        { value: 'name', text: 'First name' },
        { value: 'given_name', text: 'Given name' },
        { value: 'family_name', text: 'Last name' },
        { value: 'preferred_username', text: 'Preferred username' },
        { value: 'sub', text: 'Sub (internal id)' },
      ],
      selectedSearchField: 'auto',
      searchQuery: '',
      detectedSearchField: 'unknown',
      searchChangeCallback: this.onSearchChange || (() => null),
    };
  },
  props: {
    onSearchChange: Function,
  },
  methods: {
    handleSearchQueryUpdate() {
      if (this.selectedSearchField === 'auto') {
        this.detectSearchQueryField();
      }

      this.querySearch();
    },
    handleSearchFieldChange() {
      if (this.selectedSearchField === 'auto') {
        this.detectSearchQueryField();
      } else {
        this.resetAutoFieldName();
      }

      this.querySearch();
    },
    resetAutoFieldName() {
      this.searchFields[0].text = 'Auto-detect';
    },
    clarifyAutoFieldName(clarifiedField) {
      this.searchFields[0].text = `Auto-detect (${clarifiedField})`;
    },
    detectSearchQueryField() {
      if (this.searchQuery === '') {
        this.resetAutoFieldName();
      } else {
        this.detectedSearchField = detectedQueryTypeToField(detectSearchQueryType(this.searchQuery));
        this.clarifyAutoFieldName(this.detectedSearchField.replace('_', ' '));
      }
    },
    querySearch() {
      this.searchChangeCallback(
        this.selectedSearchField === 'auto' ? this.detectedSearchField : this.selectedSearchField,
        this.searchQuery
      );
    },
  },
};
</script>

<style scoped>
.search-bar {
  margin-bottom: 20px;
}
</style>
