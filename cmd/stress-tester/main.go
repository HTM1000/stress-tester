package main

import (
	"fmt"
	"os"

	"github.com/HTM1000/stress-tester/internal/config"
	"github.com/HTM1000/stress-tester/internal/tester"
)

func main() {
	cfg, err := config.ParseFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	report, err := tester.RunLoadTest(cfg)
	if err != nil {
		fmt.Printf("Erro ao executar o teste: %v\n", err)
		os.Exit(1)
	}

	tester.PrintReport(report)

	if cfg.SaveToFile {
		if err := tester.SaveReportToFile(report); err != nil {
			fmt.Printf("Erro ao salvar relatório: %v\n", err)
		} else {
			fmt.Println("Relatório salvo com sucesso em 'stress_test_report.json'")
		}
	}
}
