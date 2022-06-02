package stats

import "github.com/ShavqatKavrakov/Lesson10_bank_1/pkg/bank/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) (sum types.Money) {
	for _, payment := range payments {
		sum += payment.Amount
	}
	return sum / types.Money(len((payments)))
}

//TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) (sum types.Money) {
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		sum += payment.Amount
	}
	return sum
}
