async function set(id) {
    console.log(id);
    const res = await fetch(`/api/mode/${id}`)
    const response = await res.json()
    return response
}