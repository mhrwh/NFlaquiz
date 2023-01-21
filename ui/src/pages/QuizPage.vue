<template>
  <div class="quizapp">
    <!-- <QuizScreen /> -->
    <nav class="navbar navbar-expand-lg navbar-dark bg-main-green">
      <a class="navbar-brand">
        <img
          src="../assets/NFlaquizlogo.png"
          width="30"
          height="30"
          class="d-inline-block align-top mr-2"
          alt=""
        />
        NFlaquiz
      </a>
    </nav>
    <div class="container">
      <div class="row my-3">
        <div class="col-12">
          <h1 class="text-center main-title">国旗クイズ</h1>
          <p class="text-center">国名を当てよう！</p>
        </div>
      </div>

      <div class="row">
        <!-- 問題数を表示する領域 -->
        <div class="col-md-8 offset-md-2">
          第{{ currentQuizNumber }}問 / 全{{ totalQuizNumber }}問
        </div>

        <div class="col-md-8 offset-md-2">
          <div class="questionBox">
            <span class="box-title"></span>
            <p>この国旗の国名は？</p>
            <img :src="flagImgPath" class="flagImg" />
          </div>
        </div>

        <!-- 回答ボタンを表示する領域 -->
        <div class="col-md-8 offset-md-2 text-center mt-5">
          <div class="my-4">
            <button
              v-for="(option, index) in options"
              :key="index"
              class="answer-btn answer-btn-bg1 mx-4"
              data-toggle="modal"
              data-target="#answerCheckModal"
              @click="judgeAnswer(option)"
            >
              {{ option }}
            </button>
          </div>

          <p class="text-center">正解の国名をクリックしよう！</p>
        </div>
      </div>
    </div>

    <div
      class="modal"
      id="answerCheckModal"
      tabindex="-1"
      role="dialog"
      aria-labelledby="exampleModalCenterTitle"
      aria-hidden="true"
    >
      <div class="answerModal-dialog modal-dialog-centered" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalCenterTitle">問題</h5>
          </div>

          <div class="modal-body">
            <div class="row">
              <div class="col-4">
                <div class="text-center">
                  <img
                    v-if="isCorrect"
                    src="../assets/correct.png"
                    class="answer-img"
                    alt=""
                  />
                  <img
                    v-else
                    src="../assets/miss.png"
                    class="answer-img"
                    alt=""
                  />
                </div>
              </div>

              <div class="col-8">
                <div v-if="isCorrect">
                  <h3 class="answer-title">正解！</h3>
                  <p>正解は{{ currentCorrectAnswer }}！</p>
                  <p></p>
                </div>

                <div v-else>
                  <h3 class="answer-title">残念...!</h3>
                  <p>不正解だったときの説明</p>
                  <p></p>
                </div>
              </div>
            </div>
          </div>

          <div class="modal-footer">
            <button type="button" class="btn btn-primary" @click="toNextQuiz()">
              次の問題へ
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- すべての問題を回答し終わったときに表示するモーダル -->

    <div
      class="modal fade"
      id="answerEndedModal"
      tabindex="-1"
      role="dialog"
      aria-labelledby="exampleModalCenterTitle"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered" role="document">
        <div class="modal-content">
          <div class="modal-header">おつかれさまでした！</div>

          <div class="modal-body">
            <div class="row">
              <div class="col-4">
                <div class="text-center">
                  <img src="../assets/trophy.png" class="answer-img" />
                </div>
              </div>

              <div class="col-8">
                <p>すべてが終了したときの説明</p>
              </div>
            </div>
          </div>

          <div class="modal-footer">
            <a href="index.html" class="btn btn-success">クイズ一覧にもどる</a>
            <a href="#" class="btn btn-primary" @click="reload()"
              >最初から答える</a
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import QuizScreen from "@/components/quizScreen.vue";
import { onMounted, ref } from "vue";
import { computed } from "vue";
import { useStore } from "vuex";
import axios from "axios";

