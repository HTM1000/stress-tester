package tester

import (
	"encoding/json"
	"fmt"
	"os"
)

func PrintReport(r *Report) {
	fmt.Printf("\n=== Relatório de Teste ===\n")
	fmt.Printf("Tempo total: %v\n", r.TotalTime)
	fmt.Printf("Requests totais: %d\n", r.TotalRequests)
	fmt.Printf("Status 200: %d\n", r.Successful200)
	fmt.Println("Distribuição de outros códigos de status:")
	for code, count := range r.StatusCounts {
		if code != 200 {
			fmt.Printf("  %d: %d\n", code, count)
		}
	}
}

func SaveReportToFile(r *Report) error {
	file, err := os.Create("stress_test_report.json")
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(r); err != nil {
		return fmt.Errorf("erro ao salvar relatório: %v", err)
	}

	return nil
}
