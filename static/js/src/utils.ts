interface Mode {
  Name: string;
  Description: string;
  ImagePath: string;
  Url: string;
  Id: number;
  Volume: number;
};

function getIds(modes: Mode[]) {
  let ids = [];
  for (let mode of modes) {
    ids.push(mode.Id);
  }
  return ids;
}

function setCurrentModeGui(modeId: string, modes: Mode[]) {
  const ids = getIds(modes);
  for (let item of ids) {
    setSmall(`${item}`);
  }

  const modeDiv: HTMLElement = document.getElementById(modeId) as HTMLElement;
  modeDiv.style.transform = "scale(1.05)";
  modeDiv.style.filter = "brightness(100%)";
  modeDiv.style.boxShadow = "0 0 0 3px var(--clr-purple)";

  const picture: HTMLElement = document.getElementById(
    `${modeId}-picture`
  ) as HTMLElement;
  picture.style.filter = "grayscale(0)";
}
