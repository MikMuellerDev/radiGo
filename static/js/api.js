async function setCurrentMode(id) {
  const res = await fetch(`/api/mode/${id}`, {
    method: "post",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
  });
  const response = await res.json();
  return response;
}

async function getAvailableModes() {
  const url = "/api/mode/list";
  const res = await fetch(url);
  return (await res.json())["Modes"];
}

async function getCurrentMode() {
  const url = "/api/mode";
  const res = await fetch(url);
  return (await res.json())["Mode"];
}
