package report

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"enigma.com/projectmanagementhub/config"
	"enigma.com/projectmanagementhub/model"
)

// ReportToTXT adalah interface untuk menulis laporan ke file teks.
type ReportToTXT interface {
	WriteReport(report model.ShowReport, status string) error
}

type reportToTXT struct {
	cfg config.PathConfig
}

func NewReportToTXT(cfg config.PathConfig) ReportToTXT {
	return &reportToTXT{cfg: cfg}
}

// WriteReport menulis laporan ke file teks.
func (r *reportToTXT) WriteReport(report model.ShowReport, status string) error {
	// Membuat nama file dengan format "YYYY-MM-DD.txt" di dalam folder
	folderPath := r.cfg.StaticPath //"D:/final_project_enigma2/project-management-hub/report/report-result"
	fileName := filepath.Join(folderPath, time.Now().Format("2006-01-02")+".txt")

	// Membuat folder jika belum ada
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return fmt.Errorf("gagal membuat folder: %v", err)
	}

	// Mengecek apakah file untuk hari tersebut sudah ada atau belum
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		// Jika file belum ada, buat file baru
		file, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("gagal membuat file: %v", err)
		}
		defer file.Close()

		writer := bufio.NewWriter(file)

		// Menggunakan encoding/json untuk mengubah struct Report menjadi JSON
		jsonOldData, err := json.Marshal(report.Content)
		if err != nil {
			return fmt.Errorf("gagal mengonversi ke JSON: %v", err)
		}

		// Check the status and set the appropriate message
		var statusMessage string
		if status == "create" {
			statusMessage = "Create report"
		} else if status == "update" {
			statusMessage = "Update report"
		} else if status == "delete" {
			statusMessage = "Delete report"
		}

		writer.WriteString(fmt.Sprintf("%s\nDate: %s\n%s\n", statusMessage, time.Now().Format("2006-01-02"), jsonOldData))

		writer.WriteString("\n")
		writer.Flush()

		fmt.Printf("File '%s' telah dibuat.\n", fileName)
	} else if err == nil {
		// Jika file sudah ada, baca kontennya, tambahkan data baru, dan tulis kembali
		file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
		if err != nil {
			return fmt.Errorf("gagal membuka file: %v", err)
		}
		defer file.Close()

		// Baca konten lama
		scanner := bufio.NewScanner(file)
		var oldContent string
		for scanner.Scan() {
			oldContent += scanner.Text() + "\n"
		}

		// Tambahkan data baru
		jsonOldData, err := json.Marshal(report.Content)
		if err != nil {
			return fmt.Errorf("gagal mengonversi ke JSON: %v", err)
		}

		// Check the status and set the appropriate message
		var statusMessage string
		if status == "create" {
			statusMessage = "Create report"
		} else if status == "update" {
			statusMessage = "Update report"
		} else if status == "delete" {
			statusMessage = "Delete report"
		}

		newContent := fmt.Sprintf("%s\nDate: %s\n%s\n\n%s", statusMessage, time.Now().Format("2006-01-02"), jsonOldData, oldContent)

		// Pindahkan kursor ke awal file dan tulis konten baru
		file.Seek(0, 0)
		file.Truncate(0)
		file.WriteString(newContent)

		fmt.Printf("File '%s' telah diperbarui.\n", fileName)
	} else {
		return fmt.Errorf("gagal mendapatkan info file: %v", err)
	}

	return nil
}
