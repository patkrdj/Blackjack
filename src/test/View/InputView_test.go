package View

import (
	"Blackjack/src/main/View"
	"bufio"
	"strings"
	"testing"
)

func Test_ReadString_NormalInput(t *testing.T) {
	fakeInput := "400\n"
	inputView := &View.InputView{Reader: bufio.NewReader(strings.NewReader(fakeInput))}

	value, err := inputView.ReadString()

	if err != nil {
		t.Errorf("예상치 않은 에러 발생")
	}
	if value != 400 {
		t.Errorf("결과값이 다름")
	}
}

func Test_Exception_Negative(t *testing.T) {
	fakeInput := "-122\n"
	inputView := &View.InputView{Reader: bufio.NewReader(strings.NewReader(fakeInput))}

	_, err := inputView.ReadString()

	if err == nil {
		t.Errorf("예상치 않은 에러 발생: %v", err)
	} else {
		t.Logf("에러가 정상적으로 동작")
	}
}

func Test_Exception_EmptyString(t *testing.T) {
	fakeInput := "\n"
	inputView := &View.InputView{Reader: bufio.NewReader(strings.NewReader(fakeInput))}

	_, err := inputView.ReadString()

	if err == nil {
		t.Errorf("예상치 않은 에러 발생: %v", err)
	} else {
		t.Logf("에러가 정상적으로 동작")
	}
}

func Test_Exception_NotDigit(t *testing.T) {
	fakeInput := "good\n"
	inputView := &View.InputView{Reader: bufio.NewReader(strings.NewReader(fakeInput))}

	_, err := inputView.ReadString()

	if err == nil {
		t.Errorf("예상치 않은 에러 발생: %v", err)
	} else {
		t.Logf("에러가 정상적으로 동작")
	}
}
