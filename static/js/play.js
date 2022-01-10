async function set(id) {
    console.log(id);
    const res = await fetch(`/play/${id}`)
    const stations = await res.json()
}