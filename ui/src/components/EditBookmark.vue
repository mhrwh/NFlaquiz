<template>
  <div class="modal fade bd-example-modal-lg" id="editBookmark" tabindex="-1" aria-labelledby="editBookmarkTitle"
    aria-hidden="false">
    <div class="modal-dialog modal-lg modal-dialog-centered">
      <div class="modal-content quiz-filter">
        <div class="modal-header p-0 border-bottom-0">
          <h2 class="modal-title" id="editBookmarkTitle">ブックマーク編集</h2>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
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
                      <td>{{ value.bookmark }}</td>
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
export default {
  name: 'EditBookmark',
  props: {
    bookMarks: {
      type: Array,
      default: () => []
    }
  },
  methods: {
    deleteBookmark(country_id) {
      this.$emit('delete-bookmark', country_id);
    }
  }
}
</script>