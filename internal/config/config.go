package config

import (
	"errors"
	"flag"
)

type Config struct {
	URL         string
	Requests    int
	Concurrency int
	SaveToFile  bool
	Detail      bool
}

func ParseFlags() (*Config, error) {
	cfg := &Config{}
	flag.StringVar(&cfg.URL, "url", "", "URL do serviço a ser testado (obrigatório)")
	flag.IntVar(&cfg.Requests, "requests", 1, "Número total de requests (>= 1)")
	flag.IntVar(&cfg.Concurrency, "concurrency", 1, "Número de chamadas simultâneas (>= 1)")
	flag.BoolVar(&cfg.SaveToFile, "save", false, "Salvar relatório em arquivo JSON")
	flag.BoolVar(&cfg.Detail, "detail", false, "Incluir detalhes de cada request no relatório")
	flag.Parse()

	if cfg.URL == "" {
		return nil, errors.New("a URL é obrigatória")
	}
	if cfg.Requests < 1 {
		return nil, errors.New("o número de requests deve ser >= 1")
	}
	if cfg.Concurrency < 1 {
		return nil, errors.New("o número de chamadas simultâneas deve ser >= 1")
	}

	return cfg, nil
}
