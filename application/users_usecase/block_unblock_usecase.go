package users_usecase

func (u *usersUsecase) BlockUnblock(id string, isBlocked bool) error {
	err := u.repository.BlockUnblockUser(id, isBlocked)
	if err != nil {
		return err
	}

	return nil
}
