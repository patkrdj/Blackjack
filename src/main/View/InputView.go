package View

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InputView struct {
	Reader *bufio.Reader
}

func NewInputView() *InputView {
	return &InputView{Reader: bufio.NewReader(os.Stdin)}
}

func (view *InputView) ReadString() (int, error) {
	fmt.Print("배팅 금액: ")
	input, err := view.Reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("입력 읽기 실패")
	}
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("정수 변환 실패")
	}
	if num < 1 {
		return 0, fmt.Errorf("베팅 금액은 1보다 커야 함")
	}
	return num, nil
}
