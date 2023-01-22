// const name = document.getElementById("name")
const update = document.getElementById("update")
const word_table = document.getElementById("word_table")

// const fingers_update = document.getElementById("fingers_update")
const fingers_table = document.getElementById("fingers_table")

function make_table(elem, response) {
    if (response.ok) {
        response.text().then(function (data) {
            let result = JSON.parse(data);

            if (result.Result === "ok") {
                let html = "";
                for (let k = 0; k < result.Data.length; k++) {
                    for (let i = 0; i < result.Data[k].length; i++) {
                        let tds = "";
                        for (let j = 0; j < result.Data[k][i].length; j++) {
                            tds += "<td>" + result.Data[k][i][j] + "</td>";
                        }
                        html += "<tr>" + tds + "</tr>";
                    }
                    html += "<td>" + "</td>"
                }
                elem.innerHTML = html;
            }
        });
    } else {
        console.log(response.status)
    }
}

async function make_tables() {
    let data = {
        Type: document.querySelector("input[name=param]:checked").value,
    };
    console.log(data.Type)
    let response1 = await fetch("/get_words", {method: "POST", body: JSON.stringify(data)});
    let response2 = await fetch("/get_fingers", {method: "POST"});

    make_table(word_table, response1)
    make_table(fingers_table, response2)
}

update.onclick = async function () {
    await make_tables()
}

update.click()


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