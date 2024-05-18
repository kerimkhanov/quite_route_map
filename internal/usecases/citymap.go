package usecases

import "context"

func (uc *UseCase) GetMap(ctx context.Context) error {
	return uc.cr.CityMap.GetMap()
}
