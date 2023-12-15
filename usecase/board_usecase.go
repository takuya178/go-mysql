package usecase

import (
	"github.com/t-shimpo/go-mysql-docker/app/model"
	"github.com/t-shimpo/go-mysql-docker/repository"
)

type IBoardUsecase interface {
	GetAllBoards(userId uint) ([]model.BoardResponse, error)
	GetBoardById(userId uint, boardId uint) (model.BoardResponse, error)
	CreateBoard(board model.Board) (model.BoardResponse, error)
	UpdateBoard(boards model.Board, userId uint, boardId uint) (model.BoardResponse, error)
	DeleteBoard(userId uint, boardId uint) error
}

type boardUsecase struct {
	tr repository.IBoardRepository
}

func NewBoardUsecase(tr repository.IBoardRepository) IBoardUsecase {
	return &boardUsecase{tr}
}

func (tu *boardUsecase) GetAllBoards(userId uint) ([]model.BoardResponse, error) {
	boards := []model.Board{}
	if err := tu.tr.GetAllBoards(&boards, userId); err != nil {
		return nil, err
	}
	resBoards := []model.BoardResponse{}
	for _, v := range boards {
		t := model.BoardResponse{
			ID: 			 v.ID,
			Title: 		 v.Title,
			Context:   v.Context,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resBoards = append(resBoards, t)
	}
	return resBoards, nil
}

func (tu *boardUsecase) GetBoardById(userId uint, boardId uint) (model.BoardResponse, error) {
	board := model.Board{}
	if err := tu.tr.GetBoardById(&board, userId, boardId); err != nil {
		return model.BoardResponse{}, err
	}
	resBoard := model.BoardResponse {
		ID:        board.ID,
		Title:     board.Title,
		Context:   board.Context,
		CreatedAt: board.CreatedAt,
		UpdatedAt: board.UpdatedAt,		
	}
	return resBoard, nil
}

func (tu *boardUsecase) CreateBoard(board model.Board) (model.BoardResponse, error) {
	if err := tu.tr.CreateBoard(&board); err != nil {
		return model.BoardResponse{}, err
	}
	resBoard := model.BoardResponse {
		ID:        board.ID,
		Title:     board.Title,
		Context:   board.Context,
		CreatedAt: board.CreatedAt,
		UpdatedAt: board.UpdatedAt,
	}
	return resBoard, nil
}

func (tu *boardUsecase) UpdateBoard(board model.Board, userId uint, boardId uint) (model.BoardResponse, error) {
	if err := tu.tr.UpdateBoard(&board, userId, boardId); err != nil {
		return model.BoardResponse{}, err
	}
	resBoard := model.BoardResponse {
		ID:        board.ID,
		Title:     board.Title,
		Context:   board.Context,
		CreatedAt: board.CreatedAt,
		UpdatedAt: board.UpdatedAt,	
	}
	return resBoard, nil
}

func (tu *boardUsecase) DeleteBoard(userId uint, boardId uint) error {
	if err := tu.tr.DeleteBoard(userId, boardId); err != nil {
		return err
	}
	return nil
}
