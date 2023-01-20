<template>
  <div class="modal fade bd-example-modal-lg" id="editBookmark" tabindex="-1" aria-labelledby="editBookmarkTitle"
    aria-hidden="false">
    <div class="modal-dialog modal-lg modal-dialog-centered">
      <div class="modal-content quiz-filter">
        <div class="modal-header p-0 border-bottom-0">
          <h2 class="modal-title" id="editBookmarkTitle">ブックマーク編集</h2>
          <button type="button" class="close modal-close" data-dismiss="modal" @click="reload" aria-label="Close">
            <i class="bi bi-x-square"></i>
          </button>
        </div>
        <div class="modal-body p-0">
          <div class="container pl-32px">
            <div class="mt-3">
              <div v-if="bookMarks.length > 0" class="table-responsive" style="height: 600px; overflow-y: scroll;">
                <table class="table table-striped table-hover">
                  <tbody>
                    <tr v-for="value in bookMarks" :key="value.name">
                      <td>{{ value.name }}</td>
                      <td>
                        <button class="btn btn-danger" v-on:click="deleteBookmark(value.id)">削除</button>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
              <div v-if="bookMarks.length == 0">
                ブックマークがありません
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
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
        .then((res) => {
          bookMarks.value = res.data.map_info
          bookMarks.value = bookMarks.value.filter(bookmark => bookmark.bookmark === 1);
          bookMarks.value.sort((a, b) => {
            return a.name.localeCompare(b.name, 'ja');
          });
        });
    });
    const deleteBookmark = (id) => {
      axios
        .put(`http://localhost:8888/bookmark/${id}`)
        .then(() => {
          bookMarks.value = bookMarks.value.filter(bookmark => bookmark.id !== id);
        })
        .catch((error) => {
          console.log(error);
        });
    }

    const reload = () => {
      location.reload();
    }

    return {
      bookMarks,
      deleteBookmark,
      reload,
    }
  },
}
</script>
