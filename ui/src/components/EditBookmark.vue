<template>
  <div class="modal fade" id="editBookmark" tabindex="-1" aria-labelledby="editBookmarkTitle" data-backdrop="static"
    aria-hidden="false">
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content edit-bookmark-modal pb-5">
        <div class="modal-header px-4 border-bottom-0">
          <h2 class="modal-title" id="editBookmarkTitle">ブックマーク編集</h2>
          <button type="button" class="close modal-close" data-dismiss="modal" @click="reload" aria-label="Close">
            <i class="bi bi-x-square"></i>
          </button>
        </div>
        <div class="modal-body p-0">
          <div class="border-top border-bottom border-secondary">
            <div v-if="bookMarks.length > 0" class="table-responsive bookmarks-table bookmarks-country-name-col mb-0"
              style="overflow-y: scroll ">
              <table class="table table-hover">
                <tbody>
                  <tr v-for="(value, index) in bookMarks" :key="value.name"
                    :class="{ 'even': index % 2 === 0, 'odd': index % 2 !== 0 }">
                    <td scope="col" class="align-middle w-70pct pl-5">
                      {{ value.name }}
                    </td>
                    <div class="text-right">
                      <td scope="col" class="align-middle pr-5 w-30pct ">
                        <button class="btn btn-danger" v-on:click="deleteBookmark(value.id)">削除</button>
                      </td>
                    </div>
                  </tr>
                </tbody>
              </table>
            </div>
            <div v-if="bookMarks.length == 0" class="pl-5 bookmarks-table">
              ブックマークがありません
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
}


.pd-y-20px {
  padding: 0px 20px
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