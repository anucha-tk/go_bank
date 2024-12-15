package util

type CreateFakerUser struct {
	Username       string `faker:"username"`
	HashedPassword string `faker:"password"`
	FullName       string `faker:"name"`
	Email          string `faker:"email"`
}

type CreateFakerAccount struct {
	Currency string `faker:"oneof:USD,EUR,CAD"`
	Balance  int64  `faker:"boundary_start=31, boundary_end=88"`
}

type UpdateFakerAccount struct {
	ID      int64 `faker:"-"`
	Balance int64 `faker:"boundary_start=500, boundary_end=999"`
}
