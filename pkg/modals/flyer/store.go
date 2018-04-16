package flyer

import (
	"context"
	"database/sql"
	"time"

	"code.ysitd.cloud/common/go/db"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

const (
	defaultExpire = 5 * time.Minute
	cleanupExpire = 10 * time.Minute
)

type Store struct {
	Cache    *cache.Cache
	Database db.Pool
	Logger   logrus.FieldLogger
}

func NewStore(database db.Pool, logger logrus.FieldLogger) *Store {
	c := cache.New(defaultExpire, cleanupExpire)
	return &Store{
		Cache:    c,
		Database: database,
		Logger:   logger,
	}
}

func (s *Store) GetFlyer(ctx context.Context, flightNumber string) (flyer *Flyer, err error) {
	flyer, cached := s.getCachedFlyer(flightNumber)
	if cached {
		return
	}

	return s.getDBFlyer(ctx, flightNumber)
}

func (s *Store) getCachedFlyer(flightNumber string) (flyer *Flyer, cached bool) {
	val, cached := s.Cache.Get(flightNumber)
	if cached {
		flyer = val.(*Flyer)
	}

	return
}

func (s *Store) getDBFlyer(ctx context.Context, flightNumber string) (flyer *Flyer, err error) {
	conn, err := s.Database.Acquire()
	if err != nil {
		return
	}

	defer conn.Close()

	s.Logger.Debugf("Get info for hostname %s", flightNumber)

	query := "SELECT id, revision FROM flyers WHERE hostname = $1"
	row := conn.QueryRowContext(ctx, query, flightNumber)

	var f Flyer
	flyer = &f
	f.Hostname = flightNumber

	if err = row.Scan(&f.ID, &f.Revision); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return
}
