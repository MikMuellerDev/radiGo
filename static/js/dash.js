const getStationUrl = "/stations";

async function fetchStations() {
  const res = await fetch(getStationUrl);
  const stations = await res.json();
  return stations["Stations"];
}

function getIds(stations) {
  let ids = [];
  for (let station of stations) {
    ids.push(station.Id);
  }
  return ids;
}

async function addStations() {
  const stations = await fetchStations();
  const ids = getIds(stations);
  for (let station of stations) {
    console.log(station);
    const parentDiv = document.getElementById("mode-selector-div");

    const nodeItem = document.createElement("div");
    nodeItem.className = "mode-item station threeDp";
    nodeItem.id = station.Id;

    const nodeItemLeft = document.createElement("div");
    nodeItemLeft.className = "node-item-left";

    const nodeItemLeftPicture = document.createElement("div");
    nodeItemLeftPicture.className = "node-item-picture";
    nodeItemLeftPicture.style.backgroundImage = `url(${station.ImagePath})`;

    const nodeItemLabels = document.createElement("div");
    nodeItemLabels.className = "node-item-labels";

    const nodeItemTitle = document.createElement("h2");
    nodeItemTitle.innerText = station.Name;
    const nodeItemSubTitle = document.createElement("h3");
    nodeItemSubTitle.innerText = station.Description;

    parentDiv.appendChild(nodeItem);
    nodeItem.appendChild(nodeItemLeft);
    nodeItemLeft.appendChild(nodeItemLeftPicture);
    nodeItemLeft.appendChild(nodeItemLabels);
    nodeItemLabels.appendChild(nodeItemTitle);
    nodeItemLabels.appendChild(nodeItemSubTitle);
    

    setTimeout(function () {
      nodeItem.style.opacity = "1";
    }, 50)
    
    nodeItem.addEventListener("click", async function () {
      for (let item of ids) {
        setSmall(item);
      }
      nodeItem.style.transform = "scale(1.05)";
      nodeItem.style.filter = "brightness(100%)";
      nodeItem.style.border = "3px solid var(--clr-purple)";

      

      startSpinner("spinner")
      nodeItem.style.animation = 'setMode 1s linear infinite'
      const response = await set(station.Id);
      nodeItem.style.animation = 'none'
      console.log(response);
      stopSpinner("spinner")
    });
  }
}

function setSmall(id) {
  const item = document.getElementById(id);
  item.style.transform = "scale(1)";
  item.style.filter = "brightness(60%)";
  item.style.border = "3px solid transparent";
}

window.onload = function () {
  addStations().then();
};

// <div class="mode-item station threeDp">
//           <div class="node-item-left">
//             <div class="node-item-picture"></div>
//             <div class="node-item-labels">
//               <h2>Title</h2>
//               <h3>An interesting subtitle</h3>
//             </div>
//           </div>
//           <div class="node-item-select-indicator"></div>
//         </div>
