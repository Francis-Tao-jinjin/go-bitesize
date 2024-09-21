package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

var totalDuration time.Duration = 5

func getName(r io.Reader, w io.Writer) (string, error) {
	scanner := bufio.NewScanner(r)
	msg := "Your name please? Press the Enter key when done"
	fmt.Fprintln(w, msg)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("No name provided")
	}
	return name, nil
}

func getNameContext(ctx context.Context) (string, error) {
	var err error
	name := "Default Name"
	c := make(chan error, 1)

	go func() {
		name, err = getName(os.Stdin, os.Stdout)
		c <- err
	}()
	// another way to create a timeout channel if not using context
	timeoutChanner := createTimeoutChanner(1)
	select {
	case <-ctx.Done():
		return name, ctx.Err()
	case err := <-c:
		return name, err
	case <-timeoutChanner:
		return name, errors.New("Custome Timeout")
	}
}

/*
*
优点

	简洁：代码非常简洁，易于理解。
	直接：直接返回一个定时器通道，易于使用。

缺点

	缺乏灵活性：无法传递取消信号或元数据。
	无法组合：不能与其他上下文组合使用，无法传递复杂的取消信号。
	资源管理：需要手动管理定时器，可能会导致资源泄漏。
*/
func createTimeoutChanner(duration time.Duration) <-chan time.Time {
	return time.After(duration * time.Second)
}

func main() {
	allowedDuration := totalDuration * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), allowedDuration)
	defer cancel()

	name, err := getNameContext(ctx)
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		fmt.Fprintf(os.Stdout, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "Name: ", name)
}
