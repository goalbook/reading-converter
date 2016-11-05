package main

import (
    "log"
    "testing"
)

func TestNewConversionMap(t *testing.T) {
    log.Printf("conversion map:")
}

func TestNewLexileConverter(t *testing.T) {
    lc, err := NewLexileConverter()
    if err != nil {
        t.Errorf("Error converting %v", err)
    }
    rBR, err := lc.Convert("BR")
    if err != nil {
        t.Errorf("Error converting %v", err)
    }
    log.Printf("BR: %v", rBR)

    r540, err := lc.Convert("540")
    if err != nil {
        t.Errorf("Error converting %v", err)
    }
    log.Printf("540: %v", r540)

    r1000, err := lc.Convert("1000")
    if err != nil {
        t.Errorf("Error converting %v", err)
    }
    log.Printf("1000: %v", r1000)
}