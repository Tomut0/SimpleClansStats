const baseClanUrl = "/detail/clan/",
    basePlayerUrl = "/detail/player/";

function setupModalHolder() {
    let modalHolder = document.getElementById("modalHolder");
    if (!modalHolder) {
        modalHolder = document.createElement("div");
        modalHolder.id = "modalHolder";
        document.body.appendChild(modalHolder);
    }
}

function getModal() {
    let element = document.querySelector('#detailModal');
    if (!element) {
        throw new Error("Modal not loaded");
    }
    return new bootstrap.Modal(element);
}

function removeBackdrop() {
    let backdrop = document.body.querySelector(".modal-backdrop.fade.show");
    if (backdrop) {
        document.body.removeChild(backdrop);
    }
}

function showModal(url, id) {
    if (!id) {
        console.log("invalid id")
        return;
    }
    //when opening a modal from another, the backdrop doesn't disappear, so this takes care of it
    removeBackdrop();
    setupModalHolder();
    fetch(url + id).then(function (response) {
        return response.text();
    }).then(function (text) {
        document.querySelector('#modalHolder').innerHTML = text;
        getModal().show();
        //registering again, in case the modal has modal-openers
        registerModalOpeners();
    });
}

function showPlayer(name) {
    showModal(basePlayerUrl, name);
}

function showClan(tag) {
    showModal(baseClanUrl, tag);
}

function registerModalOpeners() {
    let elements = document.querySelectorAll(".modal-opener");
    for (let element of elements) {
        let tag = element.dataset.tag;
        let nick = element.dataset.nick;
        if (tag) {
            element.onclick = () => showClan(tag)
        } else if (nick) {
            element.onclick = () => showPlayer(nick);
        }
    }
}

registerModalOpeners();
