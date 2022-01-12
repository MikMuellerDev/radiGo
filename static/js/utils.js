function getIds(modes) {
  let ids = [];
  for (let mode of modes) {
    ids.push(mode.Id);
  }
  return ids;
}

function setCurrentModeGui(modeId, modes) {
  const ids = getIds(modes);
  for (let item of ids) {
    setSmall(item);
  }

  const modeDiv = document.getElementById(modeId);
  modeDiv.style.transform = "scale(1.05)";
  modeDiv.style.filter = "brightness(100%)";
  modeDiv.style.boxShadow = "0 0 0 3px var(--clr-purple)";

  const picture = document.getElementById(`${modeId}-picture`);
  picture.style.filter = "grayscale(0)";
}
