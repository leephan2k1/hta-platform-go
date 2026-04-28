package service

import (
	"context"
	"fmt"
	"hta-platform/internal/media/controller/dto"
	"hta-platform/internal/media/domain/model/entity"
	"hta-platform/internal/media/domain/repository"
	"hta-platform/utils"
	"strings"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type mediaService struct {
	mediaRepo repository.MediaRepository
	db        *gorm.DB
}

// GetMediaByUrl implements [MediaService].
func (m *mediaService) GetMediaByUrl(ctx context.Context, url string) (dto.MediaResponse, error) {
	media, err := m.mediaRepo.GetMediaByUrl(ctx, url)
	if err != nil {
		return dto.MediaResponse{}, err
	}

	var res dto.MediaResponse
	res.SetData(media)

	return res, nil
}

// GetMedias implements [MediaService].
func (m *mediaService) GetMedias(ctx context.Context, req *dto.GetMediasReq) (dto.GetMediasRes, error) {
	items, total, err := m.mediaRepo.GetMedias(ctx, req)
	if err != nil {
		return dto.GetMediasRes{}, err
	}

	var res dto.GetMediasRes
	res.SetPagination(total, req.Page, req.Limit)

	res.Items = make([]dto.MediaResponse, len(items))
	for i, media := range items {
		res.Items[i].SetData(media)
	}

	return res, nil
}

// CreateMedia implements [MediaService].
func (m *mediaService) CreateMedia(ctx context.Context, req *dto.CreateMediaReq) (entity.Media, error) {
	nameVal := strings.TrimSpace(req.Name)
	slugVal := slug.Make(strings.ToLower(nameVal))

	// Parse UUIDs from string
	statusID, err := uuid.Parse(req.StatusID)
	if err != nil {
		return entity.Media{}, fmt.Errorf("invalid statusId: %w", err)
	}
	typeID, err := uuid.Parse(req.TypeID)
	if err != nil {
		return entity.Media{}, fmt.Errorf("invalid typeId: %w", err)
	}

	var created *entity.Media

	txErr := m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. Build Media entity
		media := &entity.Media{
			Name:        nameVal,
			Description: req.Description,
			URL:         slugVal,
			StatusID:    statusID,
			TypeID:      typeID,
			IsNSFW:      req.IsNSFW,
			Thumbnail:   req.Thumbnail,
		}

		// 2. Insert media (ON CONFLICT DO NOTHING on url)
		insertedMedia, ok, err := m.mediaRepo.CreateMedia(tx, media)
		if err != nil {
			return err
		}
		if !ok {
			// Duplicate slug — return nil to commit tx, created stays nil
			return nil
		}

		// 3. Other Name
		if req.OtherName != "" {
			otherName := &entity.MediaOtherName{
				Name:     req.OtherName,
				Language: req.OtherNameLanguage,
				MediaID:  insertedMedia.ID,
			}
			if err := m.mediaRepo.InsertOtherName(tx, otherName); err != nil {
				return err
			}
		}

		// 4. Author (M2M)
		if len(req.AuthorIDs) > 0 {
			authorUUIDs, err := utils.ParseUUIDs(req.AuthorIDs)
			if err != nil {
				return fmt.Errorf("invalid authorIds: %w", err)
			}
			if err := m.mediaRepo.AttachAuthors(tx, insertedMedia.ID, authorUUIDs); err != nil {
				return err
			}
		}

		// 5. Category (M2M)
		if len(req.CategoryIDs) > 0 {
			categoryUUIDs, err := utils.ParseUUIDs(req.CategoryIDs)
			if err != nil {
				return fmt.Errorf("invalid categoryIds: %w", err)
			}
			if err := m.mediaRepo.AttachCategories(tx, insertedMedia.ID, categoryUUIDs); err != nil {
				return err
			}
		}

		created = insertedMedia
		return nil
	})

	if txErr != nil {
		return entity.Media{}, txErr
	}

	if created == nil {
		return entity.Media{}, fmt.Errorf("media with slug '%s' already exists", slug.Make(strings.ToLower(nameVal)))
	}

	return *created, nil
}

// UpdateMedia implements [MediaService].
// Follows the TS pattern: update by URL, then replace M2M relations.
func (m *mediaService) UpdateMedia(ctx context.Context, url string, req *dto.CreateMediaReq) (entity.Media, error) {
	nameVal := strings.TrimSpace(req.Name)
	slugVal := slug.Make(strings.ToLower(nameVal))

	statusID, err := uuid.Parse(req.StatusID)
	if err != nil {
		return entity.Media{}, fmt.Errorf("invalid statusId: %w", err)
	}
	typeID, err := uuid.Parse(req.TypeID)
	if err != nil {
		return entity.Media{}, fmt.Errorf("invalid typeId: %w", err)
	}

	var updated *entity.Media

	txErr := m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. Update media by URL
		media := &entity.Media{
			Name:        nameVal,
			Description: req.Description,
			URL:         slugVal,
			StatusID:    statusID,
			TypeID:      typeID,
			IsNSFW:      req.IsNSFW,
			Thumbnail:   req.Thumbnail,
		}

		updatedMedia, err := m.mediaRepo.UpdateMediaByURL(tx, url, media)
		if err != nil {
			return err
		}

		// 2. Replace other names
		if err := m.mediaRepo.DeleteOtherNamesByMediaID(tx, updatedMedia.ID); err != nil {
			return err
		}
		if req.OtherName != "" {
			otherName := &entity.MediaOtherName{
				Name:     req.OtherName,
				Language: req.OtherNameLanguage,
				MediaID:  updatedMedia.ID,
			}
			if err := m.mediaRepo.InsertOtherName(tx, otherName); err != nil {
				return err
			}
		}

		// 3. Replace authors
		if err := m.mediaRepo.DeleteAuthorsByMediaID(tx, updatedMedia.ID); err != nil {
			return err
		}
		if len(req.AuthorIDs) > 0 {
			authorUUIDs, err := utils.ParseUUIDs(req.AuthorIDs)
			if err != nil {
				return fmt.Errorf("invalid authorIds: %w", err)
			}
			if err := m.mediaRepo.AttachAuthors(tx, updatedMedia.ID, authorUUIDs); err != nil {
				return err
			}
		}

		// 4. Replace categories
		if err := m.mediaRepo.DeleteCategoriesByMediaID(tx, updatedMedia.ID); err != nil {
			return err
		}
		if len(req.CategoryIDs) > 0 {
			categoryUUIDs, err := utils.ParseUUIDs(req.CategoryIDs)
			if err != nil {
				return fmt.Errorf("invalid categoryIds: %w", err)
			}
			if err := m.mediaRepo.AttachCategories(tx, updatedMedia.ID, categoryUUIDs); err != nil {
				return err
			}
		}

		updated = updatedMedia
		return nil
	})

	if txErr != nil {
		return entity.Media{}, txErr
	}

	return *updated, nil
}

func NewMediaService(mediaRepo repository.MediaRepository, db *gorm.DB) MediaService {
	return &mediaService{
		mediaRepo: mediaRepo,
		db:        db,
	}
}
