package repositories

import (
	"fmt"
	"moysklad/core/entites"
	"moysklad/infr"
	"moysklad/infr/models"
)

type AssortmentRepository struct {
}

func (a AssortmentRepository) Create(e *entites.Assortment) {
	result := infr.Db.Create(models.AssortmentModel{
		AccountId: e.AccountId,
	})

	fmt.Println(result.Error)
}

func (a AssortmentRepository) FindAll() {
	infr.Db.Find(models.AssortmentModel{})
}
