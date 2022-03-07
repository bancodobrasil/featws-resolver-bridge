package services

import (
	"context"

	"github.com/bancodobrasil/featws-resolver-bridge/models"
	"github.com/bancodobrasil/featws-resolver-bridge/repository"
)

// CreateResolver ...
func CreateResolver(ctx context.Context, resolver *models.Resolver) (err error) {

	//TODO verifica unicidade do nome

	err = repository.GetResolversRepository().Create(ctx, resolver)
	if err != nil {
		return
	}

	return
}

// FetchResolvers ...
func FetchResolvers(ctx context.Context) (result []models.Resolver, err error) {

	result, err = repository.GetResolversRepository().Find(ctx)
	if err != nil {
		return
	}

	return
}

// FetchResolver ...
func FetchResolver(ctx context.Context, id string) (result *models.Resolver, err error) {

	result, err = repository.GetResolversRepository().Get(ctx, id)
	if err != nil {
		return
	}

	return
}

// UpdateResolver ...
func UpdateResolver(ctx context.Context, entity models.Resolver) (result *models.Resolver, err error) {

	result, err = repository.GetResolversRepository().Update(ctx, entity)
	if err != nil {
		return
	}

	return
}

// DeleteResolver ...
func DeleteResolver(ctx context.Context, id string) (deleted bool, err error) {

	deleted, err = repository.GetResolversRepository().Delete(ctx, id)
	if err != nil {
		return
	}

	return
}

