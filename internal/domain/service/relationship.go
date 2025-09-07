package service

import (
	mRelationship "clone-instagram-service/internal/domain/model/relationship"
)

type relationshipService struct {
	relationshipRepo mRelationship.RelationshipRepository
}

func NewRelationshipService(relationshipRepo mRelationship.RelationshipRepository) *relationshipService {
	return &relationshipService{
		relationshipRepo: relationshipRepo,
	}
}
