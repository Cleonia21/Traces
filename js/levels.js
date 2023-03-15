class History {
    #beforeElems
    #currentElems
    #afterElems

    #beforeLevel
    #currentLevel
    #afterLevel

    constructor(beforeLevel, currentLevel, afterLevel) {
        this.#beforeElems = document.getElementById("before").getElementsByClassName("preview_row");
        this.#currentElems = document.getElementById("current").getElementsByClassName("preview_row");
        this.#afterElems = document.getElementById("after").getElementsByClassName("preview_row");

        this.#beforeLevel = beforeLevel;
        this.#currentLevel = currentLevel;
        this.#afterLevel = afterLevel;
    }

    writeHistory() {
        this.#writeLevelInHistory(this.#beforeElems, this.#beforeLevel);
        this.#writeLevelInHistory(this.#currentElems, this.#currentLevel);
        this.#writeLevelInHistory(this.#afterElems, this.#afterLevel);
    }

    updateAnswer() {
        this.#writeLevelInHistory(this.#currentElems, this.#currentLevel);
    }

    moveHistory(newLevel) {
        this.#beforeLevel = this.#currentLevel;
        this.#currentLevel = this.#afterLevel;
        this.#afterLevel = newLevel;

        this.writeHistory();
    }

    #writeLevelInHistory(elems, level) {
        if (level === null || level === undefined) {
            for (let i = 0; i < elems.length; i++)
                elems[i].textContent = "?"
            return
        }

        for (let i = 0; i < level.length; i++) {
            let answersStr = "";
            for (let j = 0; j < level[i].answers.length; j++) {
                let answer = level[i].answers[j];
                if (answer !== "+" && answer !== "?")
                    answer = "-"
                answersStr += answer;
            }
            elems[i].textContent = level[i].str + " " + answersStr;
        }
    }
}

class Research {
    #levels = []
    #levItr
    #seriesItr
    #history

    #type

    constructor(numsArray, type) {
        this.#type = type;
        this.#levItr = 0;
        this.#seriesItr = 0;

        for (let i = 0; i < numsArray.length; i++) {
            this.#levels[i] = []
            for (let j = 0; j < numsArray[i].length; j++) {
                this.#levels[i][j] = {
                    // nums: numsArray[i][j],
                    str: "",
                    answers: ["?", "?", "?", "?"]
                }
                for (let k = 0; k < numsArray[i][j].length; k++) {
                    let letter = "*"
                    if (this.#type === "w")
                        letter = convertToLetter(numsArray[i][j][k]);
                    if (this.#type === "f")
                        letter = numsArray[i][j][k];
                    this.#levels[i][j].str += letter;
                }
            }
        }
        this.#history = new History(null, this.#levels[0], this.#levels[1])
    }

    localSave() {
        this.#levels.length = this.#levItr + 1;
        localStorage.setItem(this.#type, JSON.stringify(this.#levels))
    }

    #createElem(content) {
        let elem = document.createElement("button");
        elem.onmousedown = function () { return false; }
        elem.className = "series_elem";
        elem.textContent = content;
        elem.onclick = function () {
            let inpt = document.querySelector('input:focus');
            if (inpt) {
                inpt.value += elem.textContent;
            }
        }
        return elem;
    }

    put_series() {
        if (this.#levItr === 0 && this.#seriesItr === 0)
            this.#history.writeHistory();

        let series_container = document.getElementById("series_field");

        while (series_container.firstChild)
            series_container.removeChild(series_container.firstChild)

        let series = this.#levels[this.#levItr][this.#seriesItr];

        for (let i = 0; i < series.str.length; i++) {
            series_container.appendChild(this.#createElem(series.str[i]));
        }
    }

    #transition() {
        let correctAnswers = 0;
        let fullySeries = false;

        let level = this.#levels[this.#levItr];
        for (let i = 0; i < level.length; i++) {
            let tmpCorrectAnswers = 0;
            let answers = level[i].answers;

            for (let k = 0; k < answers.length; k++) {
                if (answers[k] === "+")
                    tmpCorrectAnswers++;
            }

            correctAnswers += tmpCorrectAnswers;
            if (tmpCorrectAnswers === answers.length)
                fullySeries = true;
        }
        return fullySeries === true || correctAnswers >= 5;
    }

    get_input () {
        let inputs = document.getElementsByClassName("series_input")
        let series = this.#levels[this.#levItr][this.#seriesItr];
        for (let i = 0; i < inputs.length; i++) {
            series.answers[i] = inputs[i].value
            if (series.str === inputs[i].value)
                series.answers[i] = "+"
            if (inputs[i].value === "" || inputs[i].value[0] === "*")
                series.answers[i] = "-"
            inputs[i].value = ""
        }

        if (this.#seriesItr < 4) {
            this.#history.updateAnswer(this.#seriesItr, series);
            this.#seriesItr++;
        } else {
            if (this.#transition() === false)
                return false;

            this.#seriesItr = 0;
            this.#levItr++;
            this.#history.moveHistory(this.#levels[this.#levItr + 1]);
        }
        if (this.#levItr === this.#levels.length)
            return false;
        this.put_series();
        return true;
    }
}

class ResearchHub {
    #current = []
    #navElems
    #itr

    constructor(research) {
        console.log(research);

        this.#current[0] = new Research(research.Words, "w")
        this.#current[1] = new Research(research.Fingers, "f")
        this.#current[2] = new Research(research.Visual, "v")

        this.#navElems = document.getElementsByClassName("navigation_elem");

        console.log(this.#current);

        this.#itr = 0;
        this.#current[this.#itr].put_series()
        this.#navElems[this.#itr].style.backgroundColor = "#ff7171";
    }

    next() {
        if (this.#current[this.#itr].get_input() === false) {
            this.#current[this.#itr].localSave();
            this.#itr++;
            if (this.#itr === 3) {
                alert("Исследование завершено полностью");
                window.location.href = "/research/results";
                this.#itr--; //Переход к результатам
            } else {
                this.#current[this.#itr].put_series()
                this.#navElems[this.#itr - 1].style.backgroundColor = "#e5e5e5";
                this.#navElems[this.#itr].style.backgroundColor = "#ff7171";
            }
        }
    }
}

let researchHub

const complete_button = document.getElementById("complete");

let get_research = new Promise(async function(resolve, reject) {
    let response = await fetch("get_research", {method: "POST"});

    if (response.ok) {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            resolve(result.Levels)
        });
    } else {
        reject(response.status)
    }
});

get_research.then(
    function (levels) {
        researchHub = new ResearchHub(levels)
    },
    function (error) {
        console.log(error)
    }
);

complete_button.onclick = function () {
    researchHub.next();
    document.getElementById("input_1").focus()
}
