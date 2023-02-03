<template>
  <div class="quizapp">
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
          <button class="hint-btn" data-toggle="modal" data-target="#hintModal">
            ヒント
          </button>
        </div>

        <div class="col-md-8 offset-md-2">
          <div class="questionBox">
            <span class="box-title"></span>
            <p>この国旗の国名は？</p>
            <img :src="flagImgPath" class="flagImg" />
          </div>
        </div>

        <!-- 回答ボタンを表示する領域 -->
        <div class="col-md-8 offset-md-2 text-center mt-2">
          <div class="btn-box">
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

          <p class="text-center mt-4">正解の国名をクリックしよう！</p>
        </div>
      </div>
    </div>

    <!-- ヒントを表示するモーダル -->

    <div
      class="modal"
      id="hintModal"
      role="dialog"
      aria-labelledby="exampleModalCenterTitle"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <p>ヒント</p>
            <button class="hint-close" data-dismiss="modal">×</button>
          </div>

          <div class="modal-body">
            <p class="m-4 text-center">{{ hint }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- 答えを表示するモーダル -->

    <div
      class="modal"
      id="answerCheckModal"
      tabindex="-1"
      role="dialog"
      aria-labelledby="exampleModalCenterTitle"
      aria-hidden="true"
      data-backdrop="static"
    >
      <div class="modal-dialog modal-dialog-centered" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalCenterTitle">
              第{{ currentQuizNumber }}問
            </h5>
            <button
              class="btn btn-light btn-circle"
              @click="updateBookmark(currentQuizNumber - 1)"
            >
              <i
                v-if="bookmarks[currentQuizNumber - 1]"
                class="bi bi-bookmark-heart-fill"
              />
              <i v-else class="bi bi-bookmark-heart"></i>
            </button>
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
                  <p>正解は{{ correctAnswer }}！</p>
                  <p></p>
                </div>

                <div v-else>
                  <h3 class="answer-title">残念...!</h3>
                  <p>正解は{{ correctAnswer }}！</p>
                  <p></p>
                </div>
              </div>
            </div>
          </div>

          <div class="modal-footer">
            <button
              v-show="currentQuizNumber < totalQuizNumber"
              type="button"
              class="btn btn-primary"
              data-dismiss="modal"
              @click="toNextQuiz()"
            >
              次の問題へ
            </button>
            <button
              v-if="currentQuizNumber == totalQuizNumber"
              type="button"
              class="btn btn-primary"
              data-toggle="modal"
              data-target="#answerEndedModal"
              @click="toNextQuiz()"
            >
              結果へ
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
      data-backdrop="static"
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
                <p v-if="correctNumber.length == totalQuizNumber">全問正解！</p>
                <p v-else>
                  {{ totalQuizNumber }}問中{{ correctNumber.length }}問正解！
                </p>
                <table>
                  <tr
                    v-for="(result, number) in results"
                    :key="result"
                    class="border-bottom"
                  >
                    <td class="px-2">第{{ number + 1 }}問</td>
                    <td class="px-2">{{ quizzes[number]["CountryName"] }}</td>
                    <td class="px-2">{{ result ? "〇" : " ×" }}</td>
                    <td class="px-2">
                      <button
                        class="btn btn-light btn-circle"
                        @click="updateBookmark(number)"
                      >
                        <i
                          v-if="bookmarks[number]"
                          class="bi bi-bookmark-heart-fill"
                        />
                        <i v-else class="bi bi-bookmark-heart"></i>
                      </button>
                    </td>
                  </tr>
                </table>
              </div>
            </div>
          </div>

          <div class="modal-footer">
            <button class="btn btn-success" @click="toTopPage()">
              トップへ戻る
            </button>
            <button class="btn btn-primary" @click="startOver()">
              最初から答える
            </button>
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
import { useRouter } from "vue-router";
import axios from "axios";

export default {
  name: "QuizPage",
  components: {
    // QuizScreen,
  },
  setup() {
    const store = useStore();
    const router = useRouter();
    const quizzes = computed(() => store.state.quizzes);

    const totalQuizNumber = ref(quizzes.value.length);
    const currentQuizNumber = ref(0);
    const correctAnswer = ref();
    const hint = ref();
    const options = ref();
    const flagImgPath = ref();
    const results = ref([]);
    const bookmarks = ref(Array(totalQuizNumber.value).fill(0));

    const isCorrect = ref();
    const correctNumber = ref([]);

    onMounted(() => {
      updateQuiz();
    });

    const startOver = () => {
      sendResult();
      router.go("/quiz");
    };

    const toTopPage = () => {
      const backdropElems = document.querySelectorAll(".modal-backdrop");
      for (const elem of backdropElems) {
        elem.remove();
      }
      sendResult();
      router.push("/");
    };

    const updateQuiz = () => {
      currentQuizNumber.value += 1;
      const currentQuiz = quizzes.value[currentQuizNumber.value - 1];
      const hints = currentQuiz["Hints"];
      hint.value = hints[Math.floor(Math.random() * hints.length)];
      correctAnswer.value = currentQuiz["CountryName"];
      options.value = currentQuiz["Options"];

      const countryID = currentQuiz["CountryID"].toLowerCase();
      flagImgPath.value =
        "https://cdn.jsdelivr.net/gh/hjnilsson/country-flags@latest/svg/" +
        countryID +
        ".svg";
    };

    const updateBookmark = (quiz_index) => {
      bookmarks.value[quiz_index] = bookmarks.value[quiz_index] === 0 ? 1 : 0;
    };

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
      isCorrect.value = option === correctAnswer.value;
      results.value.push(isCorrect.value ? 1 : 0);
      if (isCorrect.value) {
        correctNumber.value.push(1);
      }
    };

    //次の問題に行く処理（いじらない）
    const toNextQuiz = () => {
      if (currentQuizNumber.value < totalQuizNumber.value) {
        updateQuiz();
      }
    };

    return {
      currentQuizNumber,
      correctAnswer,
      hint,
      options,
      flagImgPath,
      totalQuizNumber,
      isCorrect,
      judgeAnswer,
      toNextQuiz,
      results,
      correctNumber,
      updateBookmark,
      startOver,
      toTopPage,
      quizzes,
      bookmarks,
    };
  },
};
</script>

<style>
.btn-box button:active {
  margin-top: 0.2em;
}

.hint-btn {
  display: block;
  margin: 5px 0 0 auto;
  background: rgb(221, 82, 47);
  border-radius: 10px;
  color: #fff;
  box-shadow: 0 1px 0 #988588;
  border: solid 3px #668d1f00;
  transition: 0.4s;
}

.hint-close {
  border: solid 3px #668d1f00;
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
  width: 200px;
  margin-top: 20px;
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

.btn-box {
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

.questionBox {
  margin: 7;
  padding: 5px;
  font-size: 18px;
  text-align: center;
}

/* .choiceBox {
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
} */

.answer-img {
  width: 100px;
}
</style>
