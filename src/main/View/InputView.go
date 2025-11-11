package View

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InputView struct {
	reader *bufio.Reader
}

func NewInputView() *InputView {
	return &InputView{reader: bufio.NewReader(os.Stdin)}
}

func (view *InputView) ReadString() (int, error) {
	fmt.Print("배팅 금액: ")
	input, err := view.reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return num, nil
}
