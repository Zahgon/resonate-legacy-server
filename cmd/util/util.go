package util

import (
	"math/rand" // nosemgrep

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func SafeDeref[T any](val *T) T { _ = "STUB: not implemented"; return *new(T) }

func PrettyHeaders(headers map[string]string, separator string) []string {
	_ = "STUB: not implemented"
	return nil
}

func PrettyData(data string) string { _ = "STUB: not implemented"; return "" }

func MapToBytes() mapstructure.DecodeHookFunc {
	_ = "STUB: not implemented"
	return *new(mapstructure.DecodeHookFunc)
}

// Random value generators
func RangeIntn(r *rand.Rand, min int, max int) int { _ = "STUB: not implemented"; return 0 }

func RangeInt63n(r *rand.Rand, min int64, max int64) int64 { _ = "STUB: not implemented"; return 0 }

func RangeFloat63n(r *rand.Rand, min float64, max float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func RangeMap[K comparable, V any](r *rand.Rand, m map[K]V) K {
	_ = "STUB: not implemented"
	return *new(K)
}

// nosemgrep: range-over-map

func Choose[T any](r *rand.Rand, v ...T) T { _ = "STUB: not implemented"; return *new(T) }

// Bind configuration struct fields to cobra flags and viper config
func Bind(cfg any, cmd *cobra.Command, flg *pflag.FlagSet, vip *viper.Viper, prefixes ...string) {
	_ = "STUB: not implemented"
	return
}

// TODO: support additional slice types

// TODO: support additional map types

// Extract a nested key from viper settings
func Extract(m map[string]any, path string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}
