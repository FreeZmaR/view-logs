package reader

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const consoleReaderMockText1024 = `bMyUHeabVMvaaSdWxiMrDZcQtzZfKISXmnqxmgkoPdPhzRvRBuupGRQKviHHjwxFiLnFmOqFqIfWhMHiBzVZsMr
KfhpsCSxDRWtYTuDSQLgtSeJsDAACxuBXHaZLckTVbluKKUaerwqdftAvKfdnFkYPkUooFnJtqpRqkAIWEOaaOTNWmRbzdWekYtXwGoDhLgGTKqXeiCScbcZ
aGbtolHiaQDqOOaHuMbAeIVGvIrjZYKUpBPeXvVMXqysbmMncMIxKOTNYAXzinKlWTYaYpbqyhyfAxENbNmbbqrtVexGDevhpCUQLSqrOsvKrFBpTDyjehre
YpEoyuJiAsxXgCylxTnAdaAabyHKcyifBhPFXhUrBEiFAlfytgRzLkwYswWJjpqxrneIsNyGyEaRiNYLKrFaDNFLckWCcDzxWLhNEDZtJxfgQdrBrwoGzrEU
jNQjwjTkMQhTZQhRuZcoqQQBVoweKEWcqvurfJnBXspQuZKrqnPkuPokChicEOBfCLrFEgQFaxoTHOMrMDsczwXhzonOxGanMvZCoxiifbzpHsbjIQZmojEI
iuoLTHGAptzLGUjFtMblSMAUnoxmgWEIiHHdzeCKkPRsENebJkQPBWdxyLlYQbBBbDgUSiZRErUGSHXtQYxBYfhPcptOqCfavcVntPMFPLtGFXdEbtZBKRuq
PeIWVOwLCFbzpYCCNFRdZljLrdDjaDbWalPsGsKQRpPvyLJFThVMAltxSrgcMefspwbjWSHFouGDwHIFIieIXUZlqSKLQaeasWNQJAlaguvLAWrvETMNNQWb
ScJprzFCvNrnOyiLZIToFIaYVJrFVjraSzgnmEdWbexZheGhVRedyhZfGVnKDCJDKIIrrjqzfjPPHidfihRrPgmHMmNORhuYFzjmYnpRtbhhbKAMMNkrJcuD
sAaZWHGdRBquIFAmYnlzKeidZQUdekSYqwrKWFwYTTkBijpdJWlrRHmCRLCouxCtUyVZHyrDdxziLysKjvZeAzyeO`

type consoleReaderApiTest struct {
	content []byte
	path    string
	run     RunCommandFN
	int     InterruptCommandFN
	buf     *bytes.Buffer
	cmd     *exec.Cmd
}

func newConsoleApiTest(content []byte, path string) (*consoleReaderApiTest, error) {
	api := &consoleReaderApiTest{
		content: content,
		path:    path,
		buf:     &bytes.Buffer{},
	}

	err := api.reInit()
	if err != nil {
		return nil, err
	}

	var cmd *exec.Cmd

	api.run = func(command string) error {
		splitCommand := strings.Split(command, " ")
		cmd = exec.Command(splitCommand[0], splitCommand[1:]...)
		cmd.Stdout = api.buf

		return cmd.Run()
	}

	api.int = func() error {
		if cmd == nil {
			return nil
		}

		return cmd.Process.Kill()
	}

	return api, nil
}

func (api *consoleReaderApiTest) resetBuf() {
	api.buf.Reset()
}

func (api *consoleReaderApiTest) reInit() error {
	api.buf.Reset()

	return os.WriteFile(api.path, api.content, 0777)
}

func (api *consoleReaderApiTest) cleanup() {
	api.resetBuf()

	err := os.Remove(api.path)
	if err != nil {
		fmt.Println(err)
	}
}

func TestConsoleReader_SuccessCommand(t *testing.T) {
	type testCase struct {
		name      string
		command   ConsoleCommand
		path      string
		content   []byte
		assert    []byte
		readCount int
	}

	testCases := []testCase{
		{
			name:      "Cat",
			path:      "cat.log",
			command:   CatCommand,
			content:   []byte("test"),
			assert:    []byte("test"),
			readCount: 1,
		},
		{
			name:      "Tail",
			path:      "tail.log",
			command:   TailCommand,
			content:   []byte("test\nlong\ncontent"),
			assert:    []byte("\n==> tail.log <==\ntest\nlong\ncontent"),
			readCount: 1,
		},
		{
			name:      "Head 1",
			path:      "head_1.log",
			command:   HeadCommand,
			content:   []byte("test\nlong\ncontent"),
			assert:    []byte("test\n"),
			readCount: 1,
		},
		{
			name:      "Head 2",
			path:      "head_2.log",
			command:   HeadCommand,
			content:   []byte("test\nlong\ncontent"),
			assert:    []byte("test\nlong\n"),
			readCount: 2,
		},
		{
			name:      "Head All",
			path:      "head_3.log",
			command:   HeadCommand,
			content:   []byte("test\nlong\ncontent"),
			assert:    []byte("test\nlong\ncontent"),
			readCount: 3,
		},
		{
			name:      "DD 1",
			path:      "dd_1.log",
			command:   DDCommand,
			content:   []byte(consoleReaderMockText1024),
			assert:    []byte(consoleReaderMockText1024),
			readCount: 1,
		},
		{
			name:      "DD 2",
			path:      "dd_2.log",
			command:   DDCommand,
			content:   []byte(consoleReaderMockText1024 + consoleReaderMockText1024),
			assert:    []byte(consoleReaderMockText1024),
			readCount: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			api, err := newConsoleApiTest(tc.content, tc.path)
			if err != nil {
				t.Errorf("unexpected error: %v", err)

				return
			}

			r := NewConsoleReader(tc.path, tc.command)
			InitReader(r, api.buf, api.run, api.int)

			defer func() {
				api.cleanup()

				if err = r.Close(); err != nil {
					fmt.Println("error close console reader: ", err)
				}
			}()

			var (
				result []byte
			)
			for i := 0; i < tc.readCount; i++ {
				var errRead error

				result, errRead = r.Read()
				if errRead != nil {
					t.Errorf("expected error: nil , got: %v", errRead)

					return
				}

				api.resetBuf()
			}

			if !bytes.Equal(result, tc.assert) {
				t.Errorf("expected result: %v, got: %v", tc.assert, result)
			}
		})
	}
}
