window.onload = async () => {
    const version = await getVersion();
    setVersion(version.Version, version.Name, version.Production);
  };
  