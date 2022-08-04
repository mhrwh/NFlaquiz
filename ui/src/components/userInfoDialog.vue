<template>
    
    <button class="bi bi-person-fill" type="button" @click="isShow = ! isShow"  style="position: absolute; right: 0px; z-index: 10"></button>
    <transition>
      <div class="dialog"  v-if="isShow" > 
        <div class="container">
          <div class="row">
              <div class="col" style="padding: 0;">
                <div class="form-group">
                    <button class="btn btn-primary" type="button" @click="logout">ログアウト</button>
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
  name: 'UserInfoDialog',

  setup() {
    const store = useStore()
    const isShow = ref(false)

    const logout = async() => {
      axios.get('http://localhost:8888/logout',)
      .then(res => {
          console.log(res);
          //msg.value = res.data.msg;
          store.dispatch("setAuth", false);
          location.reload();
      })
      .catch(error => {
          console.log(error.response.data.err_msg);
          //msg.value = error.response.data.err_msg;
      });
    }

    return {
      isShow,
      logout,
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