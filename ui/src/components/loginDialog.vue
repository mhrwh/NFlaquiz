<template>
    
    <button class="btn btn-primary" type="button" @click="isShow = ! isShow"  style="position: absolute; right: 0px; z-index: 10">ログイン</button>
    <transition>
      <div class="dialog"  v-if="isShow" > 
        <div class="container">
          <div class="row">
              <div class="col" style="padding: 0;">
                <div v-if="!registerActive" class="card" v-bind:class="{ error: emptyFields }">
                  <article class="card-body">
                    <button type="button" class="float-right close" aria-label="Close" @click="isShow = ! isShow">
                      <span aria-hidden="true">&times;</span>
                    </button>
                    <h4 class="card-title mb-4 mt-1">Sign in</h4>
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
                  </article>
                </div> 

                <div v-else class="card" v-bind:class="{ error: emptyFields }">
                  <article class="card-body">
                    <button type="button" class="float-right close" aria-label="Close" @click="isShow = ! isShow">
                      <span aria-hidden="true">&times;</span>
                    </button>
                    <h4 class="card-title mb-4 mt-1">Sign up</h4>
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
                        <a href="#" @click="registerActive = !registerActive, emptyFields = false">Sign In</a>
                      </div>
                    </form>
                  </article>
                </div>
              </div>
          </div>

        </div>
      </div>
    </transition>
</template>

<script>
import { ref } from 'vue'
import axios from 'axios'
import { useStore } from 'vuex'

export default {
  name: 'LoginDialog',

  setup() {
    const store = useStore()
    const isShow = ref(false)
    const registerActive = ref (false)
    const email = ref("")
    const password = ref("")

    const login = async() => {
      const params = new URLSearchParams();
      params.append('email', email.value);
      params.append('password', password.value);
      console.log(params);
      axios.post('http://localhost:8888/login',params,)
      .then(res => {
          console.log(res);
          //msg.value = res.data.msg;
          store.dispatch("setAuth", true)
      })
      .catch(error => {
          console.log(error.response.data.err_msg);
          //msg.value = error.response.data.err_msg;
      });
    }

    return {
      isShow,
      registerActive,
      email,
      password,
      login,
    };
  }
};
</script>


<style>
#dialog {
  color: #2c3e50;
  margin-top: 16px;
}
.dialog {
  position: absolute;
  width: auto;
  height: auto;
  top: 50px;
  right: 0px;
  z-index: 10;
  box-shadow: 4px 4px 8px rgba(0, 0, 0, 0.2);
  background-color: #fff;
}
.dialog__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #f0f0f0;
  padding: 8px 16px;
}
.dialog__close {
  cursor: pointer;
}
.login {
  text-align: center;
  color: #2C3E50;
  margin-top: 60px;
}

p {
   line-height: 1rem;
}

.card {
   padding: 20px;
}

.form-group {
   input {
      margin-bottom: 20px;
   }
}

.login-page {
   align-items: center;
   display: flex;
   height: 100vh;
 
   h1 {
      margin-bottom: 1.5rem;
   }
}
</style>