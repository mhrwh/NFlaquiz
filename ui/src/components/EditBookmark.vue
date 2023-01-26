<template>
  <div class="modal fade" id="editBookmark" tabindex="-1" aria-labelledby="editBookmarkTitle" data-backdrop="static"
    aria-hidden="false">
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content edit-bookmark-modal">
        <div class="modal-header p-0 border-bottom-0">
          <h2 class="modal-title" id="editBookmarkTitle">ブックマーク編集</h2>
          <button type="button" class="close modal-close" data-dismiss="modal" @click="reload" aria-label="Close">
            <i class="bi bi-x-square"></i>
          </button>
        </div>
        <div class="modal-body p-0">
          <div class="mt-3">
            <div class="container">
              <div v-if="bookMarks.length > 0" class="table-responsive bookmarks-table bookmarks-country-name-col"
                style="overflow-y: scroll ">
                <table class="table table-hover">
                  <tbody>
                    <tr v-for="(value, index) in bookMarks" :key="value.name"
                      :class="{ 'even': index % 2 === 0, 'odd': index % 2 !== 0 }">
                      <td scope="col" class="align-middle w-70pct">
                        {{ value.name }}
                      </td>
                      <td scope="col" class="align-middle w-30pct">
                        <button class="btn btn-danger" v-on:click="deleteBookmark(value.id)">削除</button>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
              <div v-if="bookMarks.length == 0" class="bookmarks-table">
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

<style>
.odd {
  background-color: #E9CE3F;
}

.even {
  background-color: #F5F2E9;
  ;
}

.w-70pct {
  width: 70%
}

.w-30pct {
  width: 30%
}

.bookmarks-table {
  height: 450px
}

.bookmarks-country-name-col {
  font-size: 24px
}
</style>