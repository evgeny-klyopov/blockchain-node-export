package export

import (
	"fmt"
	"github.com/evgeny-klyopov/blockchain-node-export/internal/node"
	each "github.com/evgeny-klyopov/each"
	"os"
	"time"
)

type fileWriter struct {
	header              *string
	outputDir           string
	name                string
	outputFileExtension string
	file                *os.File
	format              func(transaction node.Transaction) string
	writeFormat         func(lines []string) string
	bufferLength        int
	debug               bool
}

func (f *fileWriter) setHeader() error {
	var err error

	if f.header != nil {
		_, err = f.file.WriteString(*f.header)
	}

	return err
}

func (f *fileWriter) Run(transactions []node.Transaction) error {
	var err error
	fileName := fmt.Sprintf("%s_%d.%s", f.name, time.Now().Unix(), f.outputFileExtension)

	f.file, err = os.Create(f.outputDir + "/" + fileName)

	if err != nil {
		return err
	}

	defer func() {
		_ = f.file.Close()
	}()

	err = f.setHeader()
	if err != nil {
		return err
	}

	e := each.NewEach(f.bufferLength, func(lines []string, hasError bool) error {
		if hasError == true {
			return nil
		}
		_, err = f.file.WriteString(f.writeFormat(lines) + "\n")

		return err
	})

	for _, t := range transactions {
		hasError := e.Add(f.format(t))
		if true == hasError {
			break
		}
	}
	e.Close()

	return e.GetError()
}
