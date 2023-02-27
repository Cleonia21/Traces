// const update = document.getElementById("update")
// const word_table = document.getElementById("word_table")

const word_button = document.getElementById("word_button")

/*
const a = document.getElementById("result_line_0")

a.getElementsByClassName("result_elem")[0].textContent = "Daaaa"

 */

const lines_answers = []
// let research

function ChangeColor(type, elem) {
    let line = document.getElementById("result_line_text_" + elem.toString())
    let color
    if (type === "+") {
        lines_answers[elem] = true
        color = "#0dd210";
    }
    else if (type === "-") {
        lines_answers[elem] = false
        color = "#d20d0d";
    }
    line.style.backgroundColor = color;
    console.log(lines_answers)
}

word_button.onclick = async function () {
    await get_words(3)
}

async function get_words(len) {
    let param = { len: len }
    let response = await fetch("/get_words", {
        method: "POST",
        body: JSON.stringify(param)
    });

    if (response.ok) {
        response.text().then(function (data) {
            let result = JSON.parse(data);

            if (result.Result === "ok") {
                for (let i = 0; i < result.Data.length; i++) {
                    if (result.Data[i] == null)
                        break

                    const line = document.getElementById("result_line_" + i.toString())
                    let text = line.getElementsByClassName("result_line_text")
                    text[0].style.backgroundColor = '#bdb491';
                    while (text[0].firstChild) {
                        text[0].removeChild(text[0].firstChild)
                    }

                    for (let j = 0; j < result.Data[i].length; j++) {
                        if (result.Data[i][j] == null)
                            break;
                        var elem = document.createElement("div");
                        elem.className = "result_elem";
                        elem.textContent = convertToLetter(result.Data[i][j]);
                        text[0].appendChild(elem);
                    }
                }
            }
        });
    } else {
        console.log(response.status)
    }
}


/*
// const fingers_update = document.getElementById("fingers_update")
// const fingers_table = document.getElementById("fingers_table")


async function make_tables() {
    let data = {
        Type: document.querySelector("input[name=param]:checked").value,
    };
    console.log(data.Type)

    let response1 = await fetch("/get_words", {method: "POST", body: JSON.stringify(data)});
    // let response2 = await fetch("/get_fingers", {method: "POST"});

    make_table(word_table, response1)
    // make_table(fingers_table, response2)
}

update.onclick = async function () {
    await make_tables()
}

const download_file = document.getElementById("download_file")
download_file.onclick = async function () {
    let url = "/download_file"
    let options = {
        method: 'GET',
        headers: new Headers({
            'Content-Type': 'application/json',
        }),
        mode: 'cors',
        cache: 'default'
    };
    let strFileName;

    //Perform a GET Request to server
    fetch(url, options)
        .then(function (response) {
            let contentType = response.headers.get("Content-Type"); //Get File name from content type
            // strMimeType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet";
            strFileName = contentType.split(";")[1];
            return response.blob();

        }).then(function (myBlob) {
        let downloadLink = window.document.createElement('a');
        downloadLink.href = window.URL.createObjectURL(new Blob([myBlob]));
        downloadLink.download = strFileName;
        document.body.appendChild(downloadLink);
        downloadLink.click();
        document.body.removeChild(downloadLink);
    }).catch((e) => {
        console.log("e", e);
    });
}
 */
