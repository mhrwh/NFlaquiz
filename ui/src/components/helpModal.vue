<template>
  <div class="modal fade" id="helpModal" tabindex="-1" aria-labelledby="helpModalTitle" aria-hidden="false">
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content help-modal">
        <div class="modal-header px-0 border-bottom-0">
          <h2 class="modal-title" id="helpModalTitle">ヘルプ</h2>
          <button type="button" class="close modal-close" data-dismiss="modal" aria-label="Close">
            <i class="bi bi-x-square" />
          </button>
        </div>
        <div class="modal-body p-0">
          <ul class="nav nav-tabs">
            <li class="nav-item">
              <a class="nav-link active" data-toggle="tab" href="#buttonDescription" @click="selectedTab = 'buttonDescription'">
                ボタンの説明
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link" data-toggle="tab" href="#tab2" @click="selectedTab = 'tab2'">
                ログイン機能
              </a>
            </li>
          </ul>
          <div class="tab-content tab-content-area">
            <div class="tab-pane fade show active" id="tab1" v-if="selectedTab === 'tab1'">
              <h4>
                <i class="bi bi-layout-sidebar-inset" />
                マップ切り替え
              </h4>
              <div class="text-danger">※この機能はログイン時のみ使用可能です。</div>
              マップには赤系統色で正答率を表すモードと緑系統色でブックマークした国を示すモードの2種類があります。<br><br>

              <p class="h4">
                <i class="bi bi-bookmark-heart" />
                ブックマーク編集
              </p>
              <div class="text-danger">※この機能はログイン時のみ使用可能です。</div>
              問題解答後にブックマークした国をブックマークしていない状態に戻すことができます。<br><br>

              <p class="h4">
                <i class="bi bi-flag-fill" />
                クイズ絞り込み
              </p>
              クイズを始める前に国の地域と国旗の色で出題される国を絞り込むことができます。
              出題される問題は10問です。<br><br>

            </div>
            <div class="tab-pane fade" id="tab2" v-if="selectedTab === 'tab2'">
              <p class="h4">
                ログイン方法
              </p>
              ヘッダー部の
              <i class="bi bi-person-fill" />
              アイコンを押し、アカウントをお持ちの方はフォームにメールアドレスとパスワードを入力するとログインできます。<br>
              アカウントをお持ちでない方はアイコンを押したのち、フォーム下部のアカウント登録リンクからアカウントを作成できます。<br><br>
              <p class="h4">
                ログインするとできること
              </p>
              ログインすることで以下の機能が使用可能になります。
              <ul>
                <li>マップ切り替え</li>
                <li>ブックマーク機能</li>
              </ul>
              なお、クイズはログイン状態関係なく遊ぶことができます。
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showMore: false,
      selectedTab: 'tab1'
    }
  }
}
</script>

<style>
.tab-content-area {
  background-color: #FFFFFF;
  height: 500px;
  overflow-y: auto;
}
</style>