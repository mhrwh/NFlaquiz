<template>
 <div id="dialog">
    
    <button  @click="dialogToggle"
      >dialog</button
    >
    <transition>
      <div class="dialog"  v-if="isShow" > 
        <div class="container">
          <div class="row">
              <div class="col-lg-4 col-md-6 col-sm-8 mx-auto">
                <div v-if="!registerActive" class="card login" v-bind:class="{ error: emptyFields }">
                    <h1>Sign In</h1>
                    <form class="form-group">
                      <input v-model="emailLogin" type="email" class="form-control" placeholder="Email" required>
                      <input v-model="passwordLogin" type="password" class="form-control" placeholder="Password" required>
                      <input type="submit" class="btn btn-primary" @click="doLogin">
                      <p>Don't have an account? <a href="#" @click="registerActive = !registerActive, emptyFields = false">Sign up here</a>
                      </p>
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
  </div>
<!-- <div class="text-center">
    <v-dialog
      v-model="dialog"
      width="500"
    >
      <v-card>
        <v-card-title class="headline grey lighten-2">
          Vuetify で v-dialog をつかってみる
        </v-card-title>
        <v-card-text style="margin: 1rem auto;">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="primary"
            text
            @click="dialog = false"
          >
          閉じる
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div> -->
  
  <v-btn elevation="5" @click="currentComponent = 'CorrectAnswersRateMap'">正答率</v-btn>
  <v-btn elevation="5" @click="currentComponent = 'BookMarkMap'">ブックマーク</v-btn>
  <CorrectAnswersRateMap v-if="currentComponent === 'CorrectAnswersRateMap'"></CorrectAnswersRateMap>
  <BookMarkMap v-if="currentComponent === 'BookMarkMap'"></BookMarkMap>
</template>

<script>
import CorrectAnswersRateMap from '@/components/correctAnswersRateMap.vue'
import BookMarkMap from '@/components/bookMarkMap.vue'
export default {
  name: 'TopPage',
  components: {
    CorrectAnswersRateMap,
    BookMarkMap
  },
  data() {
    return {
      currentComponent: "BookMarkMap",
      isShow: false,
    };
  },
  methods: {
    dialogToggle: function() {
      this.isShow = !this.isShow;
    }
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
  width: 420px;
  box-shadow: 4px 4px 8px rgba(0, 0, 0, 0.2);
  background-color: #fff;
  z-index: 99999;
  border-radius: 2px;
  border: 1px solid #f0f0f0;
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