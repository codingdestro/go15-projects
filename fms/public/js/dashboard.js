(() => {
  const token = localStorage.getItem("token");
  if (!token) {
    alert("token not found");
    return;
  }
  fetchFolders(token);
})();

function fetchFolders(token) {
  fetch("/api/user/folders", {
    headers: {
      "Content-Type": "application/json",
      Authorization: token,
    },
    method: "POST",
  })
    .then((e) => e.json())
    .then((res) => {
      const folders = document.querySelector(".folder-structure");
      for (let i = 0; i < res.folders.length; i++) {
        const folder = createFolder(res.folders[i].id, res.folders[i].name);
        folders.append(folder);
      }
      console.log(folders);
    })
    .catch((e) => console.log(e));
}

function createFolder(id, name) {
  const folderDiv = document.createElement("div");
  folderDiv.className = "folder";
  folderDiv.id = id;

  const img = document.createElement("img");
  img.src = "/icons/folder.png";
  img.className = "icon";

  const namePara = document.createElement("p");
  namePara.textContent = name;

  folderDiv.appendChild(img);
  folderDiv.appendChild(namePara);

  return folderDiv;
}
