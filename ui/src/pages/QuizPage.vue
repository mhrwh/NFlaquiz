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
          第{{ getcurrentQuizNumber() }}問 / 全10問
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
            <!-- <button v-for="answer in questions.answers" :key="answer" class="answer-btn answer-btn-bg1 mx-4" data-toggle="modal" data-target="#answerCheckModal" @click="answerCheck(answer)">{{answer}}</button> -->
            <button
              class="answer-btn answer-btn-bg1 mx-4"
              data-toggle="modal"
              data-target="#answerCheckModal"
              @click="answerCheck(questions.answers[0])"
            >
              {{ questions.answers[0] }}
            </button>
            <button
              class="answer-btn answer-btn-bg1 mx-4"
              data-toggle="modal"
              data-target="#answerCheckModal"
              @click="answerCheck(questions.answers[1])"
            >
              {{ questions.answers[1] }}
            </button>
            <button
              class="answer-btn answer-btn-bg1 mx-4"
              data-toggle="modal"
              data-target="#answerCheckModal"
              @click="answerCheck(questions.answers[2])"
            >
              {{ questions.answers[2] }}
            </button>
            <button
              class="answer-btn answer-btn-bg1 mx-4"
              data-toggle="modal"
              data-target="#answerCheckModal"
              @click="answerCheck(questions.answers[3])"
            >
              {{ questions.answers[3] }}
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
                    src="../assets/correct.png"
                    class="answer-img"
                    alt=""
                    v-if="isCorrect"
                  />
                  <img
                    src="../assets/miss.png"
                    class="answer-img"
                    alt=""
                    v-if="!isCorrect"
                  />
                </div>
              </div>

              <div class="col-8">
                <div v-if="isCorrect">
                  <h3 class="answer-title">正解！</h3>
                  <p>正解は{{ questions.correctAnswer }}！</p>
                  <p></p>
                </div>

                <div v-if="!isCorrect">
                  <h3 class="answer-title">残念...!</h3>
                  <p>不正解だったときの説明</p>
                  <p></p>
                </div>
              </div>
            </div>
          </div>

          <div class="modal-footer">
            <button
              type="button"
              class="btn btn-primary"
              @click="nextQuestion()"
            >
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
import { ref } from "vue";
import { computed } from "vue";
import { useStore } from "vuex";

export default {
  name: "QuizPage",
  components: {
    // QuizScreen,
  },

  methods: {
    getcurrentQuizNumber() {
      return this.currentQuizNumber + 1;
    },

    //答えをチェックする処理
    answerCheck(userChoiceAnswer) {
      if (userChoiceAnswer == this.questions.correctAnswer) {
        //正解！
        this.isCorrect = true;
        this.correctAnswerNum += 1;
      } else {
        this.isCorrect = false;
      }
      //結果を表示するモーダルを表示する
    },

    //次の問題に行く処理（いじらない）
    nextQuestion() {
      if (this.currentQuizNumber < Object.keys(this.questions).length - 1) {
        this.currentQuizNumber += 1; //最後の問題ではないので問題番号を1ずつ増やす
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
    },
  },

  setup() {
    const store = useStore();
    const quizzes = computed(() => store.state.quizzes);

    let currentQuizNumber = 0;
    const countryID =
      quizzes.value[currentQuizNumber]["CountryID"].toLowerCase();
    const flagImgPath = ref();
    flagImgPath.value =
      "https://cdn.jsdelivr.net/gh/hjnilsson/country-flags@latest/svg/" +
      countryID +
      ".svg";

    return {
      quizzes,
      currentQuizNumber,
      countryID,
      flagImgPath,
      questions: {
        //問題文のデータ
        hint: quizzes.value[0].Hint1, //ヒント
        answers: [
          //答えの選択肢
          quizzes.value[currentQuizNumber].Options[0],
          quizzes.value[currentQuizNumber].Options[1],
          quizzes.value[currentQuizNumber].Options[2],
          quizzes.value[currentQuizNumber].Options[3],
        ],
        correctAnswer: quizzes.value[currentQuizNumber].CountryName,
        description: "", //答えたあとに表示する説明など
      },

      //必要となる変数を定義する
      currentQuestionData: [], //現在の問題データ
      correctAnswerNum: 0, //正解した数
      isCorrect: true, //正解しているかどうか
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