package terminal

// Complaint - жалоба на банкомат
type Complaint struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Text     string `json:"text"`
	Date     string `json:"date"`
	Status   string `json:"status"`
}

// Cassette - сущность одной кассеты
type Cassette struct {
	Type     string  `json:"type"`     // "Cash-In" или "Cash-Out"
	Currency string  `json:"currency"` // KZT
	Amount   float64 `json:"amount"`   // Текущая сумма
	Capacity float64 `json:"capacity"` // Макс. вместимость
	Status   string  `json:"status"`   // "OK", "Low", "Full"
}

// ATM - основная сущность терминала
type ATM struct {
	// --- Базовые поля ---
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	IsForte  bool    `json:"isForte"`
	District string  `json:"district"`
	Bank     string  `json:"bank,omitempty"`

	// --- Поля для Конкурентов (Оценочные данные) ---
	// 👇 Именно этих полей не хватало для исправления ошибки
	EstWithdrawalKZT float64 `json:"estWithdrawalKZT,omitempty"` // Оценка: Снятие
	EstDepositKZT    float64 `json:"estDepositKZT,omitempty"`    // Оценка: Внесение

	// --- Поля для Forte (Детальные данные) ---
	AvgCashBalanceKZT float64 `json:"avgCashBalanceKZT,omitempty"`
	TotalCashKZT      float64 `json:"totalCashKZT,omitempty"`

	WithdrawalFreqPerDay int     `json:"withdrawalFreqPerDay,omitempty"`
	DowntimePct          float64 `json:"downtimePct,omitempty"`
	EfficiencyStatus     string  `json:"efficiencyStatus,omitempty"`

	Cassettes  []Cassette  `json:"cassettes,omitempty"`
	Complaints []Complaint `json:"complaints,omitempty"`
}
