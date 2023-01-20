// const name = document.getElementById("name")
const word_update = document.getElementById("word_update")
const word_table = document.getElementById("word_table")

const fingers_update = document.getElementById("fingers_update")
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

word_update.onclick = async function () {
    let response = await fetch("/get_words", {method: "POST"});

    make_table(word_table, response)
}

fingers_update.onclick = async function () {
    let response = await fetch("/get_fingers", {method: "POST"});

    make_table(fingers_table, response)
}

// word_update.addEventListener("click", async function () {
//     let response = await fetch("/get_words", {method: "POST"});
//
//     if (response.ok) {
//         response.text().then(function (data) {
//             let result = JSON.parse(data);
//
//             if (result.Result === "ok") {
//                 let html = "";
//                 for (let k = 0; k < result.Data.length; k++) {
//                     for (let i = 0; i < result.Data[k].length; i++) {
//                         let tds = "";
//                         for (let j = 0; j < result.Data[k][i].length; j++) {
//                             tds += "<td>" + result.Data[k][i][j] + "</td>";
//                         }
//                         html += "<tr>" + tds + "</tr>";
//                     }
//                     html += "<td>" + "</td>"
//                 }
//                 word_table.innerHTML = html;
//             }
//         });
//     } else {
//         console.log(response.status)
//     }
// })