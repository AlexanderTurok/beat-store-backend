package repository

import (
	"fmt"
	"strings"

	beatstore "github.com/AlexanderTurok/beat-store-backend/pkg"
	"github.com/jmoiron/sqlx"
)

type BeatRepository struct {
	db *sqlx.DB
}

func NewBeatRepository(db *sqlx.DB) *BeatRepository {
	return &BeatRepository{
		db: db,
	}
}

func (r *BeatRepository) Create(userId int, beat beatstore.Beat) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var beatId int
	beatQuery := fmt.Sprintf("INSERT INTO %s (bpm, key, path, tag, price) VALUES ($1, $2, $3, $4, $5) RETURNING id", beatTable)
	row := tx.QueryRow(beatQuery, beat.Bpm, beat.Key, beat.Path, beat.Tag, beat.Price)
	if err := row.Scan(&beatId); err != nil {
		tx.Rollback()
		return 0, err
	}

	usersBeatQuery := fmt.Sprintf("INSERT INTO %s (user_id, beat_id) VALUES ($1, $2)", usersBeatTable)
	_, err = tx.Exec(usersBeatQuery, userId, beatId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return beatId, tx.Commit()
}

func (r *BeatRepository) GetAll() ([]beatstore.Beat, error) {
	var beats []beatstore.Beat

	query := fmt.Sprintf("SELECT * FROM %s", beatTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var beat beatstore.Beat
		if err := rows.Scan(&beat.Id, &beat.Bpm, &beat.Key, &beat.Path, &beat.Tag, &beat.Price); err != nil {
			return beats, err
		}
		beats = append(beats, beat)
	}

	return beats, rows.Err()
}

func (r *BeatRepository) Update(userId, beatId int, input beatstore.BeatUpdateInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Bpm != nil {
		setValues = append(setValues, fmt.Sprintf("bpm=$%d", argId))
		args = append(args, *input.Bpm)
		argId++
	}

	if input.Key != nil {
		setValues = append(setValues, fmt.Sprintf("key=$%d", argId))
		args = append(args, *input.Key)
		argId++
	}

	if input.Path != nil {
		setValues = append(setValues, fmt.Sprintf("path=$%d", argId))
		args = append(args, *input.Path)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.Tag != nil {
		setValues = append(setValues, fmt.Sprintf("tag=$%d", argId))
		args = append(args, *input.Tag)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s bt SET %s FROM %s ub WHERE bt.id = ub.beat_id AND ub.beat_id=$%d AND ub.user_id=$%d",
		beatTable, setQuery, usersBeatTable, argId, argId+1)
	args = append(args, beatId, userId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *BeatRepository) Delete(userId, beatId int) error {
	query := fmt.Sprintf("DELETE FROM %s bt USING %s ub WHERE bt.id = ub.beat_id AND ub.user_id=$1 AND ub.beat_id=$2",
		beatTable, usersBeatTable)
	_, err := r.db.Exec(query, userId, beatId)

	return err
}
