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
              <div class="btn-group">
                <input type="checkbox" class="btn-check" id="btn-check-2-outlined" checked autocomplete="off">
                <label class="btn btn-outline-secondary" for="btn-check-2-outlined">Checked</label><br>
              </div>
              <b-form-checkbox v-model="checked1" name="check-button" button>
                Button Checkbox
              </b-form-checkbox>
              <div class="btn2_wrap">
                  <input value="1" id="btn_demo2-1" type="checkbox">
                  <label for="btn_demo2-1">アジア</label>
              </div>
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
.btn2_wrap{
    display:inline-block;
    position: relative;
}
.btn2_wrap input{
    opacity: 0;
    position: absolute;
    left: 0;
}
.btn2_wrap label{
    padding: 10px 20px;
    border-radius: 20px;
    background: #eee;
    display: inline-block;
    cursor: pointer;
    transition: .5s;
}

.btn2_wrap input:checked + label{
    background: #af975e;
    color: #FFF;
}
.btn2_wrap input:focus + label{
    box-shadow: 0 0 4px #af975e;
}
</style>