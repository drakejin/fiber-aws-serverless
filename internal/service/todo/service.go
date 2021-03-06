package todo

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/drakejin/fiber-aws-serverless/model"
)

func (s *Service) Insert(todo *model.Todo) (*model.Todo, error) {
	todo.Status = model.StatusIdle

	tx := s.Container.ServiceDB.DB.Create(todo)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return todo, nil
}

func (s *Service) Select() ([]*model.Todo, error) {
	m := new(model.Todo)
	tx := s.Container.ServiceDB.DB.
		Model(m).
		Where("status != ?", model.StatusRemoved)
	if tx.Error != nil {
		return nil, tx.Error
	}

	rows, err := tx.Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var todos []*model.Todo
	for rows.Next() {
		t := new(model.Todo)
		err := tx.ScanRows(rows, &t)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (s *Service) SelectOne(id string) (*model.Todo, error) {
	m := new(model.Todo)
	tx := s.Container.ServiceDB.DB.
		Model(m).
		Where("status != ? AND id = ?", model.StatusRemoved, id).
		Scan(m)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return m, nil
}

func (s *Service) Update(id string, todo *model.Todo) (*model.Todo, error) {
	t := new(model.Todo)
	err := s.Container.ServiceDB.DB.Transaction(func(tx *gorm.DB) error {
		subErr := tx.Model(todo).Where("id = ?", id).Scan(t).Error
		if subErr != nil {
			return subErr
		}
		if t.ID == "" {
			return gorm.ErrRecordNotFound
		}

		if todo.Name != "" {
			t.Name = todo.Name
		}
		if todo.Sort != 0 {
			t.Sort = todo.Sort
		}
		if todo.Note != "" {
			t.Note = todo.Note
		}
		if todo.Type != "" {
			t.Type = todo.Type
		}
		subErr = tx.Model(t).
			Updates(t).
			Error
		if subErr != nil {
			return subErr
		}
		return nil
	}, &sql.TxOptions{
		ReadOnly:  false,
	})

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *Service) Delete(id string) (*model.Todo, error) {
	t := new(model.Todo)
	err := s.Container.ServiceDB.DB.Transaction(func(tx *gorm.DB) error {
		subErr := tx.Model(t).Where("id = ?", id).Scan(t).Error
		if subErr != nil {
			return subErr
		}
		if t.ID == "" {
			return gorm.ErrRecordNotFound
		}

		if t.ID != "" {
			subErr = tx.Model(t).Update("status", model.StatusRemoved).Error
			return subErr
		}

		return nil
	}, &sql.TxOptions{
		ReadOnly:  false,
	})

	if err != nil {
		return nil, err
	}

	return t, nil
}
