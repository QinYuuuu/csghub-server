package database

import (
	"context"
	"fmt"
	"strings"
)

type mirrorSourceStoreImpl struct {
	db *DB
}

type MirrorSourceStore interface {
	Create(ctx context.Context, mirrorSource *MirrorSource) (*MirrorSource, error)
	Index(ctx context.Context) ([]MirrorSource, error)
	Get(ctx context.Context, id int64) (*MirrorSource, error)
	FindByName(ctx context.Context, name string) (*MirrorSource, error)
	Update(ctx context.Context, mirrorSource *MirrorSource) (err error)
	Delete(ctx context.Context, mirrorSource *MirrorSource) (err error)
}

func NewMirrorSourceStore() MirrorSourceStore {
	return &mirrorSourceStoreImpl{
		db: defaultDB,
	}
}

func NewMirrorSourceStoreWithDB(db *DB) MirrorSourceStore {
	return &mirrorSourceStoreImpl{
		db: db,
	}
}

type MirrorSource struct {
	ID         int64  `bun:",pk,autoincrement" json:"id"`
	SourceName string `bun:",notnull,unique" json:"source_name"`
	InfoAPIUrl string `bun:",nullzero" json:"info_api_url"`

	times
}

func (s *mirrorSourceStoreImpl) Create(ctx context.Context, mirrorSource *MirrorSource) (*MirrorSource, error) {
	err := s.db.Operator.Core.NewInsert().
		Model(mirrorSource).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return mirrorSource, nil
}

func (s *mirrorSourceStoreImpl) Index(ctx context.Context) ([]MirrorSource, error) {
	var mirrorSources []MirrorSource
	err := s.db.Operator.Core.NewSelect().
		Model(&mirrorSources).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return mirrorSources, nil
}

func (s *mirrorSourceStoreImpl) Get(ctx context.Context, id int64) (*MirrorSource, error) {
	var mirrorSource MirrorSource
	err := s.db.Operator.Core.NewSelect().
		Model(&mirrorSource).
		Where("id=?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &mirrorSource, nil
}

func (s *mirrorSourceStoreImpl) FindByName(ctx context.Context, name string) (*MirrorSource, error) {
	var mirrorSource MirrorSource
	err := s.db.Operator.Core.NewSelect().
		Model(&mirrorSource).
		Where("source_name=?", name).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &mirrorSource, nil
}

func (s *mirrorSourceStoreImpl) Update(ctx context.Context, mirrorSource *MirrorSource) (err error) {
	err = assertAffectedOneRow(s.db.Operator.Core.NewUpdate().
		Model(mirrorSource).
		WherePK().
		Exec(ctx),
	)

	return
}

func (s *mirrorSourceStoreImpl) Delete(ctx context.Context, mirrorSource *MirrorSource) (err error) {
	_, err = s.db.Operator.Core.
		NewDelete().
		Model(mirrorSource).
		WherePK().
		Exec(ctx)
	return
}

func (m MirrorSource) BuildCloneURL(url, repoType, namespace, name string) string {
	namespace, _ = strings.CutPrefix(namespace, m.SourceName)
	return fmt.Sprintf("%s/%ss/%s/%s.git", url, repoType, namespace, name)
}
