<template>
  <div id="svgMap">
  </div>
</template>

<script>
import mapData from "../data/mapData";
import axios from 'axios';
import svgMap from 'svgmap';
import 'svgmap/dist/svgMap.min.css';
export default {
  name: 'WorldMap',

  props: {
    mapMode: String,
  },

  setup (props) {    
    axios.get('http://localhost:8888/map')
      .then((res =>{
        let data = res.data, jpName = {}, maxColorCode, minColorCode, noDataColorCode;
        for(let i=0; i<data.map_info.length; i++){
          const des = data.map_info[i].description.match(/.{1,25}/g);
          for(let i=0; i<8; i++){
            if(!des[i]) des[i] = '';
          }
          let colorWeight;
          if (props.mapMode == 'correctAnswersRate'){
            colorWeight = data.map_info[i].weight;
            maxColorCode = '#F28B7C';
            minColorCode = '#F6D0CF';
            noDataColorCode = '#F6D0CF';
          }else if (props.mapMode == 'bookMark'){
            colorWeight = data.map_info[i].bookmark;
            maxColorCode = '#76926A';
            minColorCode = '#D2DDBF';
            noDataColorCode = '#D2DDBF';
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
          jpName[data.map_info[i].id] = data.map_info[i].name;
        }
        new svgMap({
          targetElementID: 'svgMap',
          data: mapData,
          countryNames: jpName,
          colorMax: maxColorCode,
          colorMin: minColorCode,
          colorNoData: noDataColorCode,
        });
      }))
      return {
      }
  }
}
</script>
<style>
.svgMap-map-wrapper {
  background: #9ABBDF;
  min-height: 100vh;
}
</style>