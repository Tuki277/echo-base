package repository

import (
	"echo-base/internal/contract"
)

type ModelRepository[T any, id int | string, query contract.QueryRequest] interface {
	Read(id) (*T, error)
	GetByName(string) (*T, error)
	List(query) ([]*T, int64, error)
	GetByIdList(ids []int) ([]*T, error)
	Insert(*T) (*T, error)
	Update(*T) (*T, error)
	Delete(*T) error
	DeleteAList([]uint) error
}
type NewModelRepository[T any, id int | string, query contract.QueryRequest] interface {
	Insert(*T) (*T, error)
	ListInAdmin(query) ([]*T, int, int, error)
	List(query) ([]*T, int, int, error)
	Read(int) (*T, error)
	Update(*T) (*T, error)
	UpdateFields(*T) (*T, error)
	Delete([]uint) error
}

type ModelAuthRepository[T any, id int | string, query contract.QueryRequest] interface {
	GetByEmail(string) (*T, error)
	GetById(id uint) (*T, error)
	Add(*T) error
	List() ([]*T, error)
	ChangeStatus(request *contract.ChangeStatusRes) error
	DeleteAList([]uint) error
}

type ModelCandidateRepository[T any, id int | string, query contract.QueryRequest] interface {
	Read(id) (*T, error)
	List(query) ([]*T, int64, error)
	Insert(*T) (*T, error)
	Update(*T) (*T, error)
	Delete(*T) error
	UpdateByMap(id, map[string]interface{}) (*T, error)
}

type ModelPostRepository[T any, id int | string, query contract.QueryRequest] interface {
	Read(id) (*T, error)
	ReadByConditions(id, map[string]interface{}) (*T, error)
	List(query) ([]*T, int64, error)
	ListByConditions(query, map[string]interface{}) ([]*T, int64, error)
	Insert(*T) (*T, error)
	Update(*T) (*T, error)
	Delete(*T) error
	UpdateByMap(id, map[string]interface{}) (*T, error)
}
