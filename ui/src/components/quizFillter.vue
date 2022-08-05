<template>
    <transition>
      <div class="quiz-fillter" v-if="quizDialog" > 
        <div class="container">
          <div class="row">
              <div class="col" >
                <button type="button" class="float-right close" aria-label="Close" @click="close">
                  <span aria-hidden="true">&times;</span>
                </button>
                <h3>国旗クイズ</h3>
                
              </div>
          </div>
          <div class="row">
            <div class="col">
              地域
            </div>
            <div class="col" style="padding: 0;">
              <div class="btn_wrap">
                <input id="btn_asia" type="checkbox"  value="Asia" v-model="areas">
                <label for="btn_asia">アジア</label>
              </div>
              <div class="btn_wrap">
                <input id="btn_oceania" type="checkbox"  value="Oceania" v-model="areas">
                <label for="btn_oceania">オセアニア</label>
              </div>
              <div class="btn_wrap">
                <input id="btn_northamerica" type="checkbox"  value="NorthAmerica" v-model="areas">
                <label for="btn_northamerica">北米</label>
              </div>
              <div class="btn_wrap">
                <input id="btn_latinamericaandcaribbean" type="checkbox"  value="LatinAmericaandCaribbean" v-model="areas">
                <label for="btn_latinamericaandcaribbean">中南米</label>
              </div>
              <div class="btn_wrap">
                <input id="btn_europe" type="checkbox"  value="Europe" v-model="areas">
                <label for="btn_europe">ヨーロッパ</label>
              </div>
              <div class="btn_wrap">
                <input id="btn_middleeast" type="checkbox"  value="Middleeast" v-model="areas">
                <label for="btn_middleeast">中東</label>
              </div>
              <div class="btn_wrap">
                <input id="btn_africa" type="checkbox"  value="Africa" v-model="areas">
                <label for="btn_africa">アフリカ</label>
              </div>
            </div>
            
          </div>
          <div class="row">
            <div class="col">
              色
            </div>
            <div class="col" style="padding: 0;">
              <div class="btn_color_wrap" id="red">
                <input id="btn_red" type="checkbox"  value="red" v-model="colors">
                <label for="btn_red"></label>
              </div>
              <div class="btn_color_wrap" id="orange">
                <input id="btn_orange" type="checkbox"  value="orange" v-model="colors">
                <label for="btn_orange"></label>
              </div>
              <div class="btn_color_wrap" id="yellow">
                <input id="btn_yellow" type="checkbox"  value="yellow" v-model="colors">
                <label for="btn_yellow"></label>
              </div>
              <div class="btn_color_wrap" id="green">
                <input id="btn_green" type="checkbox"  value="green" v-model="colors">
                <label for="btn_green"></label>
              </div>
              <div class="btn_color_wrap" id="blue">
                <input id="btn_blue" type="checkbox"  value="blue" v-model="colors">
                <label for="btn_blue"></label>
              </div>
              <div class="btn_color_wrap" id="white">
                <input id="btn_white" type="checkbox"  value="white" v-model="colors">
                <label for="btn_white"></label>
              </div>
              <div class="btn_color_wrap" id="black">
                <input id="btn_black" type="checkbox"  value="black" v-model="colors">
                <label for="btn_black"></label>
              </div>
            </div>
          </div>
          <div class="row" v-if="auth">
            <div class="col">
              <div class="form-check">
                <input class="form-check-input" type="checkbox" id="bookmark" v-model="bookmark">
                <label class="form-check-label" for="bookmark">
                  ブックマークで絞り込む
                </label>
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col">
              <input class="btn btn-primary" type="submit" value="スタート" @click="quiz">
            </div>
          </div>
        </div>
      </div>
    </transition>
</template>

<script>
import { ref } from 'vue'
import axios from 'axios'
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: 'QuizFillter',

  setup() {
    const quizDialog = computed(() => store.state.quizDialog);
    const bookmark = ref(false);
    const areas = ref([]);
    const colors = ref([]);
    let url = '';
    const store = useStore();
    const auth = computed(() => store.state.auth);

    const close = () =>{
      store.dispatch("setQuizDialog", false);
    }
    const quiz = async() => {
      url = 'http://localhost:8888/quiz/select?'
      if (!colors.value[0]){
        console.log('a')
        url += 'colors[]=red&colors[]=orange&colors[]=yellow&colors[]=green&colors[]=blue&colors[]=white&colors[]=black&';
      } else{
        for(let i=0; i<colors.value.length; i++){
          url = url + 'colors[]=' + colors.value[i] + `&`;
        }
      }
      if (!areas.value[0]){
        url += 'areas[]=Asia&areas[]=Oceania&areas[]=NorthAmerica&areas[]=LatinAmericaandCaribbean&areas[]=Europe&areas[]=MiddleEast&areas[]=Africa&';
      } else{
        for(let i=0; i<areas.value.length; i++){
          url = url + 'areas[]=' + areas.value[i] + `&`;
        }
      }
      if (bookmark.value){
        url += 'bookmark=1'
      } else{
        url += 'bookmark=0'
      }
      console.log(url)
      axios.get(url)
      .then(res => {
          console.log(res.data.quizzes);
          //msg.value = res.data.msg;
      })
      .catch(error => {
          console.log(error.response.data.err_msg);
          //msg.value = error.response.data.err_msg;
      });
    }

    return {
      bookmark,
      areas,
      colors,
      quiz,
      url,
      auth,
      quizDialog,
      close,
    };
  }
};
</script>


<style>

.quiz-fillter {
  position: absolute;
  width:auto;
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
.btn_wrap{
  display:inline-block;
  position: relative;
}
.btn_wrap input{
  opacity: 0;
  position: absolute;
  left: 0;
}
.btn_wrap label{
  padding: 10px 20px;
  border-radius: 20px;
  background: #eee;
  display: inline-block;
  cursor: pointer;
  transition: .5s;
}

.btn_wrap input:checked + label{
  background: #af975e;
  color: #FFF;
}
.btn_wrap input:focus + label{
  box-shadow: 0 0 4px #af975e;
}

.btn_color_wrap{
  display:inline-block;
  position: relative;
}
.btn_color_wrap input{
  opacity: 0;
  position: absolute;
  left: 0;
}
.btn_color_wrap label{
  padding: 20px 20px;
  border-radius: 20px;
  background: #eee;
  display: inline-block;
  cursor: pointer;
  transition: .5s;
}
#red label{
  border: 2px solid #f00;
}
#red input:checked + label{
  background: #f00;
  border: 1px solid #fff;
}
#red input:focus + label{
  box-shadow: 0 0 4px #f00;
}
#orange label{
  border: 2px solid #ffa500;
}
#orange input:checked + label{
  background: #ffa500;
  border: 1px solid #fff;
}
#orange input:focus + label{
  box-shadow: 0 0 4px #ffa500;
}
#yellow label{
  border: 2px solid #ff0;
}
#yellow input:checked + label{
  background: #ff0;
  border: 1px solid #fff;
}
#yellow input:focus + label{
  box-shadow: 0 0 4px #ff0;
}
#green label{
  border: 2px solid #008000;
}
#green input:checked + label{
  background: #008000;
  border: 1px solid #fff;
}
#green input:focus + label{
  box-shadow: 0 0 4px #008000;
}
#blue label{
  border: 2px solid #00f;
}
#blue input:checked + label{
  background: #00f;
  border: 1px solid #fff;
}
#blue input:focus + label{
  box-shadow: 0 0 4px #00f;
}
#white label{
  border: 2px solid #fff;
}
#white input:checked + label{
  background: #fff;
  border: 1px solid #fff;
}
#white input:focus + label{
  box-shadow: 0 0 4px #fff;
}
#black label{
  border: 2px solid #000;
}
#black input:checked + label{
  background: #000;
  border: 1px solid #fff;
}
#black input:focus + label{
  box-shadow: 0 0 4px #000;
}

</style>