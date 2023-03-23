package usecase

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type UserLikeInteractor struct {
	UserLikeRepository UserLikeRepository
}

func (interactor *UserLikeInteractor) Add(u domain.UserLike) (domain.UserLike, error) {
	userLike, err := interactor.UserLikeRepository.Store(u)
	if err != nil {
		return domain.UserLike{}, err
	}
	return userLike, nil
}

func (interactor *UserLikeInteractor) UserLikes() (domain.UserLikes, error) {
	userLikes, err := interactor.UserLikeRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return userLikes, nil
}

func (interactor *UserLikeInteractor) UserLikeByUserID(userID uuid.UUID) (domain.UserLike, error) {
	userLike, err := interactor.UserLikeRepository.FindByUserID(userID)
	if err != nil {
		return domain.UserLike{}, err
	}
	return userLike, nil
}

func (interactor *UserLikeInteractor) UserLikeByTextID(textID uuid.UUID) (domain.UserLike, error) {
	userLike, err := interactor.UserLikeRepository.FindByTextID(textID)
	if err != nil {
		return domain.UserLike{}, err
	}
	return userLike, nil
}

func (interactor *UserLikeInteractor) Change(u domain.UserLike) error {
	err := interactor.UserLikeRepository.Update(u)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *UserLikeInteractor) Delete(u domain.UserLike) error {
	err := interactor.UserLikeRepository.Delete(u)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *UserLikeInteractor) DeleteByUserID(userID uuid.UUID) error {
	err := interactor.UserLikeRepository.DeleteByUserID(userID)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *UserLikeInteractor) DeleteByTextID(textID uuid.UUID) error {
	err := interactor.UserLikeRepository.DeleteByTextID(textID)
	if err != nil {
		return err
	}
	return nil
}
