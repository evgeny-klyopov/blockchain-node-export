package export

import (
	_ "github.com/lib/pq"
)

type sqlFileExport struct {
	fileWriter
}

func newSqlFileExport(
	name string,
	outputDir string,
	outputFileExtension string,
	databaseType string,
	tableName string,
	debug bool,
) Exporter {
	f := newFormat(databaseType, tableName)
	return &sqlFileExport{
		fileWriter{
			name:                name,
			outputDir:           outputDir,
			outputFileExtension: outputFileExtension,
			header:              nil,
			bufferLength:        300,
			writeFormat:         f.insert,
			format:              f.add,
			debug:               debug,
		},
	}
}
