

let converter = {}

converter.words = new Map()
converter.words.set(1, "дом")
converter.words.set(2, "лес")
converter.words.set(3, "кот")
converter.words.set(4, "стол")
converter.words.set(5, "звон")
converter.words.set(6, "брат")
converter.words.set(7, "мост")
converter.words.set(8, "час")
converter.words.set(9, "вол")
converter.words.set(10, "рев")

converter.letters = new Map()
converter.letters.set(0, "")
converter.letters.set(1, "д")
converter.letters.set(2, "л")
converter.letters.set(3, "к")
converter.letters.set(4, "с")
converter.letters.set(5, "з")
converter.letters.set(6, "б")
converter.letters.set(7, "м")
converter.letters.set(8, "ч")
converter.letters.set(9, "в")
converter.letters.set(10, "р")

function convertToWord(num) {
    return converter.words.get(num)
}

function convertToLetter(num) {
    return converter.letters.get(num)
}

