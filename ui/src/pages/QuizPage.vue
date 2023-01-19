<template>
    <div class="quizapp">
        <h1>This is a quiz page</h1>
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
                    第{{getCurrentQuestionNumber()}}問 / 全10問
                </div>

                <div class="col-md-8 offset-md-2">
                    <div class="questionBox">
                        <span class="box-title"></span>
                        <p>この国旗の国名は？</p>
                        <!-- <p>{{ quizzes }}</p> -->
                        <p>{{ questions[currentQuestionNumber].answers[0] }}</p>

                    </div>
                </div>

                <!-- 回答ボタンを表示する領域 -->
                <div class="col-md-8 offset-md-2 text-center mt-5">
                    <div class="my-4">
                        <a href="#" v-for="answer in questions[currentQuestionNumber].answers" :key="answer" class="answer-btn answer-btn-bg1 mx-4" @click="answerCheck(1)">{{answer}}</a>
                    </div>

                    <p class="text-center">正解の国名をクリックしよう！</p>
                </div>
            </div>
        </div>

        <!-- <div class="modal" id="answerCheckModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalCenterTitle" aria-hidden="true">
            <div class="modal-dialog modal-dialog-centered" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title" id="exampleModalCenterTitle">問題</h5>
                    </div>
                    
                    <div class="modal-body">

                        <div class="row">

                            <div class="col-4">
                                <div class="text-center">
                                    <img src="./images//correct.png" class="answer-img" alt="" v-if="isCorrect">
                                    <img src="./images/miss.png" class="answer-img" alt="" v-if="!isCorrect">
                                </div>
                            </div>

                            <div class="col-8">

                                <div v-if="isCorrect">
                                    <h3 class="answer-title">正解！</h3>
                                    <p>正解は{{questions[currentQuestionNumber].answers[questions[currentQuestionNumber].answerNum]}}！</p>
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
                        <button type="button" class="btn btn-primary" @click="nextQuestion()">次の問題へ</button>
                    </div>
                </div>
            </div>
        </div> -->

         <!-- すべての問題を回答し終わったときに表示するモーダル -->

         <!-- <div class="modal fade" id="answerEndedModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalCenterTitle" aria-hidden="true">
            <div class="modal-dialog modal-dialog-centered" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        おつかれさまでした！
                    </div>
                    
                    <div class="modal-body">

                        <div class="row">
                            <div class="col-4">
                                <div class="text-center">
                                    <img src="./images/trophy.png" class="answer-img">
                                </div>
                            </div>

                            <div class="col-8">
                                <p>すべてが終了したときの説明</p>
                            </div>

                        </div>
        
                    </div>
        
                    <div class="modal-footer">
                    
                        <a href="index.html" class="btn btn-success">クイズ一覧にもどる</a>
                        <a href="#" class="btn btn-primary" @click="reload()">最初から答える</a>

                    </div>
                </div>
            </div>
        </div> -->

    </div>
</template>

<script>
// import QuizFilter from "@/components/quizFilter.vue";
import { computed } from 'vue'
import { useStore } from 'vuex'

export default{
    name: 'QuizPage',
    components: {
        // QuizFilter,
    },

    methods:{
        getCurrentQuestionNumber(){
            return this.currentQuestionNumber + 1;
        },
        
    },

    setup() {
        const store = useStore();
        const quizzes = computed(() => store.state.quizzes)

        return{
            quizzes,
            selected_question: null,
            questions:{
                0:{ //問題文のデータ
                    hint: quizzes.value,                //ヒント
                    answers: [                          //答えの選択肢
                        quizzes.value[0].Hiragana, 
                    ],
                    answerNum: 1,                       //正解の問題番号（0から始まる）
                    description: ""                     //答えたあとに表示する説明など
                },
            },

            //必要となる変数を定義する
            currentQuestionNumber: 0,   //現在の問題番号
            currentQuestionData: [],    //現在の問題データ
            correctAnswerNum: 0,        //正解した数
            isCorrect: true,            //正解しているかどうか
        };
    }
};

</script>