<template>
    <transition>
      <div class="quiz-fillter"  v-if="isShow" > 
        <div class="container">
          <div class="row">
              <div class="col" style="padding: 0;">
                <button type="button" class="float-right close" aria-label="Close" @click="isShow = ! isShow">
                  <span aria-hidden="true">&times;</span>
                </button>
                <h3>国旗クイズ</h3>
                
              </div>
          </div>
          <div class="row">
            <div class="col" style="padding: 0;">
                <input type="checkbox" class="btn-check" id="btn-check-2-outlined" checked autocomplete="off">
                <label class="btn btn-outline-secondary" for="btn-check-2-outlined">Checked</label><br>
              </div>
          </div>

        </div>
      </div>
    </transition>
</template>

<script>
import { ref } from 'vue'


export default {
  name: 'QuizFillter',

  setup() {
    const isShow = ref(true)

    return {
      isShow,
    };
  }
};
</script>


<style>

.quiz-fillter {
  position: absolute;
  width: 300px;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  -webkit-transform: translate(-50%, -50%);
  -ms-transform: translate(-50%, -50%);
  z-index: 10;
  box-shadow: 4px 4px 8px rgba(0, 0, 0, 0.2);
  background-color: #fff;
  cursor: move;
}

</style>