<template>
  <div class="modal fade" id="quizFilterModal" tabindex="-1" aria-labelledby="quizFilterModalTitle" aria-hidden="false">
    <div class="modal-dialog modal-lg modal-dialog-centered">
      <div class="modal-content quiz-filter">
        <div class="modal-header p-0 border-bottom-0">
          <h2 class="modal-title" id="quizFilterModalTitle">国旗クイズ</h2>
          <button type="button" class="close modal-close" data-dismiss="modal" aria-label="Close">
            <i class="bi bi-x-square" />
          </button>
        </div>
        <div class="modal-body p-0">
          <div class="container pl-32px">
            <div class="row py-8px">
              <div class="col">
                <h3>絞り込み条件</h3>
              </div>
            </div>
            <div class="row filter-title">
              <div class="col">
                <p class="font-20px list-20px">地域</p>
              </div>
            </div>
            <div class="row pl-28px">
              <div class="col mx-n5px">
                <div class="btn_wrap">
                  <input id="btn_asia" type="checkbox"  value="Asia" v-model="areas">
                  <label class="font-20px" for="btn_asia">アジア</label>
                </div>
                <div class="btn_wrap">
                  <input id="btn_oceania" type="checkbox"  value="Oceania" v-model="areas">
                  <label class="font-20px" for="btn_oceania">オセアニア</label>
                </div>
                <div class="btn_wrap">
                  <input id="btn_northamerica" type="checkbox"  value="NorthAmerica" v-model="areas">
                  <label class="font-20px" for="btn_northamerica">北米</label>
                </div>
                <div class="btn_wrap">
                  <input id="btn_latinamericaandcaribbean" type="checkbox"  value="LatinAmericaandCaribbean" v-model="areas">
                  <label class="font-20px" for="btn_latinamericaandcaribbean">中南米</label>
                </div>
                <div class="btn_wrap">
                  <input id="btn_europe" type="checkbox"  value="Europe" v-model="areas">
                  <label class="font-20px" for="btn_europe">ヨーロッパ</label>
                </div>
                <div class="btn_wrap">
                  <input id="btn_middleeast" type="checkbox"  value="Middleeast" v-model="areas">
                  <label class="font-20px" for="btn_middleeast">中東</label>
                </div>
                <div class="btn_wrap">
                  <input id="btn_africa" type="checkbox"  value="Africa" v-model="areas">
                  <label class="font-20px" for="btn_africa">アフリカ</label>
                </div>
              </div> 
            </div>
            <div class="row filter-title">
              <div class="col">
                <li class="font-20px list-20px">色</li>
              </div>
            </div>
            <div class="row pl-28px">
              <div class="col mx-n12px">
                <div class="btn_color_wrap">
                  <input id="btn_red" type="checkbox" value="red" v-model="colors">
                  <label class="red" for="btn_red">
                    <i :class="colors.includes('red') ? 'bi bi-palette-fill' : 'bi bi-palette'" />
                  </label>
                </div>
                <div class="btn_color_wrap">
                  <input id="btn_orange" type="checkbox" value="orange" v-model="colors">
                  <label class="orange" for="btn_orange">
                    <i :class="colors.includes('orange') ? 'bi bi-palette-fill' : 'bi bi-palette'" />
                  </label>
                </div>
                <div class="btn_color_wrap">
                  <input id="btn_yellow" type="checkbox" value="yellow" v-model="colors">
                  <label class="yellow" for="btn_yellow">
                    <i :class="colors.includes('yellow') ? 'bi bi-palette-fill' : 'bi bi-palette'" />
                  </label>
                </div>
                <div class="btn_color_wrap">
                  <input id="btn_green" type="checkbox" value="green" v-model="colors">
                  <label class="green" for="btn_green">
                    <i :class="colors.includes('green') ? 'bi bi-palette-fill' : 'bi bi-palette'" />
                  </label>
                </div>
                <div class="btn_color_wrap">
                  <input id="btn_blue" type="checkbox" value="blue" v-model="colors">
                  <label class="blue" for="btn_blue">
                    <i :class="colors.includes('blue') ? 'bi bi-palette-fill' : 'bi bi-palette'" />
                  </label>
                </div>
                <div class="btn_color_wrap">
                  <input id="btn_white" type="checkbox" value="white" v-model="colors">
                  <label class="white" for="btn_white">
                    <i :class="colors.includes('white') ? 'bi bi-palette-fill' : 'bi bi-palette'" />
                  </label>
                </div>
                <div class="btn_color_wrap">
                  <input id="btn_black" type="checkbox" value="black" v-model="colors">
                  <label class="black" for="btn_black">
                    <i :class="colors.includes('black') ? 'bi bi-palette-fill' : 'bi bi-palette'" />
                  </label>
                </div>
              </div>
            </div>
            <div class="row filter-title" v-if="auth">
              <div class="col">
                <div class="form-check">
                  <input class="form-check-input" type="checkbox" id="bookmark" v-model="bookmark">
                  <label class="form-check-label font-20px" for="bookmark">
                    ブックマークで絞り込む
                  </label>
                </div>
              </div>
            </div>
            <div class="row filter-title" :class="errmsg ? 'd-block' : 'd-none'">
              <div class="col">
                <p class="font-20px text-danger d-flex mb-0 mt-3">
                  <i class="bi bi-exclamation-circle me-3px" />
                  {{ errmsg }}
                </p>
              </div>
            </div>
          </div>
        </div>
        <div class="modal-footer p-0 border-top-0">
          <button type="button" class="btn btn-danger btn-quiz font-20px" @click="reset" :disabled="!(areas.length || colors.length || bookmark)">リセット</button>
          <button type="button" class="btn btn-primary btn-quiz font-20px" @click="quiz">スタート</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import axios from 'axios';
import { computed } from 'vue';
import { useStore } from 'vuex';
// import { useRouter } from 'vue-router';

export default {
  name: 'QuizFilter',

  setup() {
    const bookmark = ref(false);
    let areas = ref([]);
    let colors = ref([]);
    let errmsg = ref("");
    let url = '';
    const store = useStore();
    // const router = useRouter();
    const auth = computed(() => store.state.auth);

    const reset = () =>{
      areas.value.splice(0);
      colors.value.splice(0);
      bookmark.value = false;
    }
    
    const quiz = async() => {
      url = 'http://localhost:8888/quiz/select?'
      if (colors.value.length){
        for(let i=0; i<colors.value.length; i++){
          url = url + 'colors[]=' + colors.value[i] + `&`;
        } 
      }
      if (areas.value.length){
        for(let i=0; i<areas.value.length; i++){
          url = url + 'areas[]=' + areas.value[i] + `&`;
        }
      }
      if (bookmark.value){
        url += 'bookmark=1';
      }
      axios.get(url)
      .then(res => {
        if (res.data.quizzes.length) {
          store.dispatch("setQuizzes", res.data.quizzes);
          document.querySelector('body').classList.remove('modal-open');
          document.querySelector('.modal-backdrop').remove();
          document.querySelector('#quizFilterModal').style.display = 'none';
          // router.push({name: 'QuizPage',});
        }else {
          errmsg.value = "該当する国が存在しません";
        }
      })
      .catch(error => {
        errmsg.value = error.response.data.err_msg;
      });
    }

    return {
      bookmark,
      areas,
      colors,
      quiz,
      url,
      auth,
      reset,
      errmsg,
    };
  },
};
</script>

<style>
.close i {
  font-size: 36px;
}
.pl-28px {
  padding-left:28px;
}
.py-8px {
  padding:8px 0;
}
.mx-n5px{
  margin: 0 -5px
}
.mx-n12px{
  margin: 0 -12px
}
.list-20px {
  list-style: none;
  margin: 0;
  padding: 0;
}
.list-20px::before {
  content: "";
  width: 10px;
  height: 10px;
  display: inline-block;
  background-color: #494949;
  border-radius: 50%;
  position: relative;
  top: -1px;
  margin-right: 7px;
}
.filter-title {
  padding:4px 0 4px 10px;
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
  padding: 12px 20px;
  margin: 10px 5px;
  border-radius: 30px;
  background: #D6C3E4;
  display: inline-block;
  cursor: pointer;
  transition: .5s;
  color: #FFF;
}
.btn_wrap input:checked + label{
  background: #806EB4;
  color: #FFF;
}
.btn_wrap input:focus + label{
  box-shadow: 0 0 4px #806EB4;
}

.btn_color_wrap{
  display:inline-block;
  position: relative;
  margin: 3px 12px;
}
.btn_color_wrap input{
  opacity: 0;
  position: absolute;
  left: 0;
}
.btn_color_wrap label{
  display: none;
  cursor: pointer;
  display: inline-block;
  position: relative;
  margin-bottom: 0;
}
.red i {
  color: #FFCF32;
  font-size: 36px;
}
.orange i {
  color: #f00;
  font-size: 36px;
}
.yellow i {
  color: #ff0;
  font-size: 36px;
}
.green i {
  color: #008000;
  font-size: 36px;
}
.blue i {
  color: #00f;
  font-size: 36px;
}
.white i {
  color: #fff;
  font-size: 36px;
}
.black i {
  color: #000;
  font-size: 36px;
}
</style>