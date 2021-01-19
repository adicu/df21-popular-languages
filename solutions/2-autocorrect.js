// Source: https://gist.github.com/andrei-m/982927#gistcomment-1796676
String.prototype.levenshtein = function(string) {
    let a = this, b = string + "", m = [], i, j, min = Math.min;

    if (!(a && b)) return (b || a).length;

    for (i = 0; i <= b.length; m[i] = [i++]);
    for (j = 0; j <= a.length; m[0][j] = j++);

    for (i = 1; i <= b.length; i++) {
        for (j = 1; j <= a.length; j++) {
            m[i][j] = b.charAt(i - 1) == a.charAt(j - 1)
                ? m[i - 1][j - 1]
                : m[i][j] = min(
                    m[i - 1][j - 1] + 1, 
                    min(m[i][j - 1] + 1, m[i - 1 ][j]))
        }
    }

    return m[b.length][a.length];
}

console.log("bat".levenshtein("bar"))

const dictionary = [
    "banana",
    "orange",
    "apple",
    "pear",
    "mango",
    "watermelon",
    "pineapple"
]

function autocorrect(word) {
    let shortestDistance = word.length
    let correctWord = word
    
    let currentDistance = 0
    for (const dictionaryWord of dictionary) {
        currentDistance = word.levenshtein(dictionaryWord)
        if (currentDistance < shortestDistance) {
            shortestDistance = currentDistance
            correctWord = dictionaryWord
        }
    }
    
    return correctWord
}

console.log(autocorrect("babana"))
