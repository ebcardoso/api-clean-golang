package users_usecase

func (u *usersUsecase) Destroy(id string) error {
	err := u.repository.DestroyUser(id)
	if err != nil {
		return err
	}

	return nil
}
