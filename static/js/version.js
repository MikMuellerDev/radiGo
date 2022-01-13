function setVersion(version, production) {
  console.log(version, production);
  const versionIndicator = document.getElementById(
    "navbar-application-version"
  );
  versionIndicator.innerText = `@${version}`;
  if (!production) {
    versionIndicator.style.color = "rgb(255 175 175)"

    const productionIndicator = document.getElementById("navbar-application-production")
    productionIndicator.innerText = "DEVELOPMENT SERVER"
    productionIndicator.style.marginLeft = "1rem"
    productionIndicator.style.color = "rgb(255 175 175)"
  }
}



async function getVersion() {
  const url = "/api/version";
  const res = await fetch(url);
  return await res.json();
  // {"Version": string, "Production": bool}
}