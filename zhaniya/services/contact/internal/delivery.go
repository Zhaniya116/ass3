package internal

import (
    "zhaniya.net/internal/delivery"
    "zhaniya.net/internal/usecase"
)

type ContactDelivery struct {
    ContactUsecase *usecase.ContactUsecase
}

func NewContactDelivery(contactUsecase *usecase.ContactUsecase) *ContactDelivery {
    return &ContactDelivery{
        ContactUsecase: contactUsecase,
    }
}
