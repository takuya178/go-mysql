package repository

import (
	"fmt"

	"github.com/t-shimpo/go-mysql-docker/app/model"
	"gorm.io/gorm"
)

type IBoardRepository interface {
	GetAllBoards(boards *[]model.Board, userId uint) error
	GetBoardById(board *model.Board, userId uint, boardId uint) error
	CreateBoard(board *model.Board) error
	UpdateBoard(boards *model.Board, userId uint, boardId uint) error
	DeleteBoard(userId uint, boardId uint) error
}

type boardRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) IBoardRepository {
	return &boardRepository{db}
}

func (tr *boardRepository) GetAllBoards(boards *[]model.Board, userId uint) error {
	query := `
		SELECT *
		FROM boards AS b
		INNER JOIN users u ON b.user_id = u.di
		WHERE u.id = ?
		ORDER BY b.created_at
	`
	if err := tr.db.Raw(query, userId).Scan(&boards).Error; err != nil {
		return err
	}
	return nil
}

func (tr *boardRepository) GetBoardById(board *model.Board, userId uint, boardId uint) error {
	query := `
		SELECT *
		FROM boards AS b
		INNER JOIN users u ON b.user_id = u.id
		WHERE u.id = ? AND b.id = ?
		ORDER BY b.created_at
	`
	if err := tr.db.Raw(query, userId, boardId).Scan(&board).Error; err != nil {
		return err
	}
	return nil
}

func (tr *boardRepository) CreateBoard(board *model.Board) error {
	if err := tr.db.Create(board).Error; err != nil {
		return err
	}
	return nil
}

func (tr *boardRepository) UpdateBoard(board *model.Board, userId uint, boardId uint) error {
	query := `
		SELECT *
		FROM boards AS b
		INNER JOIN users u ON b.user_id = u.id
		WHERE b.id = ? AND b.user_id = ?
	`
	result := tr.db.Raw(query, boardId, userId).Update("title", board.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exists")
	}
	return nil
}

func (tr *boardRepository) DeleteBoard(userId uint, boardId uint) error {
	return nil
}
