package utils

import "strings"

type VariantSet struct {
	variants map[string]bool
}

func NewVariantSet(variants []string) *VariantSet {
	vs := &VariantSet{
		variants: make(map[string]bool),
	}
	for _, variant := range variants {
		vs.variants[strings.ToLower(variant)] = true
	}
	return vs
}

func (vs *VariantSet) Has(variant string) bool {
	return vs.variants[strings.ToLower(variant)]
}

func (vs *VariantSet) GetAll() []string {
	variants := make([]string, 0, len(vs.variants))
	for variant := range vs.variants {
		variants = append(variants, variant)
	}
	return variants
}

func IsVariantEmpty(variant string) bool {
	return strings.TrimSpace(variant) == ""
}

func CatppuccinVariants() *VariantSet {
	return NewVariantSet([]string{"latte", "frappe", "macchiato", "mocha"})
}

func GetDefaultCatppuccinVariant() string {
	return "mocha"
}

func TokyoNightVariants() *VariantSet {
	return NewVariantSet([]string{"storm", "night", "day", "moon"})
}

func GetDefaultTokyoNightVariant() string {
	return "storm"
}
