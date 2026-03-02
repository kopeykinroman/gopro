package storage

type Storage struct {
	id      uint64
	idToUrl map[uint64]string
	urlToId map[string]uint64
}

func NewStorage() *Storage {
	return &Storage{
		idToUrl: make(map[uint64]string),
		urlToId: make(map[string]uint64),
	}
}

// Add добавляет в хранилище новый url
func (s *Storage) Add(url string) (shortLink string) {
	// если url был добавлен раннее в хранилище
	if id, ok := s.urlToId[url]; ok {
		return Encode(id)
	}

	// иначе добавляем
	s.id++
	s.idToUrl[s.id] = url
	s.urlToId[url] = s.id

	return Encode(s.id)
}

// Get возвращает из хранилища url по коротко ссылке
func (s *Storage) Get(shortLink string) (string, error) {

	id, err := Decode(shortLink)
	if err != nil {
		return "", ErrInvalidShort
	}

	if v, ok := s.idToUrl[id]; ok {
		return v, nil
	}

	return "", ErrNotFound

}
