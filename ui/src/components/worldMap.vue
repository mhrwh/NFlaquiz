<template>
  <div id="svgMap" style="position: relative;">
    <QuizFillter></QuizFillter>
    <div class="btn-toolbar" role="toolbar">
      <div class="btn-group mr-2" role="group">
        <button class="btn btn-light btn-circle m-1" style="border-radius: 50%"><i class="bi bi-person-fill"></i></button>
        <!-- <LoginDialog v-if="!auth"></LoginDialog>
        <UserInfoDialog v-if="auth"></UserInfoDialog> -->
      </div>
      <div class="btn-group mr-2" role="group">
        <button class="btn btn-light btn-circle m-1" style="border-radius: 50%"><i class="bi bi-layout-sidebar-inset-reverse"></i></button>
      </div>
      <div class="btn-group mr-2" role="group">
        <button class="btn btn-light btn-circle m-1" data-toggle="modal" data-target="#exampleModalCenter" style="border-radius: 50%"><i class="bi bi-flag-fill"></i></button>
        <!-- <button type="button" class="bi bi-flag-fill" data-toggle="modal" data-target="#exampleModalCenter">
        </button> -->
      </div>
    </div>
    
    <LoginDialog v-if="!auth"></LoginDialog>
    <UserInfoDialog v-if="auth"></UserInfoDialog>
  </div>
  
  
</template>

<script>
import mapData from "../data/mapData";
import axios from 'axios';
import { computed } from 'vue'
import { useStore } from 'vuex'
import svgMap from 'svgmap';
import 'svgmap/dist/svgMap.min.css';
import LoginDialog from '@/components/loginDialog.vue';
import UserInfoDialog from '@/components/userInfoDialog.vue';
import QuizFillter from "./quizFillter.vue";

export default {
  name: 'WorldMap',

  props: {
    mapMode: String,
  },

  components: {
    LoginDialog,
    UserInfoDialog,
    QuizFillter,
},

  setup (props) {
    const store = useStore();
    const auth = computed(() => store.state.auth);
    let data;
    
    axios.get('http://localhost:8888/map')
      .then((res =>{
        data = res.data;
        console.log(data.map_info);
        let jpName = {};
        for(let i=0; i<data.map_info.length; i++){
          const des = data.map_info[i].description.match(/.{1,25}/g);
          for(let i=0; i<8; i++){
            if(!des[i]) des[i] = '';
          }
          let colorWeight;
          console.log(props.mapMode)
          if (props.mapMode == 'correctAnswersRate'){
            colorWeight = data.map_info[i].weight;
            console.log("a")
          }else if (props.mapMode == 'bookMark'){
            colorWeight = data.map_info[i].bookmark;
            console.log("b")
          }
          let newData = {
            colorWeight : colorWeight * 100, 
            description1 : des[0],
            description2 : des[1],
            description3 : des[2],
            description4 : des[3],
            description5 : des[4],
            description6 : des[5],
            description7 : des[6],
            description8 : des[7],
          };
          
          mapData.values[data.map_info[i].id] = newData;
          console.log(mapData.values[data.map_info[i].id])
          jpName[data.map_info[i].id] = data.map_info[i].name;
        }
        new svgMap({
          targetElementID: 'svgMap',
          data: mapData,
          countryNames: jpName,
          colorMax: '#ED6E67',
          colorMin: '#F3CF85',
          colorNoData: '#E0DCDB',
        });
      }))
      return {
        auth,
      }
  }
}
</script>

<style>
.svgMap-map-wrapper{
  background: #6B9BA9;
}
.btn-toolbar{
  position: absolute;
  z-index: 1;
}
.btn-circle {
  width: 45px;
  height: 45px;
  line-height: 45px;
  text-align: center;
  padding: 0;
  border-radius: 50%;
}
.btn-circle i {
  position: relative;
  top: -1px;
}
.bi{
  color: #FFCF32;
}
</style>