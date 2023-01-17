<template>
  <div id="app">
    <table class="table table-bordered">
      <thead class="thead-dark">
        <tr>
          <th>国ID</th>
          <th>国名</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="value in bookMarks" :key="value.id">
          <td>{{ value.id }}</td>
          <td>{{ value.name }}</td>
          <td>{{ value.bookmark }}</td>
        </tr>
      </tbody>
    </table>
  </div>



  <div>
    <button class="btn btn-primary">bootstrap</button>
    <b-table striped hover :items="items"></b-table>
  </div>
</template>

<script>
import axios from 'axios'
import { ref, onMounted } from 'vue';
export default {
  name: 'EditBookmark',
  setup() {
    const bookMarks = ref([]);
    onMounted(() => {
      axios
        .get('http://localhost:8888/map')
        .then((res) => (bookMarks.value = res.data.map_info))
    });
    const fields = [
      'ID', 'Name', 'Mark'
    ]
    return {
      bookMarks, fields,
    }
  }
}


</script>