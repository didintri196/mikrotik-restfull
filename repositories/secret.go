package repositories

import (
	"github.com/didintri196/mikorm"
	"github.com/didintri196/mikrotik-restfull/domain/interfaces"
	"github.com/didintri196/mikrotik-restfull/domain/models"
)

type SecretRepository struct {
	MikORM *mikorm.MikORM
}

func (repo SecretRepository) Browse(filter models.Secret) (secrets []models.Secret, err error) {
	err = repo.MikORM.Command("/ppp/secret").Where(&filter).Scan(&secrets).Error
	return
}

func (repo SecretRepository) Add(secret models.Secret) (err error) {
	return repo.MikORM.Command("/ppp/secret").Add(&secret).Error
}

func (repo SecretRepository) Read(filter models.Secret) (secret models.Secret, err error) {
	err = repo.MikORM.Command("/ppp/secret").Where(&filter).Print(&secret).Error
	return
}

func (repo SecretRepository) Edit(filter models.Secret, data models.Secret) (err error) {
	err = repo.MikORM.Command("/ppp/secret").SetByID("", &data).Error
	return
}

func (repo SecretRepository) Remove(ID string) (err error) {
	err = repo.MikORM.Command("/ppp/secret").RemoveByID(ID).Error
	return
}

func (repo SecretRepository) Enable(ID string) (err error) {
	err = repo.MikORM.Command("/ppp/secret").EnableByID(ID).Error
	return
}

func (repo SecretRepository) Disable(ID string) (err error) {
	err = repo.MikORM.Command("/ppp/secret").DisableByID(ID).Error
	return
}

func NewSecretRepository(mikORM *mikorm.MikORM) interfaces.ISecretRepository {
	return &SecretRepository{
		MikORM: mikORM,
	}
}
