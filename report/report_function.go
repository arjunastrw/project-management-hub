package report

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"enigma.com/projectmanagementhub/model"
)

type ReportToTXT interface {
	WriteReport(report model.ShowReport) error
}

type reportToTXT struct {
}

func (r *reportToTXT) WriteReport(report model.ShowReport) error {
	// Membuat nama file dengan format "YYYY-MM-DD_HH-MM-SS.txt" di dalam folder
	folderPath := "D:/Final_project_enigma/project-management-hub/report/report_result"
	fileName := filepath.Join(folderPath, time.Now().Format("2006-01-02_15-04-05")+".txt")

	// Membuat folder jika belum ada
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Menggunakan encoding/json untuk mengubah struct Report menjadi JSON
	jsonData, err := json.Marshal(report.Content)
	if err != nil {
		return err
	}

	writer.WriteString(fmt.Sprintf("Date: %s\n%s\n", report.Date.Format("2006-01-02"), jsonData))

	writer.WriteString("\n")
	writer.Flush()

	fmt.Printf("File '%s' has created.\n", fileName)
	return nil
}
