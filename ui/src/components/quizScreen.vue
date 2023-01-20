<template>
    <div class="modal" id="answerCheckModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalCenterTitle" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="exampleModalCenterTitle">問題</h5>
                </div>
                
                <div class="modal-body">

                    <div class="row">

                        <div class="col-4">
                            <div class="text-center">
                                <img src="../assets/correct.png" class="answer-img" alt="" v-if="isCorrect">
                                <img src="../assets/miss.png" class="answer-img" alt="" v-if="!isCorrect">
                            </div>
                        </div>

                        <div class="col-8">

                            <div v-if="isCorrect">
                                <h3 class="answer-title">正解！</h3>
                                <p>正解は{{questions.correctAnswer}}！</p>
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
    </div>
</template>

<script>
// import { ref } from 'vue';
// import { computed } from 'vue';
// import { useStore } from 'vuex';

export default{
    name: 'QuizScreen',

    // setup() {
    //     const store = useStore();
    //     const quizzes = computed(() => store.state.quizzes);
    //     let currentQuestionNumber = 0;
    //     const countryID = quizzes.value[currentQuestionNumber].CountryID.toLowerCase();
    //     const flagImgPath = ref();
    //     flagImgPath.value = "https://cdn.jsdelivr.net/gh/hjnilsson/countryflags@latest/svg/" + countryID + ".svg";

    //     return{
    //         quizzes,
    //         currentQuestionNumber,
    //         questions:{ //問題文のデータ
    //             hint: quizzes.value[0].Hint1,       //ヒント
    //             answers: [                          //答えの選択肢
    //                 quizzes.value[currentQuestionNumber].Options[0], quizzes.value[currentQuestionNumber].Options[1], quizzes.value[currentQuestionNumber].Options[2], quizzes.value[currentQuestionNumber].Options[3]
    //             ],
    //             correctAnswer: quizzes.value[currentQuestionNumber].CountryName,
    //             description: ""                     //答えたあとに表示する説明など
    //         },

    //         //必要となる変数を定義する
    //         currentQuestionData: [],    //現在の問題データ
    //         correctAnswerNum: 0,        //正解した数
    //         isCorrect: true,            //正解しているかどうか
    //     };
    // }
};

</script>

<style>

.answer-btn:hover{
    color: white;
}

.answer-btn:active {
    -webkit-transform: translateY(2px);
    transform: translateY(2px);
    box-shadow: 0 0 1px rgba(0, 0, 0, 0.15);
    border-bottom: none;
}

/* ボタンの背景色 */
.answer-btn-bg1{
    background: #19c81c;
}

.answer-btn-bg2{
    background: #f2c12e;
}

.answer-btn-bg3{
    background: #f2e22e;
}

.questionBox{
    margin: 7;
    padding: 5px;
    font-size: 18px;
    text-align: center;
}


.choiceBox{
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

.answer-img{
    width: 100px;
}

</style>