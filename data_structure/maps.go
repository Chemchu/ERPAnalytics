package data_structure

import (
	"github.com/Chemchu/ERPAnalytics/types"
)

type StringProductSummaryMap map[string]types.ProductoSummary

func (m StringProductSummaryMap) Add(stringKey string, productSummary types.ProductoSummary) {
	m[stringKey] = productSummary
}

func (m StringProductSummaryMap) Remove(stringKey string, productSummary types.ProductoSummary) {
	delete(m, stringKey)
}

func (m StringProductSummaryMap) Has(stringKey string) (types.ProductoSummary, bool) {
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

func (m StringProductSummaryMap) Values() []types.ProductoSummary {
	values := []types.ProductoSummary{}
	for _, value := range m {
		values = append(values, value)
	}

	return values
}

// Comprobar que se esta actualizando correctamente el valor del Summary ya que no me queda claro si se esta
// modificando por referencia o valor
func (m *StringProductSummaryMap) UpdateFrecuenciaVentaDiaria(fechas *StringSet) {
	for _, productoSummary := range *m {
		productoSummary.FrecuenciaVentaDiaria = float64(productoSummary.CantidadVendida) / float64(fechas.Length())
	}
}
