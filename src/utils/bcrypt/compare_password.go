package bcrypt

import "golang.org/x/crypto/bcrypt"

func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) //* Si el password es correcto, err es nil
	/* if err != nil {
		return false
	}
	return true */

	return err == nil

}
