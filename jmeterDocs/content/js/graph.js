/*
   Licensed to the Apache Software Foundation (ASF) under one or more
   contributor license agreements.  See the NOTICE file distributed with
   this work for additional information regarding copyright ownership.
   The ASF licenses this file to You under the Apache License, Version 2.0
   (the "License"); you may not use this file except in compliance with
   the License.  You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
$(document).ready(function() {

    $(".click-title").mouseenter( function(    e){
        e.preventDefault();
        this.style.cursor="pointer";
    });
    $(".click-title").mousedown( function(event){
        event.preventDefault();
    });

    // Ugly code while this script is shared among several pages
    try{
        refreshHitsPerSecond(true);
    } catch(e){}
    try{
        refreshResponseTimeOverTime(true);
    } catch(e){}
    try{
        refreshResponseTimePercentiles();
    } catch(e){}
});


var responseTimePercentilesInfos = {
        data: {"result": {"minY": 0.0, "minX": 0.0, "maxY": 683.0, "series": [{"data": [[0.0, 0.0], [0.1, 0.0], [0.2, 0.0], [0.3, 0.0], [0.4, 0.0], [0.5, 0.0], [0.6, 0.0], [0.7, 0.0], [0.8, 0.0], [0.9, 0.0], [1.0, 0.0], [1.1, 0.0], [1.2, 0.0], [1.3, 0.0], [1.4, 0.0], [1.5, 0.0], [1.6, 0.0], [1.7, 0.0], [1.8, 0.0], [1.9, 0.0], [2.0, 0.0], [2.1, 0.0], [2.2, 0.0], [2.3, 0.0], [2.4, 0.0], [2.5, 0.0], [2.6, 0.0], [2.7, 0.0], [2.8, 0.0], [2.9, 0.0], [3.0, 0.0], [3.1, 0.0], [3.2, 0.0], [3.3, 0.0], [3.4, 0.0], [3.5, 0.0], [3.6, 0.0], [3.7, 0.0], [3.8, 0.0], [3.9, 0.0], [4.0, 0.0], [4.1, 0.0], [4.2, 0.0], [4.3, 0.0], [4.4, 0.0], [4.5, 0.0], [4.6, 0.0], [4.7, 0.0], [4.8, 0.0], [4.9, 0.0], [5.0, 0.0], [5.1, 0.0], [5.2, 0.0], [5.3, 0.0], [5.4, 0.0], [5.5, 0.0], [5.6, 0.0], [5.7, 0.0], [5.8, 0.0], [5.9, 0.0], [6.0, 0.0], [6.1, 0.0], [6.2, 0.0], [6.3, 0.0], [6.4, 0.0], [6.5, 0.0], [6.6, 0.0], [6.7, 0.0], [6.8, 0.0], [6.9, 0.0], [7.0, 0.0], [7.1, 0.0], [7.2, 0.0], [7.3, 0.0], [7.4, 0.0], [7.5, 0.0], [7.6, 0.0], [7.7, 0.0], [7.8, 0.0], [7.9, 0.0], [8.0, 0.0], [8.1, 0.0], [8.2, 0.0], [8.3, 0.0], [8.4, 0.0], [8.5, 0.0], [8.6, 0.0], [8.7, 0.0], [8.8, 0.0], [8.9, 0.0], [9.0, 0.0], [9.1, 0.0], [9.2, 0.0], [9.3, 0.0], [9.4, 0.0], [9.5, 0.0], [9.6, 0.0], [9.7, 0.0], [9.8, 0.0], [9.9, 0.0], [10.0, 0.0], [10.1, 0.0], [10.2, 0.0], [10.3, 0.0], [10.4, 0.0], [10.5, 0.0], [10.6, 0.0], [10.7, 0.0], [10.8, 0.0], [10.9, 0.0], [11.0, 0.0], [11.1, 0.0], [11.2, 0.0], [11.3, 0.0], [11.4, 0.0], [11.5, 0.0], [11.6, 0.0], [11.7, 0.0], [11.8, 0.0], [11.9, 0.0], [12.0, 0.0], [12.1, 0.0], [12.2, 0.0], [12.3, 0.0], [12.4, 0.0], [12.5, 0.0], [12.6, 0.0], [12.7, 0.0], [12.8, 0.0], [12.9, 0.0], [13.0, 0.0], [13.1, 0.0], [13.2, 0.0], [13.3, 1.0], [13.4, 1.0], [13.5, 1.0], [13.6, 1.0], [13.7, 1.0], [13.8, 1.0], [13.9, 1.0], [14.0, 1.0], [14.1, 1.0], [14.2, 1.0], [14.3, 1.0], [14.4, 1.0], [14.5, 1.0], [14.6, 1.0], [14.7, 1.0], [14.8, 1.0], [14.9, 1.0], [15.0, 1.0], [15.1, 1.0], [15.2, 1.0], [15.3, 1.0], [15.4, 1.0], [15.5, 1.0], [15.6, 1.0], [15.7, 1.0], [15.8, 1.0], [15.9, 1.0], [16.0, 1.0], [16.1, 1.0], [16.2, 1.0], [16.3, 1.0], [16.4, 1.0], [16.5, 1.0], [16.6, 1.0], [16.7, 1.0], [16.8, 1.0], [16.9, 1.0], [17.0, 1.0], [17.1, 1.0], [17.2, 1.0], [17.3, 1.0], [17.4, 1.0], [17.5, 1.0], [17.6, 1.0], [17.7, 1.0], [17.8, 1.0], [17.9, 1.0], [18.0, 1.0], [18.1, 1.0], [18.2, 1.0], [18.3, 1.0], [18.4, 1.0], [18.5, 1.0], [18.6, 1.0], [18.7, 1.0], [18.8, 1.0], [18.9, 1.0], [19.0, 1.0], [19.1, 1.0], [19.2, 1.0], [19.3, 1.0], [19.4, 1.0], [19.5, 1.0], [19.6, 1.0], [19.7, 1.0], [19.8, 1.0], [19.9, 1.0], [20.0, 1.0], [20.1, 1.0], [20.2, 1.0], [20.3, 1.0], [20.4, 1.0], [20.5, 1.0], [20.6, 1.0], [20.7, 1.0], [20.8, 1.0], [20.9, 1.0], [21.0, 1.0], [21.1, 1.0], [21.2, 1.0], [21.3, 1.0], [21.4, 1.0], [21.5, 1.0], [21.6, 1.0], [21.7, 1.0], [21.8, 1.0], [21.9, 1.0], [22.0, 1.0], [22.1, 1.0], [22.2, 1.0], [22.3, 1.0], [22.4, 1.0], [22.5, 1.0], [22.6, 1.0], [22.7, 1.0], [22.8, 1.0], [22.9, 1.0], [23.0, 1.0], [23.1, 1.0], [23.2, 1.0], [23.3, 1.0], [23.4, 1.0], [23.5, 1.0], [23.6, 1.0], [23.7, 1.0], [23.8, 1.0], [23.9, 1.0], [24.0, 1.0], [24.1, 1.0], [24.2, 1.0], [24.3, 1.0], [24.4, 1.0], [24.5, 1.0], [24.6, 1.0], [24.7, 1.0], [24.8, 1.0], [24.9, 1.0], [25.0, 1.0], [25.1, 1.0], [25.2, 1.0], [25.3, 1.0], [25.4, 1.0], [25.5, 1.0], [25.6, 1.0], [25.7, 1.0], [25.8, 1.0], [25.9, 1.0], [26.0, 1.0], [26.1, 1.0], [26.2, 1.0], [26.3, 1.0], [26.4, 1.0], [26.5, 1.0], [26.6, 1.0], [26.7, 1.0], [26.8, 1.0], [26.9, 1.0], [27.0, 1.0], [27.1, 1.0], [27.2, 1.0], [27.3, 1.0], [27.4, 1.0], [27.5, 1.0], [27.6, 1.0], [27.7, 1.0], [27.8, 1.0], [27.9, 1.0], [28.0, 1.0], [28.1, 1.0], [28.2, 1.0], [28.3, 1.0], [28.4, 1.0], [28.5, 1.0], [28.6, 1.0], [28.7, 1.0], [28.8, 1.0], [28.9, 1.0], [29.0, 1.0], [29.1, 1.0], [29.2, 1.0], [29.3, 1.0], [29.4, 1.0], [29.5, 1.0], [29.6, 1.0], [29.7, 1.0], [29.8, 1.0], [29.9, 1.0], [30.0, 1.0], [30.1, 1.0], [30.2, 1.0], [30.3, 1.0], [30.4, 1.0], [30.5, 1.0], [30.6, 1.0], [30.7, 1.0], [30.8, 1.0], [30.9, 1.0], [31.0, 1.0], [31.1, 1.0], [31.2, 1.0], [31.3, 1.0], [31.4, 1.0], [31.5, 1.0], [31.6, 1.0], [31.7, 1.0], [31.8, 1.0], [31.9, 1.0], [32.0, 1.0], [32.1, 1.0], [32.2, 1.0], [32.3, 1.0], [32.4, 1.0], [32.5, 1.0], [32.6, 1.0], [32.7, 1.0], [32.8, 1.0], [32.9, 1.0], [33.0, 1.0], [33.1, 1.0], [33.2, 1.0], [33.3, 1.0], [33.4, 1.0], [33.5, 1.0], [33.6, 1.0], [33.7, 1.0], [33.8, 1.0], [33.9, 1.0], [34.0, 1.0], [34.1, 1.0], [34.2, 1.0], [34.3, 1.0], [34.4, 1.0], [34.5, 1.0], [34.6, 1.0], [34.7, 1.0], [34.8, 1.0], [34.9, 1.0], [35.0, 1.0], [35.1, 1.0], [35.2, 1.0], [35.3, 1.0], [35.4, 1.0], [35.5, 1.0], [35.6, 1.0], [35.7, 1.0], [35.8, 1.0], [35.9, 1.0], [36.0, 1.0], [36.1, 1.0], [36.2, 1.0], [36.3, 1.0], [36.4, 1.0], [36.5, 1.0], [36.6, 1.0], [36.7, 1.0], [36.8, 1.0], [36.9, 1.0], [37.0, 1.0], [37.1, 1.0], [37.2, 1.0], [37.3, 1.0], [37.4, 1.0], [37.5, 1.0], [37.6, 1.0], [37.7, 1.0], [37.8, 1.0], [37.9, 1.0], [38.0, 1.0], [38.1, 1.0], [38.2, 1.0], [38.3, 1.0], [38.4, 1.0], [38.5, 1.0], [38.6, 1.0], [38.7, 1.0], [38.8, 1.0], [38.9, 1.0], [39.0, 1.0], [39.1, 1.0], [39.2, 1.0], [39.3, 1.0], [39.4, 1.0], [39.5, 1.0], [39.6, 1.0], [39.7, 1.0], [39.8, 1.0], [39.9, 1.0], [40.0, 1.0], [40.1, 1.0], [40.2, 1.0], [40.3, 1.0], [40.4, 1.0], [40.5, 1.0], [40.6, 1.0], [40.7, 1.0], [40.8, 1.0], [40.9, 1.0], [41.0, 1.0], [41.1, 1.0], [41.2, 1.0], [41.3, 1.0], [41.4, 1.0], [41.5, 1.0], [41.6, 1.0], [41.7, 1.0], [41.8, 1.0], [41.9, 1.0], [42.0, 1.0], [42.1, 1.0], [42.2, 1.0], [42.3, 1.0], [42.4, 1.0], [42.5, 1.0], [42.6, 1.0], [42.7, 1.0], [42.8, 1.0], [42.9, 1.0], [43.0, 1.0], [43.1, 1.0], [43.2, 1.0], [43.3, 1.0], [43.4, 1.0], [43.5, 1.0], [43.6, 1.0], [43.7, 1.0], [43.8, 1.0], [43.9, 1.0], [44.0, 1.0], [44.1, 1.0], [44.2, 1.0], [44.3, 1.0], [44.4, 1.0], [44.5, 1.0], [44.6, 1.0], [44.7, 2.0], [44.8, 2.0], [44.9, 2.0], [45.0, 2.0], [45.1, 2.0], [45.2, 2.0], [45.3, 2.0], [45.4, 2.0], [45.5, 2.0], [45.6, 2.0], [45.7, 2.0], [45.8, 2.0], [45.9, 2.0], [46.0, 2.0], [46.1, 2.0], [46.2, 2.0], [46.3, 2.0], [46.4, 2.0], [46.5, 2.0], [46.6, 2.0], [46.7, 2.0], [46.8, 2.0], [46.9, 2.0], [47.0, 2.0], [47.1, 2.0], [47.2, 2.0], [47.3, 2.0], [47.4, 2.0], [47.5, 2.0], [47.6, 2.0], [47.7, 2.0], [47.8, 2.0], [47.9, 2.0], [48.0, 2.0], [48.1, 2.0], [48.2, 2.0], [48.3, 2.0], [48.4, 2.0], [48.5, 2.0], [48.6, 2.0], [48.7, 2.0], [48.8, 2.0], [48.9, 2.0], [49.0, 2.0], [49.1, 2.0], [49.2, 2.0], [49.3, 2.0], [49.4, 2.0], [49.5, 2.0], [49.6, 2.0], [49.7, 2.0], [49.8, 2.0], [49.9, 2.0], [50.0, 2.0], [50.1, 2.0], [50.2, 2.0], [50.3, 2.0], [50.4, 2.0], [50.5, 3.0], [50.6, 3.0], [50.7, 3.0], [50.8, 3.0], [50.9, 3.0], [51.0, 3.0], [51.1, 3.0], [51.2, 3.0], [51.3, 3.0], [51.4, 3.0], [51.5, 3.0], [51.6, 3.0], [51.7, 3.0], [51.8, 3.0], [51.9, 3.0], [52.0, 3.0], [52.1, 3.0], [52.2, 3.0], [52.3, 3.0], [52.4, 3.0], [52.5, 3.0], [52.6, 3.0], [52.7, 3.0], [52.8, 3.0], [52.9, 3.0], [53.0, 3.0], [53.1, 3.0], [53.2, 3.0], [53.3, 3.0], [53.4, 3.0], [53.5, 3.0], [53.6, 3.0], [53.7, 3.0], [53.8, 3.0], [53.9, 3.0], [54.0, 3.0], [54.1, 3.0], [54.2, 3.0], [54.3, 3.0], [54.4, 3.0], [54.5, 3.0], [54.6, 3.0], [54.7, 3.0], [54.8, 3.0], [54.9, 3.0], [55.0, 3.0], [55.1, 3.0], [55.2, 3.0], [55.3, 3.0], [55.4, 3.0], [55.5, 3.0], [55.6, 3.0], [55.7, 3.0], [55.8, 3.0], [55.9, 3.0], [56.0, 3.0], [56.1, 3.0], [56.2, 3.0], [56.3, 3.0], [56.4, 3.0], [56.5, 3.0], [56.6, 3.0], [56.7, 3.0], [56.8, 3.0], [56.9, 3.0], [57.0, 3.0], [57.1, 3.0], [57.2, 3.0], [57.3, 3.0], [57.4, 3.0], [57.5, 3.0], [57.6, 3.0], [57.7, 3.0], [57.8, 3.0], [57.9, 3.0], [58.0, 3.0], [58.1, 3.0], [58.2, 3.0], [58.3, 3.0], [58.4, 3.0], [58.5, 3.0], [58.6, 3.0], [58.7, 3.0], [58.8, 3.0], [58.9, 3.0], [59.0, 3.0], [59.1, 3.0], [59.2, 3.0], [59.3, 3.0], [59.4, 3.0], [59.5, 3.0], [59.6, 3.0], [59.7, 3.0], [59.8, 3.0], [59.9, 3.0], [60.0, 3.0], [60.1, 3.0], [60.2, 3.0], [60.3, 3.0], [60.4, 3.0], [60.5, 3.0], [60.6, 3.0], [60.7, 3.0], [60.8, 3.0], [60.9, 3.0], [61.0, 3.0], [61.1, 3.0], [61.2, 3.0], [61.3, 3.0], [61.4, 3.0], [61.5, 3.0], [61.6, 3.0], [61.7, 3.0], [61.8, 3.0], [61.9, 3.0], [62.0, 3.0], [62.1, 3.0], [62.2, 3.0], [62.3, 3.0], [62.4, 3.0], [62.5, 3.0], [62.6, 3.0], [62.7, 3.0], [62.8, 3.0], [62.9, 3.0], [63.0, 3.0], [63.1, 3.0], [63.2, 3.0], [63.3, 3.0], [63.4, 3.0], [63.5, 3.0], [63.6, 3.0], [63.7, 3.0], [63.8, 3.0], [63.9, 3.0], [64.0, 3.0], [64.1, 3.0], [64.2, 3.0], [64.3, 3.0], [64.4, 4.0], [64.5, 4.0], [64.6, 4.0], [64.7, 4.0], [64.8, 4.0], [64.9, 4.0], [65.0, 4.0], [65.1, 4.0], [65.2, 4.0], [65.3, 4.0], [65.4, 4.0], [65.5, 4.0], [65.6, 4.0], [65.7, 4.0], [65.8, 4.0], [65.9, 4.0], [66.0, 4.0], [66.1, 4.0], [66.2, 4.0], [66.3, 4.0], [66.4, 4.0], [66.5, 4.0], [66.6, 4.0], [66.7, 4.0], [66.8, 4.0], [66.9, 4.0], [67.0, 4.0], [67.1, 4.0], [67.2, 4.0], [67.3, 4.0], [67.4, 4.0], [67.5, 4.0], [67.6, 4.0], [67.7, 4.0], [67.8, 4.0], [67.9, 4.0], [68.0, 4.0], [68.1, 4.0], [68.2, 4.0], [68.3, 4.0], [68.4, 4.0], [68.5, 4.0], [68.6, 4.0], [68.7, 4.0], [68.8, 4.0], [68.9, 4.0], [69.0, 4.0], [69.1, 4.0], [69.2, 4.0], [69.3, 4.0], [69.4, 4.0], [69.5, 4.0], [69.6, 4.0], [69.7, 4.0], [69.8, 4.0], [69.9, 4.0], [70.0, 4.0], [70.1, 4.0], [70.2, 4.0], [70.3, 4.0], [70.4, 4.0], [70.5, 4.0], [70.6, 4.0], [70.7, 4.0], [70.8, 4.0], [70.9, 4.0], [71.0, 4.0], [71.1, 4.0], [71.2, 4.0], [71.3, 4.0], [71.4, 4.0], [71.5, 4.0], [71.6, 4.0], [71.7, 4.0], [71.8, 4.0], [71.9, 4.0], [72.0, 4.0], [72.1, 4.0], [72.2, 4.0], [72.3, 4.0], [72.4, 4.0], [72.5, 4.0], [72.6, 4.0], [72.7, 4.0], [72.8, 4.0], [72.9, 4.0], [73.0, 5.0], [73.1, 5.0], [73.2, 5.0], [73.3, 5.0], [73.4, 5.0], [73.5, 5.0], [73.6, 5.0], [73.7, 5.0], [73.8, 5.0], [73.9, 5.0], [74.0, 5.0], [74.1, 5.0], [74.2, 5.0], [74.3, 5.0], [74.4, 5.0], [74.5, 5.0], [74.6, 5.0], [74.7, 5.0], [74.8, 5.0], [74.9, 5.0], [75.0, 5.0], [75.1, 5.0], [75.2, 5.0], [75.3, 5.0], [75.4, 5.0], [75.5, 5.0], [75.6, 5.0], [75.7, 5.0], [75.8, 5.0], [75.9, 5.0], [76.0, 5.0], [76.1, 5.0], [76.2, 5.0], [76.3, 5.0], [76.4, 5.0], [76.5, 5.0], [76.6, 5.0], [76.7, 5.0], [76.8, 5.0], [76.9, 5.0], [77.0, 5.0], [77.1, 5.0], [77.2, 5.0], [77.3, 5.0], [77.4, 5.0], [77.5, 5.0], [77.6, 5.0], [77.7, 5.0], [77.8, 6.0], [77.9, 6.0], [78.0, 6.0], [78.1, 6.0], [78.2, 6.0], [78.3, 6.0], [78.4, 6.0], [78.5, 6.0], [78.6, 6.0], [78.7, 6.0], [78.8, 6.0], [78.9, 6.0], [79.0, 6.0], [79.1, 6.0], [79.2, 6.0], [79.3, 6.0], [79.4, 6.0], [79.5, 6.0], [79.6, 6.0], [79.7, 6.0], [79.8, 6.0], [79.9, 6.0], [80.0, 6.0], [80.1, 6.0], [80.2, 6.0], [80.3, 6.0], [80.4, 6.0], [80.5, 6.0], [80.6, 6.0], [80.7, 6.0], [80.8, 6.0], [80.9, 6.0], [81.0, 6.0], [81.1, 6.0], [81.2, 6.0], [81.3, 6.0], [81.4, 6.0], [81.5, 6.0], [81.6, 6.0], [81.7, 6.0], [81.8, 6.0], [81.9, 6.0], [82.0, 6.0], [82.1, 6.0], [82.2, 6.0], [82.3, 6.0], [82.4, 6.0], [82.5, 6.0], [82.6, 6.0], [82.7, 6.0], [82.8, 6.0], [82.9, 6.0], [83.0, 6.0], [83.1, 6.0], [83.2, 6.0], [83.3, 6.0], [83.4, 6.0], [83.5, 6.0], [83.6, 6.0], [83.7, 6.0], [83.8, 6.0], [83.9, 6.0], [84.0, 6.0], [84.1, 6.0], [84.2, 6.0], [84.3, 6.0], [84.4, 6.0], [84.5, 6.0], [84.6, 7.0], [84.7, 7.0], [84.8, 7.0], [84.9, 7.0], [85.0, 7.0], [85.1, 7.0], [85.2, 7.0], [85.3, 7.0], [85.4, 7.0], [85.5, 7.0], [85.6, 7.0], [85.7, 7.0], [85.8, 7.0], [85.9, 7.0], [86.0, 7.0], [86.1, 7.0], [86.2, 7.0], [86.3, 7.0], [86.4, 7.0], [86.5, 7.0], [86.6, 7.0], [86.7, 7.0], [86.8, 7.0], [86.9, 7.0], [87.0, 7.0], [87.1, 7.0], [87.2, 7.0], [87.3, 8.0], [87.4, 8.0], [87.5, 8.0], [87.6, 8.0], [87.7, 8.0], [87.8, 8.0], [87.9, 8.0], [88.0, 8.0], [88.1, 8.0], [88.2, 8.0], [88.3, 8.0], [88.4, 8.0], [88.5, 8.0], [88.6, 8.0], [88.7, 8.0], [88.8, 8.0], [88.9, 8.0], [89.0, 8.0], [89.1, 8.0], [89.2, 8.0], [89.3, 8.0], [89.4, 8.0], [89.5, 8.0], [89.6, 8.0], [89.7, 8.0], [89.8, 8.0], [89.9, 8.0], [90.0, 8.0], [90.1, 8.0], [90.2, 8.0], [90.3, 8.0], [90.4, 8.0], [90.5, 8.0], [90.6, 9.0], [90.7, 9.0], [90.8, 9.0], [90.9, 9.0], [91.0, 9.0], [91.1, 9.0], [91.2, 9.0], [91.3, 9.0], [91.4, 9.0], [91.5, 9.0], [91.6, 9.0], [91.7, 9.0], [91.8, 9.0], [91.9, 9.0], [92.0, 9.0], [92.1, 9.0], [92.2, 9.0], [92.3, 9.0], [92.4, 9.0], [92.5, 9.0], [92.6, 9.0], [92.7, 9.0], [92.8, 9.0], [92.9, 9.0], [93.0, 9.0], [93.1, 10.0], [93.2, 10.0], [93.3, 10.0], [93.4, 10.0], [93.5, 10.0], [93.6, 10.0], [93.7, 10.0], [93.8, 10.0], [93.9, 10.0], [94.0, 10.0], [94.1, 10.0], [94.2, 10.0], [94.3, 10.0], [94.4, 11.0], [94.5, 11.0], [94.6, 11.0], [94.7, 11.0], [94.8, 11.0], [94.9, 11.0], [95.0, 11.0], [95.1, 11.0], [95.2, 11.0], [95.3, 11.0], [95.4, 11.0], [95.5, 11.0], [95.6, 11.0], [95.7, 11.0], [95.8, 11.0], [95.9, 12.0], [96.0, 12.0], [96.1, 12.0], [96.2, 12.0], [96.3, 12.0], [96.4, 12.0], [96.5, 12.0], [96.6, 12.0], [96.7, 12.0], [96.8, 13.0], [96.9, 13.0], [97.0, 13.0], [97.1, 13.0], [97.2, 13.0], [97.3, 13.0], [97.4, 13.0], [97.5, 14.0], [97.6, 14.0], [97.7, 14.0], [97.8, 14.0], [97.9, 14.0], [98.0, 14.0], [98.1, 15.0], [98.2, 15.0], [98.3, 15.0], [98.4, 15.0], [98.5, 16.0], [98.6, 16.0], [98.7, 16.0], [98.8, 16.0], [98.9, 17.0], [99.0, 17.0], [99.1, 17.0], [99.2, 18.0], [99.3, 18.0], [99.4, 19.0], [99.5, 20.0], [99.6, 20.0], [99.7, 21.0], [99.8, 23.0], [99.9, 26.0], [100.0, 683.0]], "isOverall": false, "label": "HTTP 要求", "isController": false}], "supportsControllersDiscrimination": true, "maxX": 100.0, "title": "Response Time Percentiles"}},
        getOptions: function() {
            return {
                series: {
                    points: { show: false }
                },
                legend: {
                    noColumns: 2,
                    show: true,
                    container: '#legendResponseTimePercentiles'
                },
                xaxis: {
                    tickDecimals: 1,
                    axisLabel: "Percentiles",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Percentile value in ms",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                tooltip: true,
                tooltipOpts: {
                    content: "%s : %x.2 percentile was %y ms"
                },
                selection: { mode: "xy" },
            };
        },
        createGraph: function() {
            var data = this.data;
            var dataset = prepareData(data.result.series, $("#choicesResponseTimePercentiles"));
            var options = this.getOptions();
            prepareOptions(options, data);
            $.plot($("#flotResponseTimesPercentiles"), dataset, options);
            // setup overview
            $.plot($("#overviewResponseTimesPercentiles"), dataset, prepareOverviewOptions(options));
        }
};

/**
 * @param elementId Id of element where we display message
 */
function setEmptyGraph(elementId) {
    $(function() {
        $(elementId).text("No graph series with filter="+seriesFilter);
    });
}

// Response times percentiles
function refreshResponseTimePercentiles() {
    var infos = responseTimePercentilesInfos;
    prepareSeries(infos.data);
    if(infos.data.result.series.length == 0) {
        setEmptyGraph("#bodyResponseTimePercentiles");
        return;
    }
    if (isGraph($("#flotResponseTimesPercentiles"))){
        infos.createGraph();
    } else {
        var choiceContainer = $("#choicesResponseTimePercentiles");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotResponseTimesPercentiles", "#overviewResponseTimesPercentiles");
        $('#bodyResponseTimePercentiles .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
}

var responseTimeDistributionInfos = {
        data: {"result": {"minY": 2.0, "minX": 0.0, "maxY": 1.7106961E7, "series": [{"data": [[0.0, 1.7106961E7], [300.0, 16.0], [600.0, 3.0], [100.0, 2084.0], [200.0, 245.0], [400.0, 4.0], [500.0, 2.0]], "isOverall": false, "label": "HTTP 要求", "isController": false}], "supportsControllersDiscrimination": true, "granularity": 100, "maxX": 600.0, "title": "Response Time Distribution"}},
        getOptions: function() {
            var granularity = this.data.result.granularity;
            return {
                legend: {
                    noColumns: 2,
                    show: true,
                    container: '#legendResponseTimeDistribution'
                },
                xaxis:{
                    axisLabel: "Response times in ms",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Number of responses",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                bars : {
                    show: true,
                    barWidth: this.data.result.granularity
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                tooltip: true,
                tooltipOpts: {
                    content: function(label, xval, yval, flotItem){
                        return yval + " responses for " + label + " were between " + xval + " and " + (xval + granularity) + " ms";
                    }
                }
            };
        },
        createGraph: function() {
            var data = this.data;
            var options = this.getOptions();
            prepareOptions(options, data);
            $.plot($("#flotResponseTimeDistribution"), prepareData(data.result.series, $("#choicesResponseTimeDistribution")), options);
        }

};

// Response time distribution
function refreshResponseTimeDistribution() {
    var infos = responseTimeDistributionInfos;
    prepareSeries(infos.data);
    if(infos.data.result.series.length == 0) {
        setEmptyGraph("#bodyResponseTimeDistribution");
        return;
    }
    if (isGraph($("#flotResponseTimeDistribution"))){
        infos.createGraph();
    }else{
        var choiceContainer = $("#choicesResponseTimeDistribution");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        $('#footerResponseTimeDistribution .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
};


var syntheticResponseTimeDistributionInfos = {
        data: {"result": {"minY": 5.0, "minX": 0.0, "ticks": [[0, "Requests having \nresponse time <= 500ms"], [1, "Requests having \nresponse time > 500ms and <= 1,500ms"], [2, "Requests having \nresponse time > 1,500ms"], [3, "Requests in error"]], "maxY": 1.710931E7, "series": [{"data": [[0.0, 1.710931E7]], "color": "#9ACD32", "isOverall": false, "label": "Requests having \nresponse time <= 500ms", "isController": false}, {"data": [[1.0, 5.0]], "color": "yellow", "isOverall": false, "label": "Requests having \nresponse time > 500ms and <= 1,500ms", "isController": false}, {"data": [], "color": "orange", "isOverall": false, "label": "Requests having \nresponse time > 1,500ms", "isController": false}, {"data": [], "color": "#FF6347", "isOverall": false, "label": "Requests in error", "isController": false}], "supportsControllersDiscrimination": false, "maxX": 1.0, "title": "Synthetic Response Times Distribution"}},
        getOptions: function() {
            return {
                legend: {
                    noColumns: 2,
                    show: true,
                    container: '#legendSyntheticResponseTimeDistribution'
                },
                xaxis:{
                    axisLabel: "Response times ranges",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                    tickLength:0,
                    min:-0.5,
                    max:3.5
                },
                yaxis: {
                    axisLabel: "Number of responses",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                bars : {
                    show: true,
                    align: "center",
                    barWidth: 0.25,
                    fill:.75
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                tooltip: true,
                tooltipOpts: {
                    content: function(label, xval, yval, flotItem){
                        return yval + " " + label;
                    }
                }
            };
        },
        createGraph: function() {
            var data = this.data;
            var options = this.getOptions();
            prepareOptions(options, data);
            options.xaxis.ticks = data.result.ticks;
            $.plot($("#flotSyntheticResponseTimeDistribution"), prepareData(data.result.series, $("#choicesSyntheticResponseTimeDistribution")), options);
        }

};

// Response time distribution
function refreshSyntheticResponseTimeDistribution() {
    var infos = syntheticResponseTimeDistributionInfos;
    prepareSeries(infos.data, true);
    if (isGraph($("#flotSyntheticResponseTimeDistribution"))){
        infos.createGraph();
    }else{
        var choiceContainer = $("#choicesSyntheticResponseTimeDistribution");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        $('#footerSyntheticResponseTimeDistribution .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
};

var activeThreadsOverTimeInfos = {
        data: {"result": {"minY": 99.42192173815081, "minX": 1.71086382E12, "maxY": 100.0, "series": [{"data": [[1.71086382E12, 99.42192173815081], [1.7108643E12, 100.0], [1.71086412E12, 100.0], [1.71086394E12, 100.0], [1.71086442E12, 99.99389707706072], [1.71086424E12, 100.0], [1.71086406E12, 100.0], [1.71086388E12, 100.0], [1.71086436E12, 100.0], [1.71086418E12, 100.0], [1.710864E12, 100.0]], "isOverall": false, "label": "test", "isController": false}], "supportsControllersDiscrimination": false, "granularity": 60000, "maxX": 1.71086442E12, "title": "Active Threads Over Time"}},
        getOptions: function() {
            return {
                series: {
                    stack: true,
                    lines: {
                        show: true,
                        fill: true
                    },
                    points: {
                        show: true
                    }
                },
                xaxis: {
                    mode: "time",
                    timeformat: getTimeFormat(this.data.result.granularity),
                    axisLabel: getElapsedTimeLabel(this.data.result.granularity),
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Number of active threads",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20
                },
                legend: {
                    noColumns: 6,
                    show: true,
                    container: '#legendActiveThreadsOverTime'
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                selection: {
                    mode: 'xy'
                },
                tooltip: true,
                tooltipOpts: {
                    content: "%s : At %x there were %y active threads"
                }
            };
        },
        createGraph: function() {
            var data = this.data;
            var dataset = prepareData(data.result.series, $("#choicesActiveThreadsOverTime"));
            var options = this.getOptions();
            prepareOptions(options, data);
            $.plot($("#flotActiveThreadsOverTime"), dataset, options);
            // setup overview
            $.plot($("#overviewActiveThreadsOverTime"), dataset, prepareOverviewOptions(options));
        }
};

// Active Threads Over Time
function refreshActiveThreadsOverTime(fixTimestamps) {
    var infos = activeThreadsOverTimeInfos;
    prepareSeries(infos.data);
    if(fixTimestamps) {
        fixTimeStamps(infos.data.result.series, 28800000);
    }
    if(isGraph($("#flotActiveThreadsOverTime"))) {
        infos.createGraph();
    }else{
        var choiceContainer = $("#choicesActiveThreadsOverTime");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotActiveThreadsOverTime", "#overviewActiveThreadsOverTime");
        $('#footerActiveThreadsOverTime .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
};

var timeVsThreadsInfos = {
        data: {"result": {"minY": 0.8571428571428572, "minX": 7.0, "maxY": 12.312499999999998, "series": [{"data": [[7.0, 12.312499999999998], [8.0, 1.4761904761904758], [9.0, 1.3437499999999998], [10.0, 1.8627450980392162], [11.0, 1.5], [12.0, 1.4696969696969697], [13.0, 1.336842105263158], [14.0, 1.3296703296703298], [15.0, 0.9174311926605505], [16.0, 1.092198581560284], [17.0, 1.1451612903225812], [18.0, 1.395904436860068], [19.0, 1.2664399092970522], [20.0, 0.9717223650385607], [23.0, 1.25], [24.0, 1.0], [25.0, 0.8571428571428572], [26.0, 1.2149532710280373], [27.0, 1.252747252747253], [28.0, 1.201183431952664], [29.0, 1.3522727272727282], [30.0, 1.438356164383562], [31.0, 1.4123711340206182], [32.0, 1.3919597989949757], [33.0, 1.3167420814479638], [34.0, 1.4269005847953213], [35.0, 1.3886255924170607], [36.0, 1.3713080168776361], [37.0, 1.425837320574163], [38.0, 1.419642857142857], [39.0, 1.5276381909547743], [40.0, 1.3640776699029127], [41.0, 1.5097276264591437], [42.0, 1.6323529411764703], [43.0, 1.5575221238938055], [44.0, 1.6176470588235299], [45.0, 1.6744186046511627], [46.0, 1.7533632286995515], [47.0, 3.0512820512820498], [48.0, 1.797235023041475], [49.0, 1.7927461139896381], [50.0, 1.748780487804879], [51.0, 1.6497461928934014], [52.0, 1.8547854785478555], [53.0, 1.7992125984251959], [54.0, 1.7195767195767202], [55.0, 1.854166666666666], [56.0, 2.0], [57.0, 1.9727626459143974], [58.0, 1.9448529411764697], [59.0, 1.9479553903345728], [60.0, 2.0615384615384587], [61.0, 2.1936619718309855], [62.0, 2.187845303867403], [63.0, 2.3992094861660096], [64.0, 2.2043478260869573], [65.0, 2.2326388888888893], [66.0, 2.3441860465116284], [67.0, 2.4291338582677144], [68.0, 2.476377952755905], [69.0, 2.572164948453607], [70.0, 2.4312267657992566], [71.0, 2.53370786516854], [72.0, 3.3609467455621305], [73.0, 2.675182481751826], [74.0, 2.5729166666666674], [75.0, 2.631782945736436], [76.0, 2.5170068027210886], [77.0, 2.5087719298245603], [78.0, 3.034090909090909], [79.0, 2.956097560975611], [80.0, 2.9704641350210963], [81.0, 3.0731707317073185], [82.0, 2.746031746031747], [83.0, 2.7962962962962963], [84.0, 3.161290322580647], [85.0, 2.776649746192894], [86.0, 2.990476190476189], [87.0, 2.5287958115183256], [88.0, 3.049019607843135], [89.0, 3.0555555555555567], [90.0, 3.1598360655737685], [91.0, 2.8832487309644663], [92.0, 3.4028436018957344], [93.0, 5.7784810126582276], [94.0, 2.5560975609756116], [95.0, 3.2178770949720654], [96.0, 3.370860927152317], [97.0, 3.0388888888888874], [98.0, 3.411504424778763], [99.0, 3.5177304964539013], [100.0, 3.475726952128761]], "isOverall": false, "label": "HTTP 要求", "isController": false}, {"data": [[99.94966303443049, 3.474200223679119]], "isOverall": false, "label": "HTTP 要求-Aggregated", "isController": false}], "supportsControllersDiscrimination": true, "maxX": 100.0, "title": "Time VS Threads"}},
        getOptions: function() {
            return {
                series: {
                    lines: {
                        show: true
                    },
                    points: {
                        show: true
                    }
                },
                xaxis: {
                    axisLabel: "Number of active threads",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Average response times in ms",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20
                },
                legend: { noColumns: 2,show: true, container: '#legendTimeVsThreads' },
                selection: {
                    mode: 'xy'
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to work
                },
                tooltip: true,
                tooltipOpts: {
                    content: "%s: At %x.2 active threads, Average response time was %y.2 ms"
                }
            };
        },
        createGraph: function() {
            var data = this.data;
            var dataset = prepareData(data.result.series, $("#choicesTimeVsThreads"));
            var options = this.getOptions();
            prepareOptions(options, data);
            $.plot($("#flotTimesVsThreads"), dataset, options);
            // setup overview
            $.plot($("#overviewTimesVsThreads"), dataset, prepareOverviewOptions(options));
        }
};

// Time vs threads
function refreshTimeVsThreads(){
    var infos = timeVsThreadsInfos;
    prepareSeries(infos.data);
    if(infos.data.result.series.length == 0) {
        setEmptyGraph("#bodyTimeVsThreads");
        return;
    }
    if(isGraph($("#flotTimesVsThreads"))){
        infos.createGraph();
    }else{
        var choiceContainer = $("#choicesTimeVsThreads");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotTimesVsThreads", "#overviewTimesVsThreads");
        $('#footerTimeVsThreads .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
};

var bytesThroughputOverTimeInfos = {
        data : {"result": {"minY": 985717.5, "minX": 1.71086382E12, "maxY": 1.0912527833333334E7, "series": [{"data": [[1.71086382E12, 9369625.066666666], [1.7108643E12, 1.0695014266666668E7], [1.71086412E12, 1.0650419833333334E7], [1.71086394E12, 1.07419573E7], [1.71086442E12, 1676912.0666666667], [1.71086424E12, 1.0754114966666667E7], [1.71086406E12, 1.0669275766666668E7], [1.71086388E12, 1.0798542016666668E7], [1.71086436E12, 1.080776585E7], [1.71086418E12, 1.0725903716666667E7], [1.710864E12, 1.0912527833333334E7]], "isOverall": false, "label": "Bytes received per second", "isController": false}, {"data": [[1.71086382E12, 5508235.433333334], [1.7108643E12, 6287088.3], [1.71086412E12, 6261043.116666666], [1.71086394E12, 6315313.1], [1.71086442E12, 985717.5], [1.71086424E12, 6322467.116666666], [1.71086406E12, 6272534.4], [1.71086388E12, 6348108.85], [1.71086436E12, 6354065.15], [1.71086418E12, 6306739.0], [1.710864E12, 6415317.9]], "isOverall": false, "label": "Bytes sent per second", "isController": false}], "supportsControllersDiscrimination": false, "granularity": 60000, "maxX": 1.71086442E12, "title": "Bytes Throughput Over Time"}},
        getOptions : function(){
            return {
                series: {
                    lines: {
                        show: true
                    },
                    points: {
                        show: true
                    }
                },
                xaxis: {
                    mode: "time",
                    timeformat: getTimeFormat(this.data.result.granularity),
                    axisLabel: getElapsedTimeLabel(this.data.result.granularity) ,
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Bytes / sec",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                legend: {
                    noColumns: 2,
                    show: true,
                    container: '#legendBytesThroughputOverTime'
                },
                selection: {
                    mode: "xy"
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                tooltip: true,
                tooltipOpts: {
                    content: "%s at %x was %y"
                }
            };
        },
        createGraph : function() {
            var data = this.data;
            var dataset = prepareData(data.result.series, $("#choicesBytesThroughputOverTime"));
            var options = this.getOptions();
            prepareOptions(options, data);
            $.plot($("#flotBytesThroughputOverTime"), dataset, options);
            // setup overview
            $.plot($("#overviewBytesThroughputOverTime"), dataset, prepareOverviewOptions(options));
        }
};

// Bytes throughput Over Time
function refreshBytesThroughputOverTime(fixTimestamps) {
    var infos = bytesThroughputOverTimeInfos;
    prepareSeries(infos.data);
    if(fixTimestamps) {
        fixTimeStamps(infos.data.result.series, 28800000);
    }
    if(isGraph($("#flotBytesThroughputOverTime"))){
        infos.createGraph();
    }else{
        var choiceContainer = $("#choicesBytesThroughputOverTime");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotBytesThroughputOverTime", "#overviewBytesThroughputOverTime");
        $('#footerBytesThroughputOverTime .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
}

var responseTimesOverTimeInfos = {
        data: {"result": {"minY": 3.347138683853729, "minX": 1.71086382E12, "maxY": 3.520676725091535, "series": [{"data": [[1.71086382E12, 3.347138683853729], [1.7108643E12, 3.5057910371783683], [1.71086412E12, 3.520676725091535], [1.71086394E12, 3.4897292192728893], [1.71086442E12, 3.436159818415598], [1.71086424E12, 3.486225390137109], [1.71086406E12, 3.5128393345504154], [1.71086388E12, 3.4711944001624726], [1.71086436E12, 3.4693330270784593], [1.71086418E12, 3.4954137492959485], [1.710864E12, 3.4348451437931358]], "isOverall": false, "label": "HTTP 要求", "isController": false}], "supportsControllersDiscrimination": true, "granularity": 60000, "maxX": 1.71086442E12, "title": "Response Time Over Time"}},
        getOptions: function(){
            return {
                series: {
                    lines: {
                        show: true
                    },
                    points: {
                        show: true
                    }
                },
                xaxis: {
                    mode: "time",
                    timeformat: getTimeFormat(this.data.result.granularity),
                    axisLabel: getElapsedTimeLabel(this.data.result.granularity),
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Average response time in ms",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                legend: {
                    noColumns: 2,
                    show: true,
                    container: '#legendResponseTimesOverTime'
                },
                selection: {
                    mode: 'xy'
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                tooltip: true,
                tooltipOpts: {
                    content: "%s : at %x Average response time was %y ms"
                }
            };
        },
        createGraph: function() {
            var data = this.data;
            var dataset = prepareData(data.result.series, $("#choicesResponseTimesOverTime"));
            var options = this.getOptions();
            prepareOptions(options, data);
            $.plot($("#flotResponseTimesOverTime"), dataset, options);
            // setup overview
            $.plot($("#overviewResponseTimesOverTime"), dataset, prepareOverviewOptions(options));
        }
};

// Response Times Over Time
function refreshResponseTimeOverTime(fixTimestamps) {
    var infos = responseTimesOverTimeInfos;
    prepareSeries(infos.data);
    if(infos.data.result.series.length == 0) {
        setEmptyGraph("#bodyResponseTimeOverTime");
        return;
    }
    if(fixTimestamps) {
        fixTimeStamps(infos.data.result.series, 28800000);
    }
    if(isGraph($("#flotResponseTimesOverTime"))){
        infos.createGraph();
    }else{
        var choiceContainer = $("#choicesResponseTimesOverTime");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotResponseTimesOverTime", "#overviewResponseTimesOverTime");
        $('#footerResponseTimesOverTime .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
};

var latenciesOverTimeInfos = {
        data: {"result": {"minY": 3.346069422625681, "minX": 1.71086382E12, "maxY": 3.5197378151736536, "series": [{"data": [[1.71086382E12, 3.346069422625681], [1.7108643E12, 3.504864848304111], [1.71086412E12, 3.5197378151736536], [1.71086394E12, 3.488755546561544], [1.71086442E12, 3.4352579086214545], [1.71086424E12, 3.4852739220406725], [1.71086406E12, 3.511892690883191], [1.71086388E12, 3.4702246050215164], [1.71086436E12, 3.468371122577596], [1.71086418E12, 3.4945080697689446], [1.710864E12, 3.4338831928227105]], "isOverall": false, "label": "HTTP 要求", "isController": false}], "supportsControllersDiscrimination": true, "granularity": 60000, "maxX": 1.71086442E12, "title": "Latencies Over Time"}},
        getOptions: function() {
            return {
                series: {
                    lines: {
                        show: true
                    },
                    points: {
                        show: true
                    }
                },
                xaxis: {
                    mode: "time",
                    timeformat: getTimeFormat(this.data.result.granularity),
                    axisLabel: getElapsedTimeLabel(this.data.result.granularity),
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Average response latencies in ms",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                legend: {
                    noColumns: 2,
                    show: true,
                    container: '#legendLatenciesOverTime'
                },
                selection: {
                    mode: 'xy'
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                tooltip: true,
                tooltipOpts: {
                    content: "%s : at %x Average latency was %y ms"
                }
            };
        },
        createGraph: function () {
            var data = this.data;
            var dataset = prepareData(data.result.series, $("#choicesLatenciesOverTime"));
            var options = this.getOptions();
            prepareOptions(options, data);
            $.plot($("#flotLatenciesOverTime"), dataset, options);
            // setup overview
            $.plot($("#overviewLatenciesOverTime"), dataset, prepareOverviewOptions(options));
        }
};

// Latencies Over Time
function refreshLatenciesOverTime(fixTimestamps) {
    var infos = latenciesOverTimeInfos;
    prepareSeries(infos.data);
    if(infos.data.result.series.length == 0) {
        setEmptyGraph("#bodyLatenciesOverTime");
        return;
    }
    if(fixTimestamps) {
        fixTimeStamps(infos.data.result.series, 28800000);
    }
    if(isGraph($("#flotLatenciesOverTime"))) {
        infos.createGraph();
    }else {
        var choiceContainer = $("#choicesLatenciesOverTime");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotLatenciesOverTime", "#overviewLatenciesOverTime");
        $('#footerLatenciesOverTime .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
};

var connectTimeOverTimeInfos = {
        data: {"result": {"minY": 0.0, "minX": 1.71086382E12, "maxY": 1.304633196324411E-4, "series": [{"data": [[1.71086382E12, 1.304633196324411E-4], [1.7108643E12, 3.417236304247338E-5], [1.71086412E12, 3.608916508604083E-5], [1.71086394E12, 3.5192989556480877E-5], [1.71086442E12, 0.0], [1.71086424E12, 3.105160659840549E-5], [1.71086406E12, 3.070834103502385E-5], [1.71086388E12, 3.85117204002424E-5], [1.71086436E12, 3.089753851552868E-5], [1.71086418E12, 2.7605018004932283E-5], [1.710864E12, 2.9447478689553017E-5]], "isOverall": false, "label": "HTTP 要求", "isController": false}], "supportsControllersDiscrimination": true, "granularity": 60000, "maxX": 1.71086442E12, "title": "Connect Time Over Time"}},
        getOptions: function() {
            return {
                series: {
                    lines: {
                        show: true
                    },
                    points: {
                        show: true
                    }
                },
                xaxis: {
                    mode: "time",
                    timeformat: getTimeFormat(this.data.result.granularity),
                    axisLabel: getConnectTimeLabel(this.data.result.granularity),
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Average Connect Time in ms",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                legend: {
                    noColumns: 2,
                    show: true,
                    container: '#legendConnectTimeOverTime'
                },
                selection: {
                    mode: 'xy'
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                tooltip: true,
                tooltipOpts: {
                    content: "%s : at %x Average connect time was %y ms"
                }
            };
        },
        createGraph: function () {
            var data = this.data;
            var dataset = prepareData(data.result.series, $("#choicesConnectTimeOverTime"));
            var options = this.getOptions();
            prepareOptions(options, data);
            $.plot($("#flotConnectTimeOverTime"), dataset, options);
            // setup overview
            $.plot($("#overviewConnectTimeOverTime"), dataset, prepareOverviewOptions(options));
        }
};

// Connect Time Over Time
function refreshConnectTimeOverTime(fixTimestamps) {
    var infos = connectTimeOverTimeInfos;
    prepareSeries(infos.data);
    if(infos.data.result.series.length == 0) {
        setEmptyGraph("#bodyConnectTimeOverTime");
        return;
    }
    if(fixTimestamps) {
        fixTimeStamps(infos.data.result.series, 28800000);
    }
    if(isGraph($("#flotConnectTimeOverTime"))) {
        infos.createGraph();
    }else {
        var choiceContainer = $("#choicesConnectTimeOverTime");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotConnectTimeOverTime", "#overviewConnectTimeOverTime");
        $('#footerConnectTimeOverTime .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
};

var responseTimePercentilesOverTimeInfos = {
        data: {"result": {"minY": 0.0, "minX": 1.71086382E12, "maxY": 683.0, "series": [{"data": [[1.71086382E12, 186.0], [1.7108643E12, 241.0], [1.71086412E12, 204.0], [1.71086394E12, 263.0], [1.71086442E12, 48.0], [1.71086424E12, 178.0], [1.71086406E12, 197.0], [1.71086388E12, 213.0], [1.71086436E12, 177.0], [1.71086418E12, 265.0], [1.710864E12, 683.0]], "isOverall": false, "label": "Max", "isController": false}, {"data": [[1.71086382E12, 8.0], [1.7108643E12, 8.0], [1.71086412E12, 8.0], [1.71086394E12, 8.0], [1.71086442E12, 8.0], [1.71086424E12, 8.0], [1.71086406E12, 9.0], [1.71086388E12, 8.0], [1.71086436E12, 8.0], [1.71086418E12, 8.0], [1.710864E12, 8.0]], "isOverall": false, "label": "90th percentile", "isController": false}, {"data": [[1.71086382E12, 16.0], [1.7108643E12, 16.0], [1.71086412E12, 17.0], [1.71086394E12, 16.0], [1.71086442E12, 17.0], [1.71086424E12, 17.0], [1.71086406E12, 17.0], [1.71086388E12, 16.0], [1.71086436E12, 17.0], [1.71086418E12, 17.0], [1.710864E12, 16.0]], "isOverall": false, "label": "99th percentile", "isController": false}, {"data": [[1.71086382E12, 10.0], [1.7108643E12, 11.0], [1.71086412E12, 11.0], [1.71086394E12, 10.0], [1.71086442E12, 11.0], [1.71086424E12, 11.0], [1.71086406E12, 11.0], [1.71086388E12, 10.0], [1.71086436E12, 11.0], [1.71086418E12, 11.0], [1.710864E12, 11.0]], "isOverall": false, "label": "95th percentile", "isController": false}, {"data": [[1.71086382E12, 0.0], [1.7108643E12, 0.0], [1.71086412E12, 0.0], [1.71086394E12, 0.0], [1.71086442E12, 0.0], [1.71086424E12, 0.0], [1.71086406E12, 0.0], [1.71086388E12, 0.0], [1.71086436E12, 0.0], [1.71086418E12, 0.0], [1.710864E12, 0.0]], "isOverall": false, "label": "Min", "isController": false}, {"data": [[1.71086382E12, 2.0], [1.7108643E12, 3.0], [1.71086412E12, 2.0], [1.71086394E12, 3.0], [1.71086442E12, 2.0], [1.71086424E12, 2.0], [1.71086406E12, 2.0], [1.71086388E12, 2.0], [1.71086436E12, 2.0], [1.71086418E12, 3.0], [1.710864E12, 2.0]], "isOverall": false, "label": "Median", "isController": false}], "supportsControllersDiscrimination": false, "granularity": 60000, "maxX": 1.71086442E12, "title": "Response Time Percentiles Over Time (successful requests only)"}},
        getOptions: function() {
            return {
                series: {
                    lines: {
                        show: true,
                        fill: true
                    },
                    points: {
                        show: true
                    }
                },
                xaxis: {
                    mode: "time",
                    timeformat: getTimeFormat(this.data.result.granularity),
                    axisLabel: getElapsedTimeLabel(this.data.result.granularity),
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Response Time in ms",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                legend: {
                    noColumns: 2,
                    show: true,
                    container: '#legendResponseTimePercentilesOverTime'
                },
                selection: {
                    mode: 'xy'
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                tooltip: true,
                tooltipOpts: {
                    content: "%s : at %x Response time was %y ms"
                }
            };
        },
        createGraph: function () {
            var data = this.data;
            var dataset = prepareData(data.result.series, $("#choicesResponseTimePercentilesOverTime"));
            var options = this.getOptions();
            prepareOptions(options, data);
            $.plot($("#flotResponseTimePercentilesOverTime"), dataset, options);
            // setup overview
            $.plot($("#overviewResponseTimePercentilesOverTime"), dataset, prepareOverviewOptions(options));
        }
};

// Response Time Percentiles Over Time
function refreshResponseTimePercentilesOverTime(fixTimestamps) {
    var infos = responseTimePercentilesOverTimeInfos;
    prepareSeries(infos.data);
    if(fixTimestamps) {
        fixTimeStamps(infos.data.result.series, 28800000);
    }
    if(isGraph($("#flotResponseTimePercentilesOverTime"))) {
        infos.createGraph();
    }else {
        var choiceContainer = $("#choicesResponseTimePercentilesOverTime");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotResponseTimePercentilesOverTime", "#overviewResponseTimePercentilesOverTime");
        $('#footerResponseTimePercentilesOverTime .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
};


var responseTimeVsRequestInfos = {
    data: {"result": {"minY": 1.0, "minX": 6218.0, "maxY": 4.0, "series": [{"data": [[22451.0, 3.0], [23747.0, 2.0], [26627.0, 3.0], [27283.0, 3.0], [27027.0, 3.0], [28611.0, 2.0], [28163.0, 3.0], [28995.0, 2.0], [29667.0, 2.0], [29043.0, 2.0], [29139.0, 2.0], [29459.0, 2.0], [28739.0, 2.0], [29219.0, 2.0], [28755.0, 2.0], [30611.0, 2.0], [30307.0, 2.0], [30243.0, 2.0], [30291.0, 2.0], [30467.0, 2.0], [30499.0, 2.0], [30515.0, 2.0], [30179.0, 2.0], [30739.0, 2.0], [30755.0, 2.0], [17090.0, 3.0], [25746.0, 3.0], [26946.0, 3.0], [27522.0, 2.0], [28242.0, 2.0], [28322.0, 3.0], [28658.0, 2.0], [28498.0, 2.0], [28450.0, 2.0], [28210.0, 2.0], [27890.0, 3.0], [28338.0, 3.0], [29650.0, 2.0], [28882.0, 2.0], [29314.0, 2.0], [29362.0, 2.0], [29586.0, 2.0], [28962.0, 3.0], [29346.0, 2.0], [29010.0, 2.0], [29282.0, 2.0], [30210.0, 2.0], [30402.0, 2.0], [30578.0, 2.0], [29842.0, 2.0], [30914.0, 2.0], [22241.0, 3.0], [27249.0, 3.0], [27377.0, 3.0], [27425.0, 3.0], [26897.0, 3.0], [27473.0, 3.0], [28017.0, 3.0], [28593.0, 2.0], [28065.0, 2.0], [28609.0, 2.0], [27649.0, 2.0], [29681.0, 2.0], [29265.0, 2.0], [29537.0, 2.0], [29329.0, 2.0], [29345.0, 2.0], [29313.0, 2.0], [29217.0, 2.0], [29441.0, 2.0], [28769.0, 2.0], [28817.0, 3.0], [28801.0, 2.0], [29153.0, 2.0], [28753.0, 2.0], [29137.0, 3.0], [29009.0, 3.0], [30673.0, 2.0], [30401.0, 2.0], [30305.0, 2.0], [29777.0, 2.0], [30721.0, 2.0], [24720.0, 3.0], [25344.0, 3.0], [25952.0, 2.0], [26480.0, 3.0], [27504.0, 3.0], [28128.0, 3.0], [27728.0, 3.0], [27936.0, 3.0], [28192.0, 3.0], [29584.0, 2.0], [29424.0, 2.0], [29264.0, 2.0], [29568.0, 2.0], [29504.0, 2.0], [29248.0, 2.0], [29184.0, 3.0], [29216.0, 2.0], [29616.0, 2.0], [29072.0, 2.0], [28912.0, 2.0], [28880.0, 2.0], [29440.0, 2.0], [29152.0, 2.0], [28768.0, 3.0], [29088.0, 2.0], [30608.0, 2.0], [30224.0, 2.0], [30016.0, 2.0], [30208.0, 2.0], [30336.0, 2.0], [30400.0, 2.0], [30480.0, 2.0], [29808.0, 2.0], [30144.0, 2.0], [29744.0, 2.0], [29696.0, 2.0], [19927.0, 4.0], [22647.0, 3.0], [22935.0, 3.0], [26007.0, 3.0], [26071.0, 3.0], [26775.0, 3.0], [27623.0, 3.0], [26983.0, 3.0], [27687.0, 3.0], [28167.0, 2.0], [28231.0, 2.0], [29511.0, 2.0], [29159.0, 2.0], [29255.0, 2.0], [28919.0, 3.0], [28679.0, 2.0], [29351.0, 2.0], [29111.0, 2.0], [29015.0, 2.0], [28967.0, 2.0], [28951.0, 2.0], [29559.0, 2.0], [29207.0, 2.0], [29239.0, 2.0], [29527.0, 2.0], [30631.0, 2.0], [30423.0, 2.0], [30247.0, 2.0], [30087.0, 2.0], [30007.0, 2.0], [30231.0, 2.0], [30471.0, 2.0], [29847.0, 2.0], [29895.0, 2.0], [30967.0, 2.0], [22774.0, 3.0], [25318.0, 3.0], [26454.0, 3.0], [26358.0, 3.0], [26006.0, 3.0], [27014.0, 3.0], [28326.0, 3.0], [27942.0, 3.0], [27734.0, 2.0], [27926.0, 3.0], [28470.0, 3.0], [28566.0, 2.0], [28758.0, 2.0], [29030.0, 2.0], [28966.0, 2.0], [28950.0, 2.0], [29046.0, 2.0], [28710.0, 3.0], [29174.0, 2.0], [29158.0, 2.0], [29542.0, 2.0], [29510.0, 2.0], [29638.0, 2.0], [28870.0, 2.0], [29462.0, 2.0], [28838.0, 2.0], [29430.0, 2.0], [29270.0, 2.0], [29222.0, 2.0], [30070.0, 2.0], [30214.0, 2.0], [29718.0, 2.0], [22021.0, 3.0], [24405.0, 3.0], [25557.0, 3.0], [25301.0, 3.0], [26149.0, 3.0], [26037.0, 3.0], [26469.0, 3.0], [26309.0, 3.0], [27509.0, 2.0], [26693.0, 3.0], [27157.0, 3.0], [26965.0, 3.0], [27765.0, 3.0], [28213.0, 2.0], [27653.0, 3.0], [27749.0, 2.0], [27717.0, 3.0], [28053.0, 3.0], [28901.0, 2.0], [29333.0, 2.0], [29061.0, 2.0], [29141.0, 2.0], [29077.0, 2.0], [29397.0, 2.0], [29413.0, 2.0], [29205.0, 2.0], [29301.0, 2.0], [29269.0, 2.0], [29461.0, 2.0], [29573.0, 2.0], [28805.0, 2.0], [28885.0, 2.0], [29157.0, 2.0], [30309.0, 2.0], [30661.0, 2.0], [30437.0, 2.0], [30325.0, 2.0], [30341.0, 2.0], [29749.0, 2.0], [29861.0, 2.0], [30053.0, 2.0], [30965.0, 2.0], [21188.0, 3.0], [21028.0, 3.0], [23732.0, 3.0], [26292.0, 2.0], [26372.0, 3.0], [26116.0, 3.0], [25636.0, 2.0], [27060.0, 2.0], [27412.0, 3.0], [26676.0, 3.0], [27076.0, 3.0], [27540.0, 3.0], [28548.0, 3.0], [27860.0, 2.0], [28148.0, 3.0], [28324.0, 3.0], [27732.0, 3.0], [28660.0, 2.0], [28788.0, 3.0], [28756.0, 2.0], [29524.0, 2.0], [29364.0, 2.0], [29316.0, 2.0], [29492.0, 2.0], [28948.0, 2.0], [29092.0, 2.0], [28980.0, 2.0], [29460.0, 2.0], [28868.0, 2.0], [28996.0, 2.0], [30372.0, 2.0], [30436.0, 2.0], [30340.0, 2.0], [30180.0, 2.0], [30068.0, 2.0], [30196.0, 2.0], [29732.0, 2.0], [30756.0, 2.0], [31140.0, 2.0], [24459.0, 2.0], [24203.0, 3.0], [25067.0, 2.0], [25147.0, 3.0], [26043.0, 3.0], [26683.0, 3.0], [28299.0, 2.0], [28571.0, 3.0], [28635.0, 3.0], [28171.0, 2.0], [28011.0, 2.0], [28059.0, 3.0], [28843.0, 2.0], [29003.0, 2.0], [29227.0, 3.0], [29691.0, 2.0], [29659.0, 2.0], [29115.0, 3.0], [29131.0, 2.0], [29099.0, 2.0], [29019.0, 2.0], [28907.0, 2.0], [28875.0, 2.0], [29451.0, 2.0], [29579.0, 2.0], [29595.0, 2.0], [29291.0, 2.0], [29307.0, 2.0], [30523.0, 2.0], [30635.0, 2.0], [30235.0, 2.0], [29739.0, 2.0], [30059.0, 2.0], [29947.0, 2.0], [30891.0, 2.0], [30859.0, 2.0], [31259.0, 2.0], [23098.0, 2.0], [24426.0, 3.0], [24634.0, 3.0], [26538.0, 3.0], [25642.0, 3.0], [25754.0, 3.0], [26186.0, 3.0], [27466.0, 2.0], [27018.0, 2.0], [28602.0, 2.0], [28074.0, 3.0], [28618.0, 3.0], [28298.0, 2.0], [28378.0, 2.0], [28650.0, 2.0], [28522.0, 2.0], [28826.0, 2.0], [29082.0, 2.0], [29306.0, 2.0], [29242.0, 2.0], [29626.0, 2.0], [28954.0, 2.0], [29018.0, 2.0], [29530.0, 2.0], [30602.0, 2.0], [30426.0, 2.0], [30074.0, 2.0], [29962.0, 2.0], [30666.0, 2.0], [30442.0, 2.0], [30186.0, 2.0], [30554.0, 2.0], [30586.0, 2.0], [30826.0, 2.0], [24073.0, 3.0], [25193.0, 3.0], [25257.0, 2.0], [24681.0, 3.0], [25337.0, 2.0], [26057.0, 2.0], [27257.0, 3.0], [27529.0, 3.0], [28537.0, 3.0], [28633.0, 2.0], [28041.0, 3.0], [28585.0, 3.0], [28329.0, 2.5], [29273.0, 2.0], [29369.0, 2.0], [29417.0, 2.0], [29129.0, 3.0], [29145.0, 2.0], [29529.0, 2.0], [29209.0, 2.0], [29241.0, 2.0], [28713.0, 2.0], [28905.0, 3.0], [29065.0, 2.0], [29017.0, 2.0], [29097.0, 2.0], [29353.0, 2.0], [29337.0, 2.0], [29321.0, 2.0], [30441.0, 2.0], [30265.0, 2.0], [30569.0, 2.0], [29721.0, 2.0], [30121.0, 2.0], [30137.0, 2.0], [30905.0, 2.0], [30857.0, 2.0], [31033.0, 2.0], [31049.0, 2.0], [6218.0, 3.0], [26152.0, 3.0], [27096.0, 3.0], [27208.0, 3.0], [28168.0, 3.0], [28616.0, 2.0], [28120.0, 3.0], [27992.0, 2.0], [28088.0, 2.0], [27816.0, 3.0], [28504.0, 3.0], [28056.0, 3.0], [28712.0, 2.0], [29512.0, 2.0], [29416.0, 2.0], [29336.0, 2.0], [28824.0, 2.0], [29128.0, 2.0], [29016.0, 2.0], [29032.0, 2.0], [29048.0, 2.0], [29112.0, 2.0], [29176.0, 2.0], [30632.0, 2.0], [30232.0, 2.0], [30056.0, 2.0], [30376.0, 2.0], [30328.0, 2.0], [30312.0, 2.0], [29912.0, 2.0], [29832.0, 2.0], [30856.0, 2.0], [21887.0, 3.0], [22751.0, 3.0], [24543.0, 3.0], [24319.0, 3.0], [24767.0, 3.0], [24735.0, 3.0], [24623.0, 2.0], [27311.0, 3.0], [26783.0, 2.0], [27391.0, 3.0], [27519.0, 3.0], [27199.0, 3.0], [27055.0, 3.0], [27679.0, 2.0], [27999.0, 2.0], [28047.0, 3.0], [28335.0, 3.0], [28847.0, 3.0], [29599.0, 2.0], [29215.0, 2.0], [29615.0, 2.0], [29199.0, 2.0], [29487.0, 2.0], [28703.0, 2.0], [28751.0, 3.0], [29439.0, 2.0], [29279.0, 2.0], [29119.0, 2.0], [29055.0, 2.0], [30271.0, 2.0], [30127.0, 2.0], [30511.0, 2.0], [29791.0, 2.0], [30767.0, 2.0], [13271.0, 1.0], [22078.0, 3.0], [23038.0, 2.0], [23278.0, 3.0], [24334.0, 3.0], [23982.0, 3.0], [24846.0, 2.0], [25678.0, 3.0], [26110.0, 2.0], [27550.0, 3.0], [26990.0, 3.0], [28590.0, 3.0], [28222.0, 3.0], [28574.0, 2.0], [27998.0, 2.0], [28174.0, 2.0], [28670.0, 3.0], [29502.0, 2.0], [29278.0, 2.0], [29390.0, 2.0], [29358.0, 2.0], [29070.0, 2.0], [29086.0, 2.0], [29630.0, 2.0], [29550.0, 2.0], [29486.0, 2.0], [28862.0, 2.0], [29470.0, 2.0], [28782.0, 2.0], [29134.0, 2.0], [29118.0, 2.0], [29198.0, 2.0], [29246.0, 2.0], [29214.0, 2.0], [29902.0, 2.0], [30478.0, 2.0], [30542.0, 2.0], [30302.0, 2.0], [30446.0, 2.0], [29870.0, 2.0], [31166.0, 2.0], [31182.0, 2.0], [24749.0, 3.0], [25005.0, 3.0], [25117.0, 2.0], [25501.0, 3.0], [25661.0, 3.0], [26813.0, 3.0], [27389.0, 3.0], [27629.0, 3.0], [28317.0, 2.0], [28109.0, 2.0], [28461.0, 2.0], [28381.0, 2.0], [27965.0, 2.0], [28029.0, 3.0], [27693.0, 3.0], [28077.0, 3.0], [28861.0, 2.5], [29037.0, 2.0], [29357.0, 2.0], [29389.0, 2.0], [29405.0, 2.0], [29261.0, 2.0], [29693.0, 2.0], [29149.0, 2.0], [29133.0, 2.0], [29165.0, 2.0], [28941.0, 2.0], [29453.0, 2.0], [28909.0, 2.0], [30317.0, 2.0], [29949.0, 2.0], [29709.0, 2.0], [30541.0, 2.0], [30477.0, 2.0], [30029.0, 2.0], [29725.0, 2.0], [31165.0, 2.0], [30893.0, 2.0], [11230.0, 3.0], [24364.0, 3.0], [24508.0, 2.0], [24732.0, 2.0], [26668.0, 3.0], [27276.0, 2.0], [27596.0, 3.0], [28348.0, 2.0], [27660.0, 3.0], [27756.0, 3.0], [28588.0, 2.0], [27932.0, 2.0], [27820.0, 3.0], [29500.0, 2.0], [29260.0, 2.0], [29052.0, 2.0], [29244.0, 2.0], [28908.0, 2.0], [29452.0, 2.0], [29084.0, 2.0], [29388.0, 2.0], [29292.0, 2.0], [28956.0, 3.0], [28716.0, 3.0], [30428.0, 2.0], [30412.0, 2.0], [30348.0, 2.0], [29804.0, 2.0]], "isOverall": false, "label": "Successes", "isController": false}], "supportsControllersDiscrimination": false, "granularity": 1000, "maxX": 31259.0, "title": "Response Time Vs Request"}},
    getOptions: function() {
        return {
            series: {
                lines: {
                    show: false
                },
                points: {
                    show: true
                }
            },
            xaxis: {
                axisLabel: "Global number of requests per second",
                axisLabelUseCanvas: true,
                axisLabelFontSizePixels: 12,
                axisLabelFontFamily: 'Verdana, Arial',
                axisLabelPadding: 20,
            },
            yaxis: {
                axisLabel: "Median Response Time in ms",
                axisLabelUseCanvas: true,
                axisLabelFontSizePixels: 12,
                axisLabelFontFamily: 'Verdana, Arial',
                axisLabelPadding: 20,
            },
            legend: {
                noColumns: 2,
                show: true,
                container: '#legendResponseTimeVsRequest'
            },
            selection: {
                mode: 'xy'
            },
            grid: {
                hoverable: true // IMPORTANT! this is needed for tooltip to work
            },
            tooltip: true,
            tooltipOpts: {
                content: "%s : Median response time at %x req/s was %y ms"
            },
            colors: ["#9ACD32", "#FF6347"]
        };
    },
    createGraph: function () {
        var data = this.data;
        var dataset = prepareData(data.result.series, $("#choicesResponseTimeVsRequest"));
        var options = this.getOptions();
        prepareOptions(options, data);
        $.plot($("#flotResponseTimeVsRequest"), dataset, options);
        // setup overview
        $.plot($("#overviewResponseTimeVsRequest"), dataset, prepareOverviewOptions(options));

    }
};

// Response Time vs Request
function refreshResponseTimeVsRequest() {
    var infos = responseTimeVsRequestInfos;
    prepareSeries(infos.data);
    if (isGraph($("#flotResponseTimeVsRequest"))){
        infos.createGraph();
    }else{
        var choiceContainer = $("#choicesResponseTimeVsRequest");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotResponseTimeVsRequest", "#overviewResponseTimeVsRequest");
        $('#footerResponseRimeVsRequest .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
};


var latenciesVsRequestInfos = {
    data: {"result": {"minY": 1.0, "minX": 6218.0, "maxY": 4.0, "series": [{"data": [[22451.0, 3.0], [23747.0, 2.0], [26627.0, 3.0], [27283.0, 3.0], [27027.0, 3.0], [28611.0, 2.0], [28163.0, 3.0], [28995.0, 2.0], [29667.0, 2.0], [29043.0, 2.0], [29139.0, 2.0], [29459.0, 2.0], [28739.0, 2.0], [29219.0, 2.0], [28755.0, 2.0], [30611.0, 2.0], [30307.0, 2.0], [30243.0, 2.0], [30291.0, 2.0], [30467.0, 2.0], [30499.0, 2.0], [30515.0, 2.0], [30179.0, 2.0], [30739.0, 2.0], [30755.0, 2.0], [17090.0, 3.0], [25746.0, 3.0], [26946.0, 3.0], [27522.0, 2.0], [28242.0, 2.0], [28322.0, 3.0], [28658.0, 2.0], [28498.0, 2.0], [28450.0, 2.0], [28210.0, 2.0], [27890.0, 3.0], [28338.0, 3.0], [29650.0, 2.0], [28882.0, 2.0], [29314.0, 2.0], [29362.0, 2.0], [29586.0, 2.0], [28962.0, 3.0], [29346.0, 2.0], [29010.0, 2.0], [29282.0, 2.0], [30210.0, 2.0], [30402.0, 2.0], [30578.0, 2.0], [29842.0, 2.0], [30914.0, 2.0], [22241.0, 3.0], [27249.0, 3.0], [27377.0, 3.0], [27425.0, 3.0], [26897.0, 3.0], [27473.0, 3.0], [28017.0, 3.0], [28593.0, 2.0], [28065.0, 2.0], [28609.0, 2.0], [27649.0, 2.0], [29681.0, 2.0], [29265.0, 2.0], [29537.0, 2.0], [29329.0, 2.0], [29345.0, 2.0], [29313.0, 2.0], [29217.0, 2.0], [29441.0, 2.0], [28769.0, 2.0], [28817.0, 3.0], [28801.0, 2.0], [29153.0, 2.0], [28753.0, 2.0], [29137.0, 3.0], [29009.0, 3.0], [30673.0, 2.0], [30401.0, 2.0], [30305.0, 2.0], [29777.0, 2.0], [30721.0, 2.0], [24720.0, 3.0], [25344.0, 3.0], [25952.0, 2.0], [26480.0, 3.0], [27504.0, 3.0], [28128.0, 3.0], [27728.0, 3.0], [27936.0, 3.0], [28192.0, 3.0], [29584.0, 2.0], [29424.0, 2.0], [29264.0, 2.0], [29568.0, 2.0], [29504.0, 2.0], [29248.0, 2.0], [29184.0, 3.0], [29216.0, 2.0], [29616.0, 2.0], [29072.0, 2.0], [28912.0, 2.0], [28880.0, 2.0], [29440.0, 2.0], [29152.0, 2.0], [28768.0, 3.0], [29088.0, 2.0], [30608.0, 2.0], [30224.0, 2.0], [30016.0, 2.0], [30208.0, 2.0], [30336.0, 2.0], [30400.0, 2.0], [30480.0, 2.0], [29808.0, 2.0], [30144.0, 2.0], [29744.0, 2.0], [29696.0, 2.0], [19927.0, 4.0], [22647.0, 3.0], [22935.0, 3.0], [26007.0, 3.0], [26071.0, 3.0], [26775.0, 3.0], [27623.0, 3.0], [26983.0, 3.0], [27687.0, 3.0], [28167.0, 2.0], [28231.0, 2.0], [29511.0, 2.0], [29159.0, 2.0], [29255.0, 2.0], [28919.0, 3.0], [28679.0, 2.0], [29351.0, 2.0], [29111.0, 2.0], [29015.0, 2.0], [28967.0, 2.0], [28951.0, 2.0], [29559.0, 2.0], [29207.0, 2.0], [29239.0, 2.0], [29527.0, 2.0], [30631.0, 2.0], [30423.0, 2.0], [30247.0, 2.0], [30087.0, 2.0], [30007.0, 2.0], [30231.0, 2.0], [30471.0, 2.0], [29847.0, 2.0], [29895.0, 2.0], [30967.0, 2.0], [22774.0, 3.0], [25318.0, 3.0], [26454.0, 3.0], [26358.0, 3.0], [26006.0, 3.0], [27014.0, 3.0], [28326.0, 3.0], [27942.0, 3.0], [27734.0, 2.0], [27926.0, 3.0], [28470.0, 3.0], [28566.0, 2.0], [28758.0, 2.0], [29030.0, 2.0], [28966.0, 2.0], [28950.0, 2.0], [29046.0, 2.0], [28710.0, 3.0], [29174.0, 2.0], [29158.0, 2.0], [29542.0, 2.0], [29510.0, 2.0], [29638.0, 2.0], [28870.0, 2.0], [29462.0, 2.0], [28838.0, 2.0], [29430.0, 2.0], [29270.0, 2.0], [29222.0, 2.0], [30070.0, 2.0], [30214.0, 2.0], [29718.0, 2.0], [22021.0, 3.0], [24405.0, 3.0], [25557.0, 3.0], [25301.0, 3.0], [26149.0, 3.0], [26037.0, 3.0], [26469.0, 3.0], [26309.0, 3.0], [27509.0, 2.0], [26693.0, 3.0], [27157.0, 3.0], [26965.0, 3.0], [27765.0, 3.0], [28213.0, 2.0], [27653.0, 3.0], [27749.0, 2.0], [27717.0, 3.0], [28053.0, 3.0], [28901.0, 2.0], [29333.0, 2.0], [29061.0, 2.0], [29141.0, 2.0], [29077.0, 2.0], [29397.0, 2.0], [29413.0, 2.0], [29205.0, 2.0], [29301.0, 2.0], [29269.0, 2.0], [29461.0, 2.0], [29573.0, 2.0], [28805.0, 2.0], [28885.0, 2.0], [29157.0, 2.0], [30309.0, 2.0], [30661.0, 2.0], [30437.0, 2.0], [30325.0, 2.0], [30341.0, 2.0], [29749.0, 2.0], [29861.0, 2.0], [30053.0, 2.0], [30965.0, 2.0], [21188.0, 3.0], [21028.0, 3.0], [23732.0, 3.0], [26292.0, 2.0], [26372.0, 3.0], [26116.0, 3.0], [25636.0, 2.0], [27060.0, 2.0], [27412.0, 3.0], [26676.0, 3.0], [27076.0, 3.0], [27540.0, 3.0], [28548.0, 3.0], [27860.0, 2.0], [28148.0, 3.0], [28324.0, 3.0], [27732.0, 3.0], [28660.0, 2.0], [28788.0, 3.0], [28756.0, 2.0], [29524.0, 2.0], [29364.0, 2.0], [29316.0, 2.0], [29492.0, 2.0], [28948.0, 2.0], [29092.0, 2.0], [28980.0, 2.0], [29460.0, 2.0], [28868.0, 2.0], [28996.0, 2.0], [30372.0, 2.0], [30436.0, 2.0], [30340.0, 2.0], [30180.0, 2.0], [30068.0, 2.0], [30196.0, 2.0], [29732.0, 2.0], [30756.0, 2.0], [31140.0, 2.0], [24459.0, 2.0], [24203.0, 3.0], [25067.0, 2.0], [25147.0, 3.0], [26043.0, 3.0], [26683.0, 3.0], [28299.0, 2.0], [28571.0, 3.0], [28635.0, 3.0], [28171.0, 2.0], [28011.0, 2.0], [28059.0, 3.0], [28843.0, 2.0], [29003.0, 2.0], [29227.0, 3.0], [29691.0, 2.0], [29659.0, 2.0], [29115.0, 3.0], [29131.0, 2.0], [29099.0, 2.0], [29019.0, 2.0], [28907.0, 2.0], [28875.0, 2.0], [29451.0, 2.0], [29579.0, 2.0], [29595.0, 2.0], [29291.0, 2.0], [29307.0, 2.0], [30523.0, 2.0], [30635.0, 2.0], [30235.0, 2.0], [29739.0, 2.0], [30059.0, 2.0], [29947.0, 2.0], [30891.0, 2.0], [30859.0, 2.0], [31259.0, 2.0], [23098.0, 2.0], [24426.0, 3.0], [24634.0, 3.0], [26538.0, 3.0], [25642.0, 3.0], [25754.0, 3.0], [26186.0, 3.0], [27466.0, 2.0], [27018.0, 2.0], [28602.0, 2.0], [28074.0, 3.0], [28618.0, 3.0], [28298.0, 2.0], [28378.0, 2.0], [28650.0, 2.0], [28522.0, 2.0], [28826.0, 2.0], [29082.0, 2.0], [29306.0, 2.0], [29242.0, 2.0], [29626.0, 2.0], [28954.0, 2.0], [29018.0, 2.0], [29530.0, 2.0], [30602.0, 2.0], [30426.0, 2.0], [30074.0, 2.0], [29962.0, 2.0], [30666.0, 2.0], [30442.0, 2.0], [30186.0, 2.0], [30554.0, 2.0], [30586.0, 2.0], [30826.0, 2.0], [24073.0, 3.0], [25193.0, 3.0], [25257.0, 2.0], [24681.0, 3.0], [25337.0, 2.0], [26057.0, 2.0], [27257.0, 3.0], [27529.0, 3.0], [28537.0, 3.0], [28633.0, 2.0], [28041.0, 3.0], [28585.0, 3.0], [28329.0, 2.0], [29273.0, 2.0], [29369.0, 2.0], [29417.0, 2.0], [29129.0, 3.0], [29145.0, 2.0], [29529.0, 2.0], [29209.0, 2.0], [29241.0, 2.0], [28713.0, 2.0], [28905.0, 3.0], [29065.0, 2.0], [29017.0, 2.0], [29097.0, 2.0], [29353.0, 2.0], [29337.0, 2.0], [29321.0, 2.0], [30441.0, 2.0], [30265.0, 2.0], [30569.0, 2.0], [29721.0, 2.0], [30121.0, 2.0], [30137.0, 2.0], [30905.0, 2.0], [30857.0, 2.0], [31033.0, 2.0], [31049.0, 2.0], [6218.0, 3.0], [26152.0, 3.0], [27096.0, 3.0], [27208.0, 3.0], [28168.0, 3.0], [28616.0, 2.0], [28120.0, 3.0], [27992.0, 2.0], [28088.0, 2.0], [27816.0, 3.0], [28504.0, 3.0], [28056.0, 3.0], [28712.0, 2.0], [29512.0, 2.0], [29416.0, 2.0], [29336.0, 2.0], [28824.0, 2.0], [29128.0, 2.0], [29016.0, 2.0], [29032.0, 2.0], [29048.0, 2.0], [29112.0, 2.0], [29176.0, 2.0], [30632.0, 2.0], [30232.0, 2.0], [30056.0, 2.0], [30376.0, 2.0], [30328.0, 2.0], [30312.0, 2.0], [29912.0, 2.0], [29832.0, 2.0], [30856.0, 2.0], [21887.0, 3.0], [22751.0, 3.0], [24543.0, 3.0], [24319.0, 3.0], [24767.0, 3.0], [24735.0, 3.0], [24623.0, 2.0], [27311.0, 3.0], [26783.0, 2.0], [27391.0, 3.0], [27519.0, 3.0], [27199.0, 3.0], [27055.0, 3.0], [27679.0, 2.0], [27999.0, 2.0], [28047.0, 3.0], [28335.0, 3.0], [28847.0, 3.0], [29599.0, 2.0], [29215.0, 2.0], [29615.0, 2.0], [29199.0, 2.0], [29487.0, 2.0], [28703.0, 2.0], [28751.0, 3.0], [29439.0, 2.0], [29279.0, 2.0], [29119.0, 2.0], [29055.0, 2.0], [30271.0, 2.0], [30127.0, 2.0], [30511.0, 2.0], [29791.0, 2.0], [30767.0, 2.0], [13271.0, 1.0], [22078.0, 3.0], [23038.0, 2.0], [23278.0, 3.0], [24334.0, 3.0], [23982.0, 3.0], [24846.0, 2.0], [25678.0, 3.0], [26110.0, 2.0], [27550.0, 3.0], [26990.0, 3.0], [28590.0, 3.0], [28222.0, 3.0], [28574.0, 2.0], [27998.0, 2.0], [28174.0, 2.0], [28670.0, 3.0], [29502.0, 2.0], [29278.0, 2.0], [29390.0, 2.0], [29358.0, 2.0], [29070.0, 2.0], [29086.0, 2.0], [29630.0, 2.0], [29550.0, 2.0], [29486.0, 2.0], [28862.0, 2.0], [29470.0, 2.0], [28782.0, 2.0], [29134.0, 2.0], [29118.0, 2.0], [29198.0, 2.0], [29246.0, 2.0], [29214.0, 2.0], [29902.0, 2.0], [30478.0, 2.0], [30542.0, 2.0], [30302.0, 2.0], [30446.0, 2.0], [29870.0, 2.0], [31166.0, 2.0], [31182.0, 2.0], [24749.0, 3.0], [25005.0, 3.0], [25117.0, 2.0], [25501.0, 3.0], [25661.0, 3.0], [26813.0, 3.0], [27389.0, 3.0], [27629.0, 3.0], [28317.0, 2.0], [28109.0, 2.0], [28461.0, 2.0], [28381.0, 2.0], [27965.0, 2.0], [28029.0, 3.0], [27693.0, 3.0], [28077.0, 3.0], [28861.0, 2.5], [29037.0, 2.0], [29357.0, 2.0], [29389.0, 2.0], [29405.0, 2.0], [29261.0, 2.0], [29693.0, 2.0], [29149.0, 2.0], [29133.0, 2.0], [29165.0, 2.0], [28941.0, 2.0], [29453.0, 2.0], [28909.0, 2.0], [30317.0, 2.0], [29949.0, 2.0], [29709.0, 2.0], [30541.0, 2.0], [30477.0, 2.0], [30029.0, 2.0], [29725.0, 2.0], [31165.0, 2.0], [30893.0, 2.0], [11230.0, 3.0], [24364.0, 3.0], [24508.0, 2.0], [24732.0, 2.0], [26668.0, 3.0], [27276.0, 2.0], [27596.0, 3.0], [28348.0, 2.0], [27660.0, 3.0], [27756.0, 3.0], [28588.0, 2.0], [27932.0, 2.0], [27820.0, 3.0], [29500.0, 2.0], [29260.0, 2.0], [29052.0, 2.0], [29244.0, 2.0], [28908.0, 2.0], [29452.0, 2.0], [29084.0, 2.0], [29388.0, 2.0], [29292.0, 2.0], [28956.0, 3.0], [28716.0, 3.0], [30428.0, 2.0], [30412.0, 2.0], [30348.0, 2.0], [29804.0, 2.0]], "isOverall": false, "label": "Successes", "isController": false}], "supportsControllersDiscrimination": false, "granularity": 1000, "maxX": 31259.0, "title": "Latencies Vs Request"}},
    getOptions: function() {
        return{
            series: {
                lines: {
                    show: false
                },
                points: {
                    show: true
                }
            },
            xaxis: {
                axisLabel: "Global number of requests per second",
                axisLabelUseCanvas: true,
                axisLabelFontSizePixels: 12,
                axisLabelFontFamily: 'Verdana, Arial',
                axisLabelPadding: 20,
            },
            yaxis: {
                axisLabel: "Median Latency in ms",
                axisLabelUseCanvas: true,
                axisLabelFontSizePixels: 12,
                axisLabelFontFamily: 'Verdana, Arial',
                axisLabelPadding: 20,
            },
            legend: { noColumns: 2,show: true, container: '#legendLatencyVsRequest' },
            selection: {
                mode: 'xy'
            },
            grid: {
                hoverable: true // IMPORTANT! this is needed for tooltip to work
            },
            tooltip: true,
            tooltipOpts: {
                content: "%s : Median Latency time at %x req/s was %y ms"
            },
            colors: ["#9ACD32", "#FF6347"]
        };
    },
    createGraph: function () {
        var data = this.data;
        var dataset = prepareData(data.result.series, $("#choicesLatencyVsRequest"));
        var options = this.getOptions();
        prepareOptions(options, data);
        $.plot($("#flotLatenciesVsRequest"), dataset, options);
        // setup overview
        $.plot($("#overviewLatenciesVsRequest"), dataset, prepareOverviewOptions(options));
    }
};

// Latencies vs Request
function refreshLatenciesVsRequest() {
        var infos = latenciesVsRequestInfos;
        prepareSeries(infos.data);
        if(isGraph($("#flotLatenciesVsRequest"))){
            infos.createGraph();
        }else{
            var choiceContainer = $("#choicesLatencyVsRequest");
            createLegend(choiceContainer, infos);
            infos.createGraph();
            setGraphZoomable("#flotLatenciesVsRequest", "#overviewLatenciesVsRequest");
            $('#footerLatenciesVsRequest .legendColorBox > div').each(function(i){
                $(this).clone().prependTo(choiceContainer.find("li").eq(i));
            });
        }
};

var hitsPerSecondInfos = {
        data: {"result": {"minY": 4433.366666666667, "minX": 1.71086382E12, "maxY": 28864.95, "series": [{"data": [[1.71086382E12, 24785.133333333335], [1.7108643E12, 28288.033333333333], [1.71086412E12, 28170.916666666668], [1.71086394E12, 28414.75], [1.71086442E12, 4433.366666666667], [1.71086424E12, 28447.216666666667], [1.71086406E12, 28222.533333333333], [1.71086388E12, 28562.716666666667], [1.71086436E12, 28589.116666666665], [1.71086418E12, 28376.516666666666], [1.710864E12, 28864.95]], "isOverall": false, "label": "hitsPerSecond", "isController": false}], "supportsControllersDiscrimination": false, "granularity": 60000, "maxX": 1.71086442E12, "title": "Hits Per Second"}},
        getOptions: function() {
            return {
                series: {
                    lines: {
                        show: true
                    },
                    points: {
                        show: true
                    }
                },
                xaxis: {
                    mode: "time",
                    timeformat: getTimeFormat(this.data.result.granularity),
                    axisLabel: getElapsedTimeLabel(this.data.result.granularity),
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Number of hits / sec",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20
                },
                legend: {
                    noColumns: 2,
                    show: true,
                    container: "#legendHitsPerSecond"
                },
                selection: {
                    mode : 'xy'
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                tooltip: true,
                tooltipOpts: {
                    content: "%s at %x was %y.2 hits/sec"
                }
            };
        },
        createGraph: function createGraph() {
            var data = this.data;
            var dataset = prepareData(data.result.series, $("#choicesHitsPerSecond"));
            var options = this.getOptions();
            prepareOptions(options, data);
            $.plot($("#flotHitsPerSecond"), dataset, options);
            // setup overview
            $.plot($("#overviewHitsPerSecond"), dataset, prepareOverviewOptions(options));
        }
};

// Hits per second
function refreshHitsPerSecond(fixTimestamps) {
    var infos = hitsPerSecondInfos;
    prepareSeries(infos.data);
    if(fixTimestamps) {
        fixTimeStamps(infos.data.result.series, 28800000);
    }
    if (isGraph($("#flotHitsPerSecond"))){
        infos.createGraph();
    }else{
        var choiceContainer = $("#choicesHitsPerSecond");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotHitsPerSecond", "#overviewHitsPerSecond");
        $('#footerHitsPerSecond .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
}

var codesPerSecondInfos = {
        data: {"result": {"minY": 4435.033333333334, "minX": 1.71086382E12, "maxY": 28864.95, "series": [{"data": [[1.71086382E12, 24783.466666666667], [1.7108643E12, 28287.966666666667], [1.71086412E12, 28170.966666666667], [1.71086394E12, 28414.75], [1.71086442E12, 4435.033333333334], [1.71086424E12, 28447.266666666666], [1.71086406E12, 28222.516666666666], [1.71086388E12, 28562.733333333334], [1.71086436E12, 28589.116666666665], [1.71086418E12, 28376.483333333334], [1.710864E12, 28864.95]], "isOverall": false, "label": "200", "isController": false}], "supportsControllersDiscrimination": false, "granularity": 60000, "maxX": 1.71086442E12, "title": "Codes Per Second"}},
        getOptions: function(){
            return {
                series: {
                    lines: {
                        show: true
                    },
                    points: {
                        show: true
                    }
                },
                xaxis: {
                    mode: "time",
                    timeformat: getTimeFormat(this.data.result.granularity),
                    axisLabel: getElapsedTimeLabel(this.data.result.granularity),
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Number of responses / sec",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                legend: {
                    noColumns: 2,
                    show: true,
                    container: "#legendCodesPerSecond"
                },
                selection: {
                    mode: 'xy'
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                tooltip: true,
                tooltipOpts: {
                    content: "Number of Response Codes %s at %x was %y.2 responses / sec"
                }
            };
        },
    createGraph: function() {
        var data = this.data;
        var dataset = prepareData(data.result.series, $("#choicesCodesPerSecond"));
        var options = this.getOptions();
        prepareOptions(options, data);
        $.plot($("#flotCodesPerSecond"), dataset, options);
        // setup overview
        $.plot($("#overviewCodesPerSecond"), dataset, prepareOverviewOptions(options));
    }
};

// Codes per second
function refreshCodesPerSecond(fixTimestamps) {
    var infos = codesPerSecondInfos;
    prepareSeries(infos.data);
    if(fixTimestamps) {
        fixTimeStamps(infos.data.result.series, 28800000);
    }
    if(isGraph($("#flotCodesPerSecond"))){
        infos.createGraph();
    }else{
        var choiceContainer = $("#choicesCodesPerSecond");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotCodesPerSecond", "#overviewCodesPerSecond");
        $('#footerCodesPerSecond .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
};

var transactionsPerSecondInfos = {
        data: {"result": {"minY": 4435.033333333334, "minX": 1.71086382E12, "maxY": 28864.95, "series": [{"data": [[1.71086382E12, 24783.466666666667], [1.7108643E12, 28287.966666666667], [1.71086412E12, 28170.966666666667], [1.71086394E12, 28414.75], [1.71086442E12, 4435.033333333334], [1.71086424E12, 28447.266666666666], [1.71086406E12, 28222.516666666666], [1.71086388E12, 28562.733333333334], [1.71086436E12, 28589.116666666665], [1.71086418E12, 28376.483333333334], [1.710864E12, 28864.95]], "isOverall": false, "label": "HTTP 要求-success", "isController": false}], "supportsControllersDiscrimination": true, "granularity": 60000, "maxX": 1.71086442E12, "title": "Transactions Per Second"}},
        getOptions: function(){
            return {
                series: {
                    lines: {
                        show: true
                    },
                    points: {
                        show: true
                    }
                },
                xaxis: {
                    mode: "time",
                    timeformat: getTimeFormat(this.data.result.granularity),
                    axisLabel: getElapsedTimeLabel(this.data.result.granularity),
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Number of transactions / sec",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20
                },
                legend: {
                    noColumns: 2,
                    show: true,
                    container: "#legendTransactionsPerSecond"
                },
                selection: {
                    mode: 'xy'
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                tooltip: true,
                tooltipOpts: {
                    content: "%s at %x was %y transactions / sec"
                }
            };
        },
    createGraph: function () {
        var data = this.data;
        var dataset = prepareData(data.result.series, $("#choicesTransactionsPerSecond"));
        var options = this.getOptions();
        prepareOptions(options, data);
        $.plot($("#flotTransactionsPerSecond"), dataset, options);
        // setup overview
        $.plot($("#overviewTransactionsPerSecond"), dataset, prepareOverviewOptions(options));
    }
};

// Transactions per second
function refreshTransactionsPerSecond(fixTimestamps) {
    var infos = transactionsPerSecondInfos;
    prepareSeries(infos.data);
    if(infos.data.result.series.length == 0) {
        setEmptyGraph("#bodyTransactionsPerSecond");
        return;
    }
    if(fixTimestamps) {
        fixTimeStamps(infos.data.result.series, 28800000);
    }
    if(isGraph($("#flotTransactionsPerSecond"))){
        infos.createGraph();
    }else{
        var choiceContainer = $("#choicesTransactionsPerSecond");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotTransactionsPerSecond", "#overviewTransactionsPerSecond");
        $('#footerTransactionsPerSecond .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
};

var totalTPSInfos = {
        data: {"result": {"minY": 4435.033333333334, "minX": 1.71086382E12, "maxY": 28864.95, "series": [{"data": [[1.71086382E12, 24783.466666666667], [1.7108643E12, 28287.966666666667], [1.71086412E12, 28170.966666666667], [1.71086394E12, 28414.75], [1.71086442E12, 4435.033333333334], [1.71086424E12, 28447.266666666666], [1.71086406E12, 28222.516666666666], [1.71086388E12, 28562.733333333334], [1.71086436E12, 28589.116666666665], [1.71086418E12, 28376.483333333334], [1.710864E12, 28864.95]], "isOverall": false, "label": "Transaction-success", "isController": false}, {"data": [], "isOverall": false, "label": "Transaction-failure", "isController": false}], "supportsControllersDiscrimination": true, "granularity": 60000, "maxX": 1.71086442E12, "title": "Total Transactions Per Second"}},
        getOptions: function(){
            return {
                series: {
                    lines: {
                        show: true
                    },
                    points: {
                        show: true
                    }
                },
                xaxis: {
                    mode: "time",
                    timeformat: getTimeFormat(this.data.result.granularity),
                    axisLabel: getElapsedTimeLabel(this.data.result.granularity),
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20,
                },
                yaxis: {
                    axisLabel: "Number of transactions / sec",
                    axisLabelUseCanvas: true,
                    axisLabelFontSizePixels: 12,
                    axisLabelFontFamily: 'Verdana, Arial',
                    axisLabelPadding: 20
                },
                legend: {
                    noColumns: 2,
                    show: true,
                    container: "#legendTotalTPS"
                },
                selection: {
                    mode: 'xy'
                },
                grid: {
                    hoverable: true // IMPORTANT! this is needed for tooltip to
                                    // work
                },
                tooltip: true,
                tooltipOpts: {
                    content: "%s at %x was %y transactions / sec"
                },
                colors: ["#9ACD32", "#FF6347"]
            };
        },
    createGraph: function () {
        var data = this.data;
        var dataset = prepareData(data.result.series, $("#choicesTotalTPS"));
        var options = this.getOptions();
        prepareOptions(options, data);
        $.plot($("#flotTotalTPS"), dataset, options);
        // setup overview
        $.plot($("#overviewTotalTPS"), dataset, prepareOverviewOptions(options));
    }
};

// Total Transactions per second
function refreshTotalTPS(fixTimestamps) {
    var infos = totalTPSInfos;
    // We want to ignore seriesFilter
    prepareSeries(infos.data, false, true);
    if(fixTimestamps) {
        fixTimeStamps(infos.data.result.series, 28800000);
    }
    if(isGraph($("#flotTotalTPS"))){
        infos.createGraph();
    }else{
        var choiceContainer = $("#choicesTotalTPS");
        createLegend(choiceContainer, infos);
        infos.createGraph();
        setGraphZoomable("#flotTotalTPS", "#overviewTotalTPS");
        $('#footerTotalTPS .legendColorBox > div').each(function(i){
            $(this).clone().prependTo(choiceContainer.find("li").eq(i));
        });
    }
};

// Collapse the graph matching the specified DOM element depending the collapsed
// status
function collapse(elem, collapsed){
    if(collapsed){
        $(elem).parent().find(".fa-chevron-up").removeClass("fa-chevron-up").addClass("fa-chevron-down");
    } else {
        $(elem).parent().find(".fa-chevron-down").removeClass("fa-chevron-down").addClass("fa-chevron-up");
        if (elem.id == "bodyBytesThroughputOverTime") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshBytesThroughputOverTime(true);
            }
            document.location.href="#bytesThroughputOverTime";
        } else if (elem.id == "bodyLatenciesOverTime") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshLatenciesOverTime(true);
            }
            document.location.href="#latenciesOverTime";
        } else if (elem.id == "bodyCustomGraph") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshCustomGraph(true);
            }
            document.location.href="#responseCustomGraph";
        } else if (elem.id == "bodyConnectTimeOverTime") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshConnectTimeOverTime(true);
            }
            document.location.href="#connectTimeOverTime";
        } else if (elem.id == "bodyResponseTimePercentilesOverTime") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshResponseTimePercentilesOverTime(true);
            }
            document.location.href="#responseTimePercentilesOverTime";
        } else if (elem.id == "bodyResponseTimeDistribution") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshResponseTimeDistribution();
            }
            document.location.href="#responseTimeDistribution" ;
        } else if (elem.id == "bodySyntheticResponseTimeDistribution") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshSyntheticResponseTimeDistribution();
            }
            document.location.href="#syntheticResponseTimeDistribution" ;
        } else if (elem.id == "bodyActiveThreadsOverTime") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshActiveThreadsOverTime(true);
            }
            document.location.href="#activeThreadsOverTime";
        } else if (elem.id == "bodyTimeVsThreads") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshTimeVsThreads();
            }
            document.location.href="#timeVsThreads" ;
        } else if (elem.id == "bodyCodesPerSecond") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshCodesPerSecond(true);
            }
            document.location.href="#codesPerSecond";
        } else if (elem.id == "bodyTransactionsPerSecond") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshTransactionsPerSecond(true);
            }
            document.location.href="#transactionsPerSecond";
        } else if (elem.id == "bodyTotalTPS") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshTotalTPS(true);
            }
            document.location.href="#totalTPS";
        } else if (elem.id == "bodyResponseTimeVsRequest") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshResponseTimeVsRequest();
            }
            document.location.href="#responseTimeVsRequest";
        } else if (elem.id == "bodyLatenciesVsRequest") {
            if (isGraph($(elem).find('.flot-chart-content')) == false) {
                refreshLatenciesVsRequest();
            }
            document.location.href="#latencyVsRequest";
        }
    }
}

/*
 * Activates or deactivates all series of the specified graph (represented by id parameter)
 * depending on checked argument.
 */
function toggleAll(id, checked){
    var placeholder = document.getElementById(id);

    var cases = $(placeholder).find(':checkbox');
    cases.prop('checked', checked);
    $(cases).parent().children().children().toggleClass("legend-disabled", !checked);

    var choiceContainer;
    if ( id == "choicesBytesThroughputOverTime"){
        choiceContainer = $("#choicesBytesThroughputOverTime");
        refreshBytesThroughputOverTime(false);
    } else if(id == "choicesResponseTimesOverTime"){
        choiceContainer = $("#choicesResponseTimesOverTime");
        refreshResponseTimeOverTime(false);
    }else if(id == "choicesResponseCustomGraph"){
        choiceContainer = $("#choicesResponseCustomGraph");
        refreshCustomGraph(false);
    } else if ( id == "choicesLatenciesOverTime"){
        choiceContainer = $("#choicesLatenciesOverTime");
        refreshLatenciesOverTime(false);
    } else if ( id == "choicesConnectTimeOverTime"){
        choiceContainer = $("#choicesConnectTimeOverTime");
        refreshConnectTimeOverTime(false);
    } else if ( id == "choicesResponseTimePercentilesOverTime"){
        choiceContainer = $("#choicesResponseTimePercentilesOverTime");
        refreshResponseTimePercentilesOverTime(false);
    } else if ( id == "choicesResponseTimePercentiles"){
        choiceContainer = $("#choicesResponseTimePercentiles");
        refreshResponseTimePercentiles();
    } else if(id == "choicesActiveThreadsOverTime"){
        choiceContainer = $("#choicesActiveThreadsOverTime");
        refreshActiveThreadsOverTime(false);
    } else if ( id == "choicesTimeVsThreads"){
        choiceContainer = $("#choicesTimeVsThreads");
        refreshTimeVsThreads();
    } else if ( id == "choicesSyntheticResponseTimeDistribution"){
        choiceContainer = $("#choicesSyntheticResponseTimeDistribution");
        refreshSyntheticResponseTimeDistribution();
    } else if ( id == "choicesResponseTimeDistribution"){
        choiceContainer = $("#choicesResponseTimeDistribution");
        refreshResponseTimeDistribution();
    } else if ( id == "choicesHitsPerSecond"){
        choiceContainer = $("#choicesHitsPerSecond");
        refreshHitsPerSecond(false);
    } else if(id == "choicesCodesPerSecond"){
        choiceContainer = $("#choicesCodesPerSecond");
        refreshCodesPerSecond(false);
    } else if ( id == "choicesTransactionsPerSecond"){
        choiceContainer = $("#choicesTransactionsPerSecond");
        refreshTransactionsPerSecond(false);
    } else if ( id == "choicesTotalTPS"){
        choiceContainer = $("#choicesTotalTPS");
        refreshTotalTPS(false);
    } else if ( id == "choicesResponseTimeVsRequest"){
        choiceContainer = $("#choicesResponseTimeVsRequest");
        refreshResponseTimeVsRequest();
    } else if ( id == "choicesLatencyVsRequest"){
        choiceContainer = $("#choicesLatencyVsRequest");
        refreshLatenciesVsRequest();
    }
    var color = checked ? "black" : "#818181";
    if(choiceContainer != null) {
        choiceContainer.find("label").each(function(){
            this.style.color = color;
        });
    }
}

