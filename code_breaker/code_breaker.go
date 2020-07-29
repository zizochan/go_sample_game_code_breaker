package main

import (
        "fmt"
        "strconv"
        "math/rand"
        "time"
        "sort"
        "strings"
        "regexp"
)

const QUESTION_COUNT int = 4
const MAX_NUMBER int = 9
const INITIAL_LIFE int = 5

const RESULT_HIT = 1
const RESULT_BLOW = 2
const RESULT_NONE = 0

var Question []int
var Life = INITIAL_LIFE

func main() {
        resetQuestion()

        p("=============================")
        p("HIT & BLOW")
        p("")
        p("答えは" + strconv.Itoa(QUESTION_COUNT) + "桁で、範囲は 1〜" + strconv.Itoa(MAX_NUMBER) + " です")
        p("終わるときは exit と入力して下さい")
        p("")
        p("ヒント")
        p("・答えに同じ数字が２つ出ることはありません")
        p("・答えは数字の順番に並んでます（1345とか5789とか）")
        p("=============================")

        for {
                isExit := mainLoop()

                if isExit {
                        p("")
                        p("おしまい")
                        break
                }
        }
}

func mainLoop() bool {
        var input string

        p("")
        p("ライフは残り" + strconv.Itoa(Life) +  "です")
        fmt.Print("input? ")
        fmt.Scan(&input)

        if input == "exit" {
                return true
        }

        if !AnswerValidation(input) {
                p(strconv.Itoa(QUESTION_COUNT) + "桁の数値を入力して下さい")
                return false
        }

        answers := SplitStringToInt(input)
        result := CheckAnswers(answers)

        var isExit bool = false;
        if (result[RESULT_HIT] == QUESTION_COUNT) {
                GameClear()
                isExit = true
        } else {
                isExit = AnswerMiss(result)
        }

        if (isExit) {
                return true
        }

        return false
}

func resetQuestion() {
        rand.Seed(time.Now().UnixNano())

        var tmp []int
        var q int
        for i := 0; i < QUESTION_COUNT; i++ {
                for {
                        q = choiceQuestion()
                        isBreak := true

                        for j := 0; j < len(tmp); j++ {
                                if (tmp[j] == q) {
                                        isBreak = false
                                }
                        }

                        if (isBreak) {
                                break
                        }
                }

                tmp = append(tmp, q)
        }

        sort.Slice(tmp, func(i, j int) bool {
                return tmp[i] < tmp[j]
        })

        Question = tmp
}

func choiceQuestion() int {
        return rand.Intn(MAX_NUMBER) + 1
}

func allQuestion() string {
        var result = ""

        for i := 0; i < QUESTION_COUNT; i++ {
                result += strconv.Itoa(Question[i])
        }

        return result
}

func p(message string) {
        fmt.Println(message)
}

func AnswerValidation(answer string) bool {
        if len(answer) != QUESTION_COUNT {
                return false
        }

        flag, _ := regexp.MatchString("^[0-9]+$", answer)
        return flag
}

func SplitStringToInt(input string) []int {
        var result []int
        slice := strings.Split(input, "")

        for i := 0; i < len(input); i++ {
                s := slice[i]
                v, _ := strconv.Atoi(s)
                result = append(result, v)
        }

        return result
}

func CheckAnswers(answers []int) map[int]int {
        var result = make(map[int]int)

        for i := 0; i < len(Question); i++ {
                resultCode := CheckAnswer(answers, i)
                result[resultCode] += 1
        }

        return result
}

func CheckAnswer(answers []int, questionIndex int) int {
        var isBlow bool = false

        q := Question[questionIndex]

        for i := 0; i < len(answers); i++ {
                a := answers[i]

                if (a != q) {
                        continue
                }

                if (i == questionIndex) {
                        return RESULT_HIT
                } else {
                        isBlow = true
                }
        }

        if (isBlow) {
                return RESULT_BLOW
        }

        return RESULT_NONE
}

func AnswerMiss(result map[int]int) bool {
        p("HIT:" + strconv.Itoa(result[RESULT_HIT]) + " BLOW:" + strconv.Itoa(result[RESULT_BLOW]) + " でした")

        Life -= 1
        if (Life <= 0) {
                GameOver()
                return true
        }

        return false
}

func GameClear() {
        p("")
        p("正解です！おめでとう！")
}

func GameOver() {
        p("")
        p("正解は" + allQuestion() + "でした")
        p("ゲームオーバー")
}
