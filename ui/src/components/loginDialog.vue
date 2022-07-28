<template>
    
    <button @click="isShow = ! isShow">ログイン</button>
    <transition>
      <div class="dialog"  v-if="isShow" > 
        <div class="container">
          <div class="row">
              <div class="col" style="padding: 0;">
                <div v-if="!registerActive" class="card login" v-bind:class="{ error: emptyFields }">
                    <h1>Sign In</h1>
                    <form class="form-group">
                      <input v-model="email" type="email"  placeholder="Email" required>
                      <input v-model="password" type="password" class="form-control" placeholder="Password" required>
                      <input type="submit" class="btn btn-primary" @click="login">
                      <p>{{email}}{{password}}</p>
                      <a href="#" @click="registerActive = !registerActive, emptyFields = false">アカウント登録</a>
                    </form>
                </div>

                <div v-else class="card register" v-bind:class="{ error: emptyFields }">
                    <h1>Sign Up</h1>
                    <form class="form-group">
                      <input v-model="emailReg" type="email" class="form-control" placeholder="Email" required>
                      <input v-model="passwordReg" type="password" class="form-control" placeholder="Password" required>
                      <input v-model="confirmReg" type="password" class="form-control" placeholder="Confirm Password" required>
                      <input type="submit" class="btn btn-primary" @click="doRegister">
                      <p>Already have an account? <a href="#" @click="registerActive = !registerActive, emptyFields = false">Sign in here</a>
                      </p>
                    </form>
                </div>
              </div>
          </div>

        </div>
      </div>
    </transition>
</template>

<script>
import { ref } from 'vue'
import axios from 'axios';

export default {
  name: 'LoginDialog',

  setup() {
    const isShow = ref(false)
    const registerActive = ref (false)
    const email = ref("")
    const password = ref("")

    const login = async() => {
      const params = new URLSearchParams();
      params.append('email', email.value);
      params.append('password', password.value);
      console.log(params);
      axios.post('http://kubernetes.docker.internal:8888/login',params,)
      .then(res => {
          console.log(res);
          //msg.value = res.data.msg;
          
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
  width: 300px;
  box-shadow: 4px 4px 8px rgba(0, 0, 0, 0.2);
  background-color: #fff;
  cursor: move;
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