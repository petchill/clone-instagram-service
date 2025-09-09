package service

import (
	mRelationship "clone-instagram-service/internal/domain/model/relationship"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	relationshipRepo *mRelationship.MockRelationshipRepository
}

func Test_FollowUser(t *testing.T) {

	type Args struct {
		userID       string
		targetUserID string
	}

	type TestCase struct {
		name        string
		args        Args
		beforeTest  func(mockRepo mockRepo)
		expectedErr error
	}

	testCases := []TestCase{
		{
			name: "[negative]_GivenUserIDIsAlreadyFollowTargetUserID_WhenUserIDTryToFollowAgain_ThenReturnErrorBecauseItIsAlreadyFollowed",
			args: Args{
				userID:       "1",
				targetUserID: "2",
			},
			beforeTest: func(mockRepo mockRepo) {
				mockRepo.relationshipRepo.EXPECT().GetAllFollowingIDsByUserID(mock.Anything, "1").Return([]string{"2"}, nil)
			},
			expectedErr: fmt.Errorf("this following is already exists"),
		},
		{
			name: "[positive]_GivenUserIDNeverFollowTargetUserID_WhenUserIDTryToFollow_ThenRecordTheFollowingAndReturnSuccess",
			args: Args{
				userID:       "1",
				targetUserID: "2",
			},
			beforeTest: func(mockRepo mockRepo) {
				mockRepo.relationshipRepo.EXPECT().GetAllFollowingIDsByUserID(mock.Anything, "1").Return([]string{}, nil)
				mockRepo.relationshipRepo.EXPECT().InsertFollowing(mock.Anything, mRelationship.Following{
					UserId:       "1",
					TargetUserID: "2",
				}).Return(nil)
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {

		mockRelationshipRepo := new(mRelationship.MockRelationshipRepository)
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mockRepo{
				relationshipRepo: mockRelationshipRepo,
			}
			tc.beforeTest(mockRepo)

			relationshipService := &relationshipService{
				relationshipRepo: mockRepo.relationshipRepo,
			}

			err := relationshipService.FollowUser(context.Background(), tc.args.userID, tc.args.targetUserID)
			assert.Equal(t, tc.expectedErr, err)

		})
	}
}
