package repository

type Reservation struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

type ReservationRepo struct {
	data map[string]Reservation
}

func NewReservationRepo() *ReservationRepo {
	return &ReservationRepo{data: make(map[string]Reservation)}
}

func (r *ReservationRepo) Save(res Reservation) {
	r.data[res.UserID] = res
}

func (r *ReservationRepo) GetAll() []Reservation {
	var list []Reservation
	for _, v := range r.data {
		list = append(list, v)
	}
	return list
}
