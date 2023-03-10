
let wordsData = JSON.parse(localStorage.getItem("w"))
let fingersData = JSON.parse(localStorage.getItem("f"))
let visualData = JSON.parse(localStorage.getItem("v"))

let resultArea = document.getElementById("results_area");
resultArea.textContent = "";

let headerWords = document.createElement("h1");
let headerFingers = document.createElement("h1");
let headerVisual = document.createElement("h1");
headerWords.textContent = "слухоречевые";
headerFingers.textContent = "кинестетические";
headerVisual.textContent = "зрительные";


resultArea.appendChild(headerWords);
resultArea.appendChild(createTable(wordsData));
resultArea.appendChild(headerFingers);
resultArea.appendChild(createTable(fingersData));
resultArea.appendChild(headerVisual);
resultArea.appendChild(createTable(visualData));








function createTable(data) {
    let table = document.createElement("div");
    table.className = "table";
    for (let i = 0; i < data.length; i++) {
        let level = document.createElement("div");
        level.className = "research_level";
        for (let j = 0; j < data[i].length; j++) {
            let series = document.createElement("div");
            series.className = "research_series";

            let seriesData = document.createElement("div");
            seriesData.className = "series_row_data";
            seriesData.textContent = data[i][j].str;

            series.appendChild(seriesData);

            for (let k = 0; k < data[i][j].answers.length; k++) {
                let seriesAnswer = document.createElement("div");
                seriesAnswer.className = "series_row_answer";
                if (k >= 2)
                    seriesAnswer.textContent = "Э.";
                seriesAnswer.textContent += data[i][j].answers[k];
                series.appendChild(seriesAnswer);
            }
            level.appendChild(series);
        }
        table.appendChild(level);
    }
    return table;
}

