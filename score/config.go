package score

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type Action int

const (
	increase Action = iota
	decrease
)

type limit struct {
	Duration duration
	Count    uint `yaml:"count" validate:"required"`
}

type duration struct {
	Count uint   `yaml:"count" validate:"required"`
	Unit  string `yaml:"unit" validate:"required"`
}

// the config contains some rules:
// some events can only score per day
// some events can only score once all the time
// the event can cause increase or decrease score
type Rule struct {
	Name       string `yaml:"event" validate:"required"`
	Limit      limit
	Action     string `yaml:"action" validate:"required"`
	realAction Action `validate:"required"`
	Score      uint   `yaml:"score" validate:"required"`
}

// a config item in yaml format looks like
//
// following:
//
//	MaxADay: 1
//	MaxWholeLife: 1
//	action: increase
//	score: 5
func ParseRule(path string) (*Rule, error) {
	var cfg Rule
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file, err: %w", err)
	}

	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, fmt.Errorf("parse config failed, err: %w", err)
	}

	realAction := increase
	if cfg.Action == "decrease" {
		realAction = decrease
	} else if cfg.Action == "increase" {
		realAction = increase
	} else {
		log.Fatalf("invalid action: %s", cfg.Action)
	}
	cfg.realAction = realAction

	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		return nil, fmt.Errorf("validate config failed, err: %w", err)
	}
	return &cfg, nil
}
