package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Input parse IO data by line
type Input struct {
	br *bufio.Reader
}

func NewInput(r io.Reader) *Input {
	return &Input{
		br: bufio.NewReader(r),
	}
}

// read a line without \n
func (in *Input) readline() ([]byte, error) {
	buf, err := in.br.ReadBytes('\n')
	if err == io.EOF {
		return buf, nil
	}
	if err != nil {
		return nil, err
	}
	return buf[:len(buf)-1], nil
}

// Num parse one line and return a num
func (in *Input) Num() (int, error) {
	buf, err := in.readline()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(buf))
}

// Num parse one line and return a num slice
func (in *Input) Nums() ([]int, error) {
	buf, err := in.readline()
	if err != nil {
		return nil, err
	}
	nums := make([]int, 0, 4)
	i, j := 0, 0
	for ; j < len(buf); j++ {
		if buf[j] == ' ' {
			continue
		}
		i = j
		for ; j < len(buf) && buf[j] != ' '; j++ {
		}
		num, err := strconv.Atoi(string(buf[i:j]))
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func (in *Input) String() (string, error) {
	s, err := in.readline()
	return string(s), err
}

func (in *Input) Bytes() ([]byte, error) {
	return in.readline()
}

func main() {
	in := NewInput(os.Stdin)
	fmt.Println("输入值数量:")
	line, _ := in.Num()
	for i := 0; i < line; i++ {
		fmt.Println("输入值num:")
		num, _ := in.Num()
		fmt.Println(num)

		fmt.Println("输入值nums:")
		nums, _ := in.Nums()
		fmt.Println(nums)
		fmt.Println(nums[1])
		//doSomething(num, nums)
	}
}
