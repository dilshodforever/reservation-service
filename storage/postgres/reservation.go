package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/dilshodforever/reservation-service/genprotos"

	"github.com/google/uuid"
)

type Reservationstorage struct {
	db *sql.DB
}

func NewReservationstorage(db *sql.DB) *Reservationstorage {
	return &Reservationstorage{db: db}
}

func (p *Reservationstorage) CreateReservation(reservation *pb.Reservation) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO reservations (id, user_id, restaurant_id, reservation_time, status)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := p.db.Exec(query, id, reservation.UserId, reservation.RestaurantId, reservation.ReservationTime, reservation.Status)
	return nil, err
}

func (p *Reservationstorage) GetByIdReservation(id *pb.ById) (*pb.Reservation, error) {
	query := `
			SELECT user_id, restaurant_id, reservation_time, status from reservations 
			where id =$1 and delated_at=0
		`
	row := p.db.QueryRow(query, id.Id)

	var Reservation pb.Reservation

	err := row.Scan(&Reservation.UserId, &Reservation.RestaurantId, &Reservation.ReservationTime, &Reservation.Status)
	if err != nil {
		return nil, err
	}

	return &Reservation, nil
}

func (p *Reservationstorage) GetAllReservation(res *pb.Reservation) (*pb.GetAllReservations, error) {
	Reservations := &pb.GetAllReservations{}
	var query string
	query = ` SELECT user_id, restaurant_id, reservation_time, status from reservations 
			where delated_at=0 `
	var arr []interface{}
	count := 1
	if len(res.Status) > 0 {
		query += fmt.Sprintf(" and status=$%d", count)
		count++
		arr = append(arr, res.Status)
	}
	if len(res.ReservationTime) > 0 {
		query += fmt.Sprintf(" and reservation_time=$%d", count)
		count++
		arr = append(arr, res.ReservationTime)
	}

	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var Reservation pb.Reservation
		err := row.Scan(&Reservation.UserId, &Reservation.RestaurantId, &Reservation.ReservationTime, &Reservation.Status)
		if err != nil {
			return nil, err
		}

		Reservations.Reservations = append(Reservations.Reservations, &Reservation)
	}

	return Reservations, nil
}

func (p *Reservationstorage) UpdateReservation(reservation *pb.Reservation) (*pb.Void, error) {
	query := `
		UPDATE reservations
		SET user_id = $1, restaurant_id = $2, reservation_time = $3, status = $4
		WHERE id = $5
	`
	_, err := p.db.Exec(query, reservation.UserId, reservation.RestaurantId, reservation.ReservationTime, reservation.Status, reservation.Id)
	return nil, err
}
func (p *Reservationstorage) DeleteReservation(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE reservations
		SET delated_at = $1
		WHERE id = $2
	`
	_, err := p.db.Exec(query, time.Now().Unix(), id.Id)
	return nil, err
}