export default {
  name: "QuizPage",
  components: {
    // QuizScreen,
  },
  setup() {
    const store = useStore();
    const quizzes = computed(() => store.state.quizzes);

    const totalQuizNumber = ref(quizzes.value.length);
    const currentQuizNumber = ref(0);
    const currentCorrectAnswer = ref();
    const hints = ref();
    const options = ref();
    const flagImgPath = ref();
    const results = ref([]);
    const bookmarks = ref(Array(totalQuizNumber.value).fill(0));

    const isCorrect = ref();

    onMounted(() => {
      updateQuiz();
    });

    // ※このコメントは仕様が理解出来たら削除してOK
    // クイズ回答後, 次のクイズに更新する処理(問題番号のカウントアップなどもここで実装)
    // 次の問題に行く処理の中で呼び出される想定
    const updateQuiz = () => {
      currentQuizNumber.value += 1;
      const currentQuiz = quizzes.value[currentQuizNumber.value - 1];
      currentCorrectAnswer.value = currentQuiz["CountryName"];
      hints.value = currentQuiz["Hints"];
      options.value = currentQuiz["Options"];

      const countryID = currentQuiz["CountryID"].toLowerCase();
      flagImgPath.value =
        "https://cdn.jsdelivr.net/gh/hjnilsson/country-flags@latest/svg/" +
        countryID +
        ".svg";
    };

    // ※このコメントは仕様が理解出来たら削除してOK
    // ブックマーク登録の有無を反転させる
    // 例(5問中)
    // 2問目のクイズをブックマーク登録する
    // updateBookmark(1) => bookmarks = [0,1,0,0,0]
    // 3問目のクイズをブックマーク登録する
    // updateBookmark(2) => bookmarks = [0,1,1,0,0]
    // 3問目のクイズのブックマークを解除
    // updateBookmark(2) => bookmarks = [0,1,0,0,0]
    const updateBookmark = (quiz_index) => {
      bookmarks.value[quiz_index] = bookmarks.value[quiz_index] === 0 ? 1 : 0;
    };

    // ※このコメントは仕様が理解出来たら削除してOK
    // APIにデータを送信する
    // ゲーム終了後, "トップへ戻る"/"最初からやり直す"ボタンがクリックされた時に呼び出される想定
    const sendResult = async () => {
      const sendingData = [];
      for (let i = 0; i < totalQuizNumber.value; i++) {
        sendingData.push({
          country_id: quizzes.value[i]["CountryID"],
          answer: results.value[i],
          bookmark: bookmarks.value[i],
        });
      }
      try {
        const url = "http://localhost:8888/result/update";
        await axios.post(url, sendingData);
      } catch (e) {
        console.log(e);
      }
    };

    //答えをチェックする処理
    const judgeAnswer = (option) => {
      isCorrect.value = option === currentCorrectAnswer.value;
      results.value.push(isCorrect.value ? 1 : 0);
      //結果を表示するモーダルを表示する
    };

    //次の問題に行く処理（いじらない）
    const toNextQuiz = () => {
      console.log(currentQuizNumber.value)
      console.log(totalQuizNumber.value)
      if (currentQuizNumber.value < totalQuizNumber.value) {
        updateQuiz();
        // $('#answerCheckModal').modal('hide');    //モーダルを隠す
      } else {
        //最後の問題なので問題番号は増やさない
        // $('#answerCheckModal').modal('hide');     //モーダルを隠す
        //すべての問題が解き終わったので、最終結果モーダルを表示する
        // $('#answerEndedModal').modal({
        //     keyboard: false,
        //     backdrop: "static"
        // });
      }
    };

    return {
      currentQuizNumber,
      currentCorrectAnswer,
      hints,
      options,
      flagImgPath,
      totalQuizNumber,
      isCorrect,
      judgeAnswer,
      toNextQuiz,
    };
  },
};
</script>

<style>
#answerCheckModal {
  position: fixed;
  background-color: rgba(0, 0, 0, 0.5);
}

.answerModal-dialog {
  position: absolute;
  width: 500px;
  left: 50%;
  transform: translateX(-50%);
}

.flagImg {
  height: 300px;
  width: auto;
  border: 1px solid black;
}

.bg-main-green {
  background: #3ebebec8;
}

/*== 個別のUIデザイン ==*/
.main-title {
  font-size: 28px;
  font-weight: bold;
}

/*== 回答ボタンのデザイン ==*/
.answer-btn {
  /* font-size: 20px; */
  display: inline-block;
  text-decoration: none;
  background: #81ff89;
  color: #fff;
  width: 130px;
  height: 80px;
  line-height: 80px;
  border-radius: 50%;
  text-align: center;
  font-weight: bold;
  overflow: hidden;
  box-shadow: 0px 2px 2px rgba(0, 0, 0, 0);
  border: solid 3px #668d1f00;
  transition: 0.4s;
}

.my-4 {
  display: inline;
}

.answer-btn:hover {
  color: white;
}

.answer-btn:active {
  -webkit-transform: translateY(2px);
  transform: translateY(2px);
  box-shadow: 0 0 1px rgba(0, 0, 0, 0.15);
  border-bottom: none;
}

/* ボタンの背景色 */
.answer-btn-bg1 {
  background: #19c81c;
}

.answer-btn-bg2 {
  background: #f2c12e;
}

.answer-btn-bg3 {
  background: #f2e22e;
}

.questionBox {
  margin: 7;
  padding: 5px;
  font-size: 18px;
  text-align: center;
}

.choiceBox {
  padding: 0.5em 1em;
  margin: 2em 0;
  color: #5d627b;
  background: white;
  border-top: solid 5px #5d627b;
  box-shadow: 0 3px 5px rgba(0, 0, 0, 0.22);
}

.choiceBox p {
  margin: 0;
  padding: 0;
}

.answer-img {
  width: 100px;
}
</style>