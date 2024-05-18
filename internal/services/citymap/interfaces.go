package citymap

type (
	Repository interface {
		GetMap() error
	}
)

func (cr *CityMap) GetMap() error {
	err := cr.cityMapRepository.GetMap()
	if err != nil {
		return err
	}
	return nil
}
