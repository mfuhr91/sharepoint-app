/*window.onload = () => {
    checkQuantityUpdatePrice()
    /!*validateForms()*!/
}*/

/*const searchProduct = async (customerId) => {
    const itemList = getItemList()

    const options = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(itemList),
    };
    const url = "/products?search=true&customerId=" + customerId

    await fetch(url, options)
    location.href = url;
}

const saveFile = async (customerId) => {

    const finalPrice = document.querySelector("#finalPrice").textContent
    const orderId = document.getElementsByName("id")[0].value

    const itemList = getItemList()

    const customer = {
        id: customerId,
    }
    const order = {
        id: orderId,
        customer: customer,
        items: itemList,
        price: finalPrice,
        time: new Date().toISOString()
    }

    const options = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(order),
    };

    await fetch("/orders/save", options)
    location.href = "/orders"
}*/

/*
const checkQuantityUpdatePrice = () => {
    const list = getItemList()
    const finalPriceElem = document.querySelector("#finalPrice")
    let saveButtonDisabled = true;

    let saveOrderBtn = document.querySelector("#saveOrderBtn")

    let finalPrice = 0
    let errors = false
    if (finalPriceElem != null) {
        finalPriceElem.textContent = finalPrice.toString()
    }
    for (const i in list) {

        let span = document.querySelector("#error-" + list[i].index)

        finalPrice += list[i].price * list[i].quantity

        finalPriceElem.textContent = finalPrice.toString()

        if ((Number(list[i].quantity) > Number(list[i].stock)) ||
            (Number(list[i].quantity) <= 0 || list[i].quantity === "")) {
            span.classList.remove("d-none");
            span.classList.add("animate__headShake");
            finalPriceElem.textContent = "0";
            saveButtonDisabled = true
            errors = true
        } else {
            if (span != null) {
                span.classList.add("d-none");
                span.classList.remove("animate__headShake");
            }
        }

        if (!errors) {
            saveButtonDisabled = false
        }
    }
    if (saveOrderBtn != null) {

        if (saveButtonDisabled) {
            saveOrderBtn.setAttribute("disabled", "")
        } else {
            saveOrderBtn.removeAttribute("disabled")

        }
    }

}*/

const setModalValues = (file) => {
    console.log(file)
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
    const newTime = new Date(arrTime[1]).toLocaleTimeString("es-AR")
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
    modalImg.setAttribute("src","none")

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
    const form = document.getElementById("form")
    const spanError = document.querySelector("#error")

    let saveButtonDisabled = false;
    if (form != null) {
        const inputs = form.getElementsByTagName("input")
        /*for (let i in inputs) {
            if (i > 0) {
                if (inputs[i].value.length === 0 ||
                    (i <= maxIndexToCheckHasNumber && hasNumber(inputs[i].value))) {
                    console.log(hasNumber(inputs[i].value))
                    saveButtonDisabled = true
                    spanError.classList.remove("d-none");
                    spanError.classList.add("animate__headShake");
                    break
                } else {
                    if (spanError != null) {
                        spanError.classList.add("d-none");
                        spanError.classList.remove("animate__headShake");
                    }
                }
                saveButtonDisabled = false
            }
        }*/
    }

    if (saveBtn != null) {
        if (saveButtonDisabled) {
            saveBtn.setAttribute("disabled", "")
        } else {
            saveBtn.removeAttribute("disabled")

        }
    }
}

const eliminarFile = (element) => {

    let itemCard = element.closest(".card")

    itemCard.remove()
    checkQuantityUpdatePrice()
    refreshItems()
}

const refreshItems = () => {
    const itemsDivs = document.querySelectorAll(".item-info")
    itemsDivs.forEach((item, key) => {
        item.nextElementSibling.setAttribute("id", "error-" + key)
    })
}

const setBlur = () => {
    const container = document.getElementsByClassName("container")
    console.log(body)
    //container[0].setAttribute("style":"blur(1002px);")
}
/*const dropDown = (id) => {
    console.log(id)
    const dropdown = document.getElementsByClassName("dropdown")

    dropdown.setAttribute("aria-expanded", true)
}*/
