<template>
  <div id="svgMap" style="position: relative;">
    <QuizFillter></QuizFillter>
    <LoginDialog v-if="!auth"></LoginDialog>
    <UserInfoDialog v-if="auth"></UserInfoDialog>
  </div>
  <input type="checkbox" class="openSidebarMenu" id="openSidebarMenu">
  <label for="openSidebarMenu" class="sidebarIconToggle">
    <div class="spinner diagonal part-1"></div>
    <div class="spinner horizontal"></div>
    <div class="spinner diagonal part-2"></div>
  </label>
  <div id="sidebarMenu">
    <ul class="sidebarMenuInner">
      <li></li>
      <li><a href="#" data-toggle="modal" data-target="#exampleModalCenter">quiz</a></li>
    </ul>
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
          countryNames: jpName
        });
      }))
      return {
        auth,
      }
  }
}
</script>

<style>
#sidebarMenu {
    height: 100%;
    position: fixed;
    left: 0;
    width: 250px;
    transform: translateX(-250px);
    transition: transform 250ms ease-in-out;
    background: #fff;
}
.sidebarMenuInner{
    margin:0;
    padding:0;
    border-top: 1px solid rgba(255, 255, 255, 0.10);
}
.sidebarMenuInner li{
    list-style: none;
    color: #fff;
    font-weight: bold;
    padding: 20px;
    cursor: pointer;
    border-bottom: 1px solid rgba(255, 255, 255, 0.10);
}
/* .sidebarMenuInner li span{
    display: block;
    font-size: 14px;
    color: #fff;
} */
.sidebarMenuInner li a{
    color: #000;
    font-weight: bold;
    cursor: pointer;
    text-decoration: none;
}
input[type="checkbox"]:checked ~ #sidebarMenu {
    transform: translateX(0);
}

input[type=checkbox] {
    transition: all 0.3s;
    box-sizing: border-box;
    display: none;
}
.sidebarIconToggle {
    transition: all 0.3s;
    box-sizing: border-box;
    cursor: pointer;
    position: absolute;
    z-index: 99;
    height: 100%;
    width: 100%;
    top: 22px;
    left: 15px;
    height: 22px;
    width: 22px;
}
.spinner {
    transition: all 0.3s;
    box-sizing: border-box;
    position: absolute;
    height: 3px;
    width: 100%;
    background-color: #000;
}
.horizontal {
    transition: all 0.3s;
    box-sizing: border-box;
    position: relative;
    float: left;
    margin-top: 3px;
}
.diagonal.part-1 {
    position: relative;
    transition: all 0.3s;
    box-sizing: border-box;
    float: left;
}
.diagonal.part-2 {
    transition: all 0.3s;
    box-sizing: border-box;
    position: relative;
    float: left;
    margin-top: 3px;
}
input[type=checkbox]:checked ~ .sidebarIconToggle > .horizontal {
    transition: all 0.3s;
    box-sizing: border-box;
    opacity: 0;
}
input[type=checkbox]:checked ~ .sidebarIconToggle > .diagonal.part-1 {
    transition: all 0.3s;
    box-sizing: border-box;
    transform: rotate(135deg);
    margin-top: 8px;
}
input[type=checkbox]:checked ~ .sidebarIconToggle > .diagonal.part-2 {
    transition: all 0.3s;
    box-sizing: border-box;
    transform: rotate(-135deg);
    margin-top: -9px;
}
</style>