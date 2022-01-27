function setVersion(version: string, production: boolean) {
  const versionIndicator: HTMLSpanElement = document.getElementById(
    "navbar-application-version"
  ) as HTMLSpanElement;
  versionIndicator.innerText = `@${version}`;
  if (!production) {
    versionIndicator.style.color = "rgb(255 175 175)";
    const productionIndicator: HTMLSpanElement = document.getElementById(
      "navbar-application-production"
    ) as HTMLSpanElement;
    productionIndicator.innerText = "Development Server";
    productionIndicator.style.marginLeft = "1rem";
    productionIndicator.style.color = "rgb(255 175 175)";
  }
}

async function getVersion() {
  const url = "/api/version";
  const res = await fetch(url);
  return await res.json();
  // {"Version": string, "Production": bool}
}
