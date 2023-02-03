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
              <a class="nav-link" data-toggle="tab" href="#loginDescription" @click="selectedTab = 'loginDescription'">
                ログイン機能
              </a>
            </li>
          </ul>
          <div class="tab-content tab-content-area p-3">
            <div class="tab-pane fade show active" id="buttonDescription" v-if="selectedTab === 'buttonDescription'">
              <h4>
                <i class="bi bi-layout-sidebar-inset" />
                マップ切り替え
              </h4>
              <div class="text-danger">※この機能はログイン時のみ使用可能です。</div>
              <p>
                ログインすると地図が赤もしくは緑系統色で色分けされています。<br>
                赤系統色は正答率を表すモードで緑系統色はブックマークした国を示すモードの2種類があり、このボタンを押すことで切り替えることが可能です。
              </p>

              <h4>
                <i class="bi bi-bookmark-heart" />
                ブックマーク編集
              </h4>
              <div class="text-danger">※この機能はログイン時のみ使用可能です。</div>
              <p>
                問題解答後にブックマークした国をブックマークしていない状態に戻すことができます。
              </p>

              <h4>
                <i class="bi bi-flag-fill" />
                クイズ絞り込み
              </h4>
              <p>
                このボタンからクイズを始めることができます。<br>
                クイズを始める前に国の地域と国旗の色で出題される国を絞り込むことができます。
                出題される問題は10問です。
              </p>

            </div>
            <div class="tab-pane fade" id="loginDescription" v-if="selectedTab === 'loginDescription'">
              <h4>
                ログイン方法
              </h4>
              <p>
                画面右上の
                <i class="bi bi-person-fill" />
                アイコンを押し、アカウントをお持ちの方はフォームにメールアドレスとパスワードを入力するとログインできます。<br>
                アカウントをお持ちでない方はアイコンを押したのち、フォーム下部のアカウント登録リンクからアカウントを作成できます。
              </p>

              <h4>
                ログアウト方法
              </h4>
              <p>
                ログイン状態で画面右上の
                <i class="bi bi-person-check-fill" />
                アイコンを押すとログアウトできます。
              </p>

              <h4>
                ログインするとできること
              </h4>
              <p>
                ログインすることで以下の機能が使用可能になります。
                <ul>
                  <li>地図が色分けされる</li>
                  <li>地図表示切り替え</li>
                  <li>ブックマーク追加・削除</li>
                </ul>
                なお、クイズはログイン状態関係なく遊ぶことができます。
              </p>
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
      selectedTab: 'buttonDescription'
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