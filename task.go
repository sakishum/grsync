package grsync

import (
	"bufio"
	"grsync/rsync"
	"io"
	"strconv"
	"strings"
)

// Task is high-level API under rsync
type Task struct {
	rsync *rsync.Rsync

	state *State
	log   Log
}

// State contains information about rsync process
type State struct {
	Remain   int     `json:"remain"`
	Total    int     `json:"total"`
	Speed    string  `json:"speed"`
	Progress float64 `json:"progress"`
}

// Log contains raw stderr and stdout outputs
type Log struct {
	Stderr string `json:"stderr"`
	Stdout string `json:"stdout"`
}

// State returns inforation about rsync processing task
func (t Task) State() State {
	return *t.state
}

// Log return structure which contains raw stderr and stdout outputs
func (t Task) Log() Log {
	return t.log
}

// Run starts rsync process with options
func (t *Task) Run() error {
	stderr, err := t.rsync.StderrPipe()
	if err != nil {
		return err
	}
	defer stderr.Close()

	stdout, err := t.rsync.StdoutPipe()
	if err != nil {
		return err
	}
	defer stdout.Close()

	go processStdout(t, stdout)
	go processStderr(t, stderr)

	return t.rsync.Run()
}

// New returns new rsync task
func New(source, destination string, rsyncOptions rsync.Options) *Task {
	// Force set required options
	rsyncOptions.Stats = true
	rsyncOptions.Progress = true
	rsyncOptions.HumanReadable = true

	return &Task{
		rsync: rsync.New(source, destination, rsyncOptions),
		state: &State{},
	}
}

var progressMatcher = newMatcher(`\(.+to-chk=(\d+.\d+)`)
var speedMatcher = newMatcher(`(\d+\.\d+.{2}/s)`)

func processStdout(task *Task, stdout io.Reader) {
	const maxPercents = 100
	// Extract data from string:
	//         999,999 99%  999.99kB/s    0:00:59 (xfr#9, to-chk=999/9999)

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		logStr := scanner.Text()
		if progressMatcher.Match(logStr) {
			task.state.Remain, task.state.Total = getTaskProgress(progressMatcher.Extract(logStr))
			task.state.Progress = float64((task.state.Total-task.state.Remain)/task.state.Total) * maxPercents
		}

		if speedMatcher.Match(logStr) {
			task.state.Speed = speedMatcher.Extract(logStr)
		}

		task.log.Stdout += logStr
	}
}

func processStderr(task *Task, stderr io.Reader) {
	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		task.log.Stderr += scanner.Text()
	}
}

func getTaskProgress(remTotalString string) (int, int) {
	const remTotalSeparator = "/"
	const numbersCount = 2
	const (
		indexRem = iota
		indexTotal
	)

	info := strings.Split(remTotalString, remTotalSeparator)
	if len(info) < numbersCount {
		return 0, 0
	}

	remain, _ := strconv.Atoi(info[indexRem])
	total, _ := strconv.Atoi(info[indexTotal])

	return remain, total
}
