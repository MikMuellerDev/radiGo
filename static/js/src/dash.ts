let operationLock: boolean = false;
let currentMode: string;
let modeBefore: string;

async function addStations(modes: any) {
  const ids = getIds(modes);
  for (let mode of modes) {
    const parentDiv: HTMLElement = document.getElementById(
      "mode-selector-div"
    ) as HTMLElement;

    const nodeItem = document.createElement("div");
    nodeItem.className = "mode-item station threeDp";
    nodeItem.id = mode.Id;

    const nodeItemLeft = document.createElement("div");
    nodeItemLeft.className = "node-item-left";

    const nodeItemLeftPicture: HTMLPictureElement =
      document.createElement("div");
    nodeItemLeftPicture.className = "node-item-picture";
    nodeItemLeftPicture.id = `${mode.Id}-picture`;
    nodeItemLeftPicture.style.backgroundImage = `url(${mode.ImagePath})`;

    const nodeItemLabels: HTMLDivElement = document.createElement("div");
    nodeItemLabels.className = "node-item-labels";

    const nodeItemTitle: HTMLHeadingElement = document.createElement("h2");
    nodeItemTitle.innerText = mode.Name;
    const nodeItemSubTitle: HTMLHeadingElement = document.createElement("h3");
    nodeItemSubTitle.innerText = mode.Description;

    parentDiv.appendChild(nodeItem);
    nodeItem.appendChild(nodeItemLeft);
    nodeItemLeft.appendChild(nodeItemLeftPicture);
    nodeItemLeft.appendChild(nodeItemLabels);
    nodeItemLabels.appendChild(nodeItemTitle);
    nodeItemLabels.appendChild(nodeItemSubTitle);

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

      for (let item of ids) setSmall(`${item}`);

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
          for (let item of ids) setSmall(`${item}`);
          const node: HTMLElement = document.getElementById(
            modeBefore
          ) as HTMLElement;
          const ItemLeftPicture: HTMLElement = document.getElementById(
            `${modeBefore}-picture`
          ) as HTMLElement;
          applyActiveCss(node, ItemLeftPicture);
          node.style.transform = "scale(1.05)";
        }, 1000);
      }
    });
  }
}

function applyActiveCss(node: HTMLElement, picture: HTMLElement) {
  picture.style.filter = "grayscale(0)";
  node.style.filter = "brightness(100%)";
  node.style.boxShadow = "0 0 0 3px var(--clr-purple)";
}

function setSmall(id: string) {
  const item: HTMLElement = document.getElementById(id) as HTMLElement;
  item.style.transform = "scale(1)";
  item.style.filter = "brightness(60%)";
  item.style.boxShadow = "none";

  const picture: HTMLElement = document.getElementById(
    `${id}-picture`
  ) as HTMLElement;
  picture.style.filter = "grayscale(30%)";
}

async function mainloop(modes: [key: Mode]) {
  currentMode = await getCurrentMode(true);
  setTimeout(function () {
    setCurrentModeGui(currentMode, modes);
  }, 100);
  setTimeout(function () {
    mainloop(modes).then();
  }, 10000);
}

window.onload = async function () {
  currentMode = await getCurrentMode(false);
  const modes = await getAvailableModes();
  const version = await getVersion();
  await addStations(modes);
  setVersion(version.Version, version.Production);
  mainloop(modes).then();
};
