package export

import (
	"errors"
)

const fileExportType = "file"
const sqlFileType = "sql"
const csvFileType = "csv"

func newFileExport(
	name string,
	outputFileExtension string,
	outputDir string,
	databaseType string,
	tableName string,
	debug bool,
) (Exporter, error) {
	if outputFileExtension == sqlFileType {
		return newSqlFileExport(name, outputDir, outputFileExtension, databaseType, tableName, debug), nil
	} else if outputFileExtension == csvFileType {
		return newCsvFileExport(name, outputDir, outputFileExtension, debug), nil
	}

	return nil, errors.New("not support file type export")
}
