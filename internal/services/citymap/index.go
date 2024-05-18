package citymap

type CityMap struct {
	cityMapRepository Repository
}

func New(cityMapRepository Repository) *CityMap {
	return &CityMap{
		cityMapRepository: cityMapRepository}
}
