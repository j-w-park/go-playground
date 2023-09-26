package dep_injection

type DataStore interface {
	GetUserNameFromId(id string) (string, bool)
}

type SimpleDataStore struct {
	users map[string]string
}

func (sds SimpleDataStore) GetUserNameFromId(id string) (string, bool) {
	user, ok := sds.users[id]
	return user, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		users: map[string]string{
			"1": "John",
			"2": "Jane",
		},
	}
}
