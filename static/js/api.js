async function setCurrentMode(id) {
  try {
    const res = await fetch(`/api/mode/${id}`, { method: "post" });
    return await res.json();
  } catch {
    return { Success: false, ErrorCode: -1, Title: "", Message: "" };
  }
}

async function getAvailableModes() {
  const url = "/api/mode/list";
  const res = await fetch(url);
  return (await res.json())["Modes"];
}

async function getCurrentMode(useKeepalive) {
  let url = "/api/mode";
  if (useKeepalive) url = "/api/mode/keepalive";
  const res = await fetch(url);
  return (await res.json())["Mode"];
}
