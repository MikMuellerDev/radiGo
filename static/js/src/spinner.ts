function startSpinner(spinnerId: string) {
  const spinner: HTMLElement = document.getElementById(
    spinnerId
  ) as HTMLElement;
  if (spinner.style.opacity !== "100%")
    if (spinner.style.animation !== "spinnerAnimation 0.9s linear infinite")
      spinner.style.animation = "spinnerAnimation 0.9s linear infinite";
  spinner.style.opacity = "100%";
}

function stopSpinner(spinnerId: string) {
  const spinner: HTMLElement = document.getElementById(
    spinnerId
  ) as HTMLElement;
  if (spinner.style.opacity === "1") {
    setTimeout(function () {
      spinner.style.opacity = "0";
      setTimeout(function () {
        spinner.style.animation = "none";
        spinner.style.opacity = "none";
      }, 200);
    }, 800);
  }
}
