var operationLock = false;
var currentMode;

async function addStations(modes) {
  const ids = getIds(modes);
  for (let mode of modes) {
    // console.log(mode);
    const parentDiv = document.getElementById("mode-selector-div");

    const nodeItem = document.createElement("div");
    nodeItem.className = "mode-item station threeDp";
    nodeItem.id = mode.Id;

    const nodeItemLeft = document.createElement("div");
    nodeItemLeft.className = "node-item-left";

    const nodeItemLeftPicture = document.createElement("div");
    nodeItemLeftPicture.className = "node-item-picture";
    nodeItemLeftPicture.id = `${mode.Id}-picture`;
    nodeItemLeftPicture.style.backgroundImage = `url(${mode.ImagePath})`;

    const nodeItemLabels = document.createElement("div");
    nodeItemLabels.className = "node-item-labels";

    const nodeItemTitle = document.createElement("h2");
    nodeItemTitle.innerText = mode.Name;
    const nodeItemSubTitle = document.createElement("h3");
    nodeItemSubTitle.innerText = mode.Description;

    parentDiv.appendChild(nodeItem);
    nodeItem.appendChild(nodeItemLeft);
    nodeItemLeft.appendChild(nodeItemLeftPicture);
    nodeItemLeft.appendChild(nodeItemLabels);
    nodeItemLabels.appendChild(nodeItemTitle);
    nodeItemLabels.appendChild(nodeItemSubTitle);

    // <div class="spinner" id="spinner"></div>
    const spinner = document.createElement("div");
    spinner.className = "spinner";
    spinner.id = `${mode.Id}-spinner`;

    nodeItem.appendChild(spinner);

    setTimeout(function () {
      nodeItem.style.opacity = "1";
    }, 50);

    nodeItem.addEventListener("click", async function () {
      if (operationLock) return;
      modeBefore = currentMode;
      operationLock = true;

      for (let item of ids) setSmall(item);

      startSpinner(`${mode.Id}-spinner`);
      applyActiveCss(nodeItem, nodeItemLeftPicture);

      const response = await setCurrentMode(mode.Id);
      await stopSpinner(`${mode.Id}-spinner`);
      console.log(response);
      if (response.Success) {
        nodeItem.style.boxShadow = "0 0 0 3px var(--clr-success)";
        setTimeout(function () {
          nodeItem.style.boxShadow = "0 0 0 3px var(--clr-purple)";
          nodeItem.style.transform = "scale(1.05)";
          operationLock = false;
        }, 1000);
      } else {
        nodeItem.style.boxShadow = "0 0 0 3px var(--clr-error)";
        setTimeout(function () {
          nodeItem.style.boxShadow = "none";
          operationLock = false;
          for (let item of ids) setSmall(item);
          const node = document.getElementById(modeBefore);
          const ItemLeftPicture = document.getElementById(
            `${modeBefore}-picture`
          );
          applyActiveCss(node, ItemLeftPicture);
          node.style.transform = "scale(1.05)";
        }, 1000);
      }
    });
  }
}

function applyActiveCss(node, picture) {
  picture.style.filter = "grayscale(0)";
  node.style.filter = "brightness(100%)";
  node.style.boxShadow = "0 0 0 3px var(--clr-purple)";
}

function setSmall(id) {
  const item = document.getElementById(id);
  item.style.transform = "scale(1)";
  item.style.filter = "brightness(60%)";
  item.style.boxShadow = "none";

  const picture = document.getElementById(`${id}-picture`);
  picture.style.filter = "grayscale(30%)";
}

async function mainloop(modes) {
  currentMode = await getCurrentMode();
  setTimeout(function () {
    setCurrentModeGui(currentMode, modes);
  }, 100);
  setTimeout(function () {
    mainloop(modes).then()
  }, 10000)
}

window.onload = async function () {
  const modes = await getAvailableModes();
  await addStations(modes);
  mainloop(modes).then()
};
