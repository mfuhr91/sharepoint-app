
const setModalValues = (file) => {
    removeModalValues()

    const id = file.id
    const name = file.name
    const ext = file.ext
    const size = file.size
    const icon = file.icon
    const url = file.url
    const time = file.time

    let arrTime = time.split(" ")

    const modalLabel = document.getElementById("viewModalLabel")
    const modalImg = document.getElementById("modalImg")
    const modalAudio = document.getElementById("modalAudio")
    const modalVideo = document.getElementById("modalVideo")
    const modalEmbed = document.getElementById("modalEmbed")
    const modalFooter = document.getElementById("modalFooter")

    modalLabel.textContent = name + "." + ext

    if (icon.includes("image")) {
        modalImg.classList.remove("d-none")
        modalImg.setAttribute("src", url)
    }

    if (icon.includes("audio")) {
        modalAudio.classList.remove("d-none")
        modalAudio.setAttribute("src", url)
        modalAudio.setAttribute("type", "audio/" + ext)
    }
    if (icon.includes("video")) {
        modalVideo.classList.remove("d-none")
        modalVideo.setAttribute("src", url)
        modalVideo.setAttribute("type", "video/" + ext)
    }

    if (!icon.includes("video") &&
        !icon.includes("image") &&
        !icon.includes("audio")) {
        modalEmbed.classList.remove("d-none")
        modalEmbed.setAttribute("src", url)
        modalEmbed.setAttribute("type", "application/" + ext)
    }


    const date = new Date(arrTime[0]).toLocaleDateString("es-AR")
    modalFooter.textContent = arrTime[1].split(".")[0] + " " + date
}

const removeModalValues = () => {

    const modalLabel = document.getElementById("viewModalLabel")
    const modalImg = document.getElementById("modalImg")
    const modalAudio = document.getElementById("modalAudio")
    const modalVideo = document.getElementById("modalVideo")
    const modalEmbed = document.getElementById("modalEmbed")
    const modalFooter = document.getElementById("modalFooter")

    modalLabel.textContent = ""
    modalFooter.textContent = ""

    modalImg.classList.add("d-none")
    modalImg.setAttribute("src", "none")

    modalAudio.classList.add("d-none")
    modalAudio.setAttribute("src", "none")
    modalAudio.setAttribute("type", "audio/mp3")

    modalVideo.classList.add("d-none")
    modalVideo.setAttribute("src", "none")
    modalVideo.setAttribute("type", "video/mp4")

    modalEmbed.classList.add("d-none")
    modalEmbed.setAttribute("src", "none")
    modalEmbed.setAttribute("type", "application/pdf")
}

const validateFile = () => {
    const saveBtn = document.getElementById("saveBtn")
    const inputFileLabel = document.getElementById("inputFileLabel")
    const form = document.getElementById("form")
    const divError = document.querySelector("#error")

    let saveButtonDisabled = false;
    let inputFile
    if (form != null) {
        inputFile = document.getElementById("inputFile")
        if (inputFile.value == null) {
            saveButtonDisabled = true
        }
        let fileNameArr = inputFile.value.split("fakepath\\")
        inputFileLabel.textContent = fileNameArr[1]

        const fileSize = inputFile.files[0].size
        if (fileSize > 10485760) { // 10MB
            saveButtonDisabled = true
            divError.classList.add("animate__headShake")
            divError.children[0].classList.remove("bg-warning")
            divError.children[0].classList.add("bg-danger")
            divError.children[0].classList.add("text-white")
            divError.children[0].classList.remove("text-dark")

        } else {
            divError.classList.remove("animate__headShake")
            divError.children[0].classList.add("bg-warning")
            divError.children[0].classList.remove("bg-danger")
            divError.children[0].classList.remove("text-white")
            divError.children[0].classList.add("text-dark")


        }
    }

    if (saveBtn != null) {
        if (saveButtonDisabled) {
            saveBtn.setAttribute("disabled", "")
        } else {
            saveBtn.removeAttribute("disabled")

        }
    }
}

const setBlur = () => {
    const title = document.getElementById("title")
    const backdropDiv = document.getElementById("backdropDiv")

    title.setAttribute("style", "animation-name: blurOn;z-index: 2;")
    backdropDiv.setAttribute("style", "animation-name: blurOn;z-index: 2;")
}

const removeBlur = () => {
    const backdropDiv = document.getElementById("backdropDiv")

    backdropDiv.setAttribute("style", "animation-name: blurOut; z-index: -1; transition: z-index linear 1s;")
}

const hideAddBtn = () => {
    const addBtn = document.getElementById("addBtn")
    addBtn.classList.add("d-none")
}

const cleanInputFile = () => {
    hideAddBtn()
    const saveBtn = document.getElementById("saveBtn")
    const inputFile = document.getElementById("inputFile")
    const inputFileLabel = document.getElementById("inputFileLabel")
    const divError = document.querySelector("#error")

    saveBtn.setAttribute("disabled", "")
    inputFile.value = ""
    inputFileLabel.textContent = "Subir un archivo"

    divError.classList.remove("animate__headShake")
    divError.children[0].classList.add("bg-warning")
    divError.children[0].classList.remove("bg-danger")
    divError.children[0].classList.add("text-dark")
    divError.children[0].classList.remove("text-white")

}

const showAddBtn = () => {
    const addBtn = document.getElementById("addBtn")

    setTimeout(() => {
        addBtn.classList.remove("d-none")
    },300)
    console.log("remove")
}
