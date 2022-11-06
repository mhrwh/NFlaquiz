<template>
  <div class="modal bd-example-modal-sm" id="loginModal" tabindex="-1" aria-labelledby="loginModalTitle">
    <div class="modal-dialog modal-sm">
      <div v-if="!registerActive" class="modal-content">
        <div class="modal-header">
          <h4 class="modal-title" id="loginModalTitle">ログイン</h4>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span>&times;</span>
          </button>
        </div>
        <div class="modal-body">
          <div v-bind:class="{ error: emptyFields }">
            <form>
              <div class="form-group">
                <label>メールアドレス</label>
                <input v-model="email" type="email" class="form-control" placeholder="Email" required>
              </div> 
              <div class="form-group">
                <label>パスワード</label>
                <input v-model="password" type="password" class="form-control" placeholder="Password" required>
              </div> 
              <div class="form-group">
                <input type="submit" class="btn btn-primary" @click="login">
                <br/>
                <a href="#" @click="registerActive = !registerActive, emptyFields = false">アカウント登録</a>
              </div>
            </form>
          </div>
        </div>
      </div>
      <div v-if="registerActive" class="modal-content">
        <div class="modal-header">
          <h4 class="modal-title" id="loginModalTitle">アカウント登録</h4>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span>&times;</span>
          </button>
        </div>
        <div class="modal-body">
          <div v-bind:class="{ error: emptyFields }">
            <form>
              <div class="form-group">
                <label>メールアドレス</label>
                <input v-model="email" type="email" class="form-control" placeholder="Email" required>
              </div> 
              <div class="form-group">
                <label>パスワード</label>
                <input v-model="password" type="password" class="form-control" placeholder="Password" required>
              </div> 
              <div class="form-group">
                <label>パスワード(確認)</label>
                <input v-model="password" type="password" class="form-control" placeholder="Password" required>
              </div> 
              <div class="form-group">
                <input type="submit" class="btn btn-primary" >
                <br/>
                <a href="#" @click="registerActive = !registerActive, emptyFields = false">ログイン画面へ</a>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="modal bd-example-modal-sm" id="logoutModal" tabindex="-1" aria-labelledby="logoutModalTitle">
    <div class="modal-dialog modal-sm">
      <div class="modal-content">
        <div class="modal-header">
          <h4 class="modal-title" id="logoutModalTitle">ログアウト</h4>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span>&times;</span>
          </button>
        </div>
        <div class="modal-body">
          <button type="button" class="btn btn-primary" @click="logout">ログアウト</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useStore } from 'vuex'

export default {
  name: 'LoginDialog',

  setup() {
    const store = useStore()
    const registerActive = ref (false)
    const email = ref("")
    const password = ref("")

    const login = async() => {
      console.log("login")
      store.dispatch("setAuth", true)
    }

    const logout = async() => {
      console.log("logout")
      store.dispatch("setAuth", false)
    }

    return {
      registerActive,
      email,
      password,
      login,
      logout,
    };
  }
};
</script>