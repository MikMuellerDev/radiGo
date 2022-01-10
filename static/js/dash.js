const getStationUrl = "/stations"

async function fetchStations() {
    const res = await fetch(getStationUrl)
    const stations = await res.json()
    return stations["Stations"]
}

async function addStations() {
    const stations = await fetchStations()
    for (let station of stations) {
        console.log(station);
        const parentDiv = document.getElementById("mode-selector-div")

        const nodeItem = document.createElement("div")
        nodeItem.className = 'mode-item station threeDp'

        const nodeItemLeft = document.createElement("div")
        nodeItemLeft.className = 'node-item-left'

        const nodeItemLeftPicture = document.createElement("div")
        nodeItemLeftPicture.className = 'node-item-picture'
        nodeItemLeftPicture.style.backgroundImage = `url(${station.ImagePath})`

        const nodeItemLabels = document.createElement("div")
        nodeItemLabels.className = 'node-item-labels'

        const nodeItemSelect = document.createElement("div")
        nodeItemSelect.className = 'node-item-select-indicator'

        const nodeItemTitle = document.createElement("h2")
        nodeItemTitle.innerText = station.Name
        const nodeItemSubTitle = document.createElement("h3")
        nodeItemSubTitle.innerText = station.Description

        parentDiv.appendChild(nodeItem)
        nodeItem.appendChild(nodeItemLeft)
        nodeItemLeft.appendChild(nodeItemLeftPicture)
        nodeItemLeft.appendChild(nodeItemLabels)
        nodeItemLabels.appendChild(nodeItemTitle)
        nodeItemLabels.appendChild(nodeItemSubTitle)
        nodeItem.appendChild(nodeItemSelect)

        nodeItem.addEventListener('click', function () {
            set(station.Id)
        })
    }
}

window.onload = function () {
    addStations().then()
}

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