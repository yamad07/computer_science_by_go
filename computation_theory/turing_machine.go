// チューリングマシンを擬似的に実装したプログラム

package main
import (
    "fmt"
)

type TuringMachine struct{
    state int;
}

func (tm *TuringMachine) run(t *Tape){
    current_color := t.color[t.index]
    state, direction, color := tm.rule(current_color)
    tm.write(state)
    t.write(color, direction)
}

func (tm *TuringMachine) rule(color int) (int, int, int){
    switch{
        case tm.state == 1 && color == 1:
            return 3, 0, 1
        case tm.state == 1 && color == 0:
            return 2, 1, -1
        case tm.state == 3 && color == 1:
            return 2, 1, 1
        case tm.state == 3 && color == 0:
            return 1, 1, 1
        case tm.state == 2 && color == 1:
            return 1, 0, -1
        case tm.state == 2 && color == 0:
            return 3, 1, 1
    }
    return 0, 0, 0
}

func (tm *TuringMachine) write(state int){
    tm.state = state
}

type Tape struct{
    color [50]int;
    index int;
}

func (t *Tape) write(color int, direction int){
    t.color[t.index] = color
    t.index = direction
}

func main() {
    tm := &TuringMachine{1}
    colors := [50] int{}
    colors [25] = 1
    t := &Tape{colors, 25}
    fmt.Println(t.color[24], t.color[25], t.color[26])
    fmt.Println(tm.state)
    tm.run(t)
    fmt.Println(t.color[24], t.color[25], t.color[26])
    fmt.Println(tm.state)
}
