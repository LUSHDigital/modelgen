package migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/LUSHDigital/modelgen/model"
)

var (
	autoincrementRegExp = regexp.MustCompile(`(?ms) AUTO_INCREMENT=[0-9]*\b`)
)

type MigrationWriter struct {
	outputPath string
}

func NewMigrationWriter(outputPath string) *MigrationWriter {
	return &MigrationWriter{
		outputPath: outputPath,
	}
}

func (w *MigrationWriter) WriteMigrations(migrations []model.Migration) error {
	var (
		upFile   *os.File
		downFile *os.File
		err      error
	)

	w.archive()

	os.Mkdir(w.outputPath, 0777)

	now := time.Now().Unix()

	for _, migration := range migrations {
		// Create the up migration
		upFilePath := filepath.Join(w.outputPath, fmt.Sprintf("%d_create_%s.up.sql", now, migration.TableName))
		if upFile, err = os.Create(upFilePath); err != nil {
			return err
		}

		auto := autoincrementRegExp.FindString(migration.Up)
		migration.Up = strings.Replace(migration.Up, auto, "", 1)
		if _, err = upFile.WriteString(migration.Up + ";"); err != nil {
			return err
		}

		// Create the down migration
		downFilePath := filepath.Join(w.outputPath, fmt.Sprintf("%d_create_%s.down.sql", now, migration.TableName))
		if downFile, err = os.Create(downFilePath); err != nil {
			return err
		}

		if _, err = downFile.WriteString(migration.Down); err != nil {
			return err
		}

		// Ensure the time increments properly
		now++
	}

	return nil
}

// if the folder you are trying to output the migrations in already exists
// archive will move the previous migrations into a timestamped archived folder
// so you can do easy rollbacks
func (w *MigrationWriter) archive() error {
	f, err := os.Stat(w.outputPath)
	if err != nil || !f.IsDir() {
		return nil
	}
	archive := fmt.Sprintf("%s_%s", w.outputPath, time.Now().Format("2006_01_02_15_04_05"))
	if err := os.Rename(w.outputPath, archive); err != nil {
		return fmt.Errorf("cannot archive %s folder", w.outputPath)
	}
}
