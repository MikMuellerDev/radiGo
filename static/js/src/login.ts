window.onload = async function () {
  const version = await getVersion();
  setVersion(version.Version, version.Name, version.Production);
};
