package database

import "errors"

func (n Database) AcquireAdvisoryLock(lockKey int64) error {
	result := n.DB.Exec("SELECT pg_advisory_lock(?)", lockKey)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("failed to acquire advisory lock")
	}

	return nil
}

func (n Database) ReleaseAdvisoryLock(lockKey int64) error {
	return n.DB.Exec("SELECT pg_advisory_unlock(?)", lockKey).Error
}
