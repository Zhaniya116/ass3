package internal

import (
    "zhaniya.net/internal/usecase"
    "zhaniya.net/internal/repository"
)

type ContactUsecase struct {
    ContactService *ContactService
}

func NewContactUsecase(contactService *ContactService) *ContactUsecase {
    return &ContactUsecase{
        ContactService: contactService,
    }
}
