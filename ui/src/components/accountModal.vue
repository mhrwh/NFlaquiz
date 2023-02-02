<template>
  <div class="modal" id="loginModal" tabindex="-1" aria-labelledby="accountModalTitle">
    <div class="modal-dialog account-modal-dialog modal-sm">
      <div v-if="!registerActive" class="modal-content account-modal">
        <div class="modal-header account-modal-header">
          <p class="modal-title" id="accountModalTitle">ログイン</p>
          <button type="button" class="close modal-close" data-dismiss="modal" aria-label="Close">
            <i class="bi bi-x-lg" style="font-size: 12px" />
          </button>
        </div>
        <div class="modal-body p-0">
          <form>
            <div class="account-form-body">
              <div class="form-group login-form-content">
                <label class="font-10px d-flex">
                  <i class="bi bi-envelope-plus mr-3px" />
                  メールアドレス
                </label>
                <input v-model="email" type="email" class="form-control form-control-sm" required>
              </div> 
              <div class="form-group login-form-content">
                <label class="font-10px d-flex">
                  <i class="bi bi-ui-radios-grid mr-3px" />
                  パスワード
                </label>
                <div class="d-flex">
                  <input v-model="password" :type="visible ? 'text' : 'password'" class="form-control form-control-sm" required>
                  <i class="bi pass-view" :class="visible ? 'bi-eye-slash' : 'bi-eye'" @click="visible = !visible" />
                </div>
              </div>
              <p class="font-10px text-danger d-flex">
                <i class="bi bi-exclamation-circle mr-3px" :class="errmsg ? 'd-block' : 'd-none'" />
                {{ errmsg }}
              </p>
            </div>
            <div class="account-form-footer">
              <button type="submit" class="btn btn-primary btn-account font-10px" @click.prevent="login">
                <i class="bi bi-envelope-paper" />
                送信
              </button>
              <br/>
              <a href="#" class="font-10px" @click="registerActive = !registerActive">アカウント登録</a>
            </div>
          </form>
        </div>
      </div>
      <div v-if="registerActive" class="modal-content account-modal">
        <div class="modal-header account-modal-header">
          <p class="modal-title" id="accountModalTitle">アカウント登録</p>
          <button type="button" class="close modal-close" data-dismiss="modal" aria-label="Close">
            <i class="bi bi-x-lg" style="font-size: 12px" />
          </button>
        </div>
        <div class="modal-body p-0">
          <form>
            <div class="account-form-body">
              <div class="form-group signup-form-content">
                <label class="font-10px d-flex">
                  <i class="bi bi-envelope-plus mr-3px" />
                  メールアドレス
                </label>
                <input v-model="registerEmail" type="email" class="form-control form-control-sm" required>
              </div> 
              <div class="form-group signup-form-content">
                <label class="font-10px d-flex">
                  <i class="bi bi-ui-radios-grid mr-3px" />
                  パスワード
                </label>
                <div class="d-flex">
                  <input v-model="registerPassword" :type="visible ? 'text' : 'password'" class="form-control form-control-sm" required>
                  <i class="bi pass-view" :class="visible ? 'bi-eye-slash' : 'bi-eye'" @click="visible = !visible" />
                </div>
              </div>
              <div class="form-group signup-form-content">
                <label class="font-10px d-flex">
                  <i class="bi bi-ui-radios-grid mr-3px" />
                  パスワード(確認)
                </label>
                <div class="d-flex">
                  <input v-model="registerPasswordConfirm" :type="visible ? 'text' : 'password'" class="form-control form-control-sm" required>
                  <i class="bi pass-view" :class="visible ? 'bi-eye-slash' : 'bi-eye'" @click="visible = !visible" />
                </div>
              </div>
              <p class="font-10px text-danger d-flex">
                <i class="bi bi-exclamation-circle mr-3px" :class="registerErrmsg ? 'd-block' : 'd-none'" />
                {{ registerErrmsg }}
              </p>
            </div>
            <div class="account-form-footer">
              <button type="submit" class="btn btn-primary btn-account font-10px" @click.prevent="signup">
                <i class="bi bi-envelope-paper" />
                送信
              </button>
              <br/>
              <a href="#" class="font-10px" @click="registerActive = !registerActive">ログイン画面へ</a>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import axios from 'axios';
import { useStore } from 'vuex';

export default {
  name: 'AccountModal',

  setup() {
    const store = useStore();
    const registerActive = ref (false);
    const visible = ref (false);
    const email = ref("");
    const password = ref("");
    const registerEmail = ref("");
    const registerPassword = ref("");
    const registerPasswordConfirm = ref("");
    let errmsg = ref("");
    let registerErrmsg = ref("");

    const login = async() => {
      const params = new URLSearchParams();
      params.append('email', email.value);
      params.append('password', password.value);
      await axios.post('http://localhost:8888/login',params,)
      .then(() => {
        store.dispatch("setAuth", true);
        location.reload();
      })
      .catch(error => {
        errmsg.value = error.response.data.err_msg;
      });
    }

    const signup = async() => {
      const params = new URLSearchParams();
      params.append('email', registerEmail.value);
      params.append('password', registerPassword.value);
      params.append('password_confirm', registerPasswordConfirm.value);
      await axios.post('http://localhost:8888/signup',params,)
      .then(() => {
        store.dispatch("setAuth", true);
        location.reload();
      })
      .catch(error => {
        registerErrmsg.value = error.response.data.err_msg;
      });
    }

    return {
      registerActive,
      visible,
      email,
      password,
      registerEmail,
      registerPassword,
      registerPasswordConfirm,
      login,
      signup,
      errmsg,
      registerErrmsg,
    };
  },
};
</script>
<style>
.account-form-body {
  padding: 4px 15px 4px 17px;
  min-height: 201px;
}
.account-form-body p {
  margin: 0;
}
.account-form-footer {
  border-top: 1px solid #dee2e6;
  padding: 10px 15px 8px 17px;
}
.pass-view{
  margin-left: 7px;
  cursor: pointer;
}
</style>