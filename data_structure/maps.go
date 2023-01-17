package data_structure

import "github.com/Chemchu/ERPAnalytics/types"

type StringProductSummaryMap map[string]types.ProductSummary

func (m StringProductSummaryMap) Add(stringKey string, productSummary types.ProductSummary) {
	m[stringKey] = productSummary
}

func (m StringProductSummaryMap) Remove(stringKey string, productSummary types.ProductSummary) {
	delete(m, stringKey)
}

func (m StringProductSummaryMap) Has(stringKey string) (types.ProductSummary, bool) {
	p, ok := m[stringKey]
	return p, ok
}

func (m StringProductSummaryMap) Keys() []string {
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func (m StringProductSummaryMap) Values() []types.ProductSummary {
	values := []types.ProductSummary{}
	for _, value := range m {
		values = append(values, value)
	}

	return values
}
