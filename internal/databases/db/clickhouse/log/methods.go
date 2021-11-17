package log

import "github.com/gna69/grpc-users/internal/databases/model"

func (s *service) Create(log *model.LogRecord) error {
	t := s.db.MustBegin()
	_, err := t.NamedExec(`insert into logs (first_name, last_name, user_id, action_day, action_time) values 
					(:first_name, :last_name, :user_id, :action_day, :action_time)`, log)
	t.Commit()

	if err != nil {
		return err
	}
	return nil
}
