// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"testing"

	"github.com/golang/mock/gomock"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/fixtures"
	"github.com/mongodb/mongocli/internal/mocks"
)

func TestAtlasDBUserUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDatabaseUserStore(ctrl)

	defer ctrl.Finish()

	expected := fixtures.DatabaseUser()

	updateOpts := &atlasDBUsersUpdateOpts{
		globalOpts: newGlobalOpts(),
		username:   "test4",
		password:   "US",
		roles:      []string{"admin@admin"},
		store:      mockStore,
	}

	dbUser := atlas.DatabaseUser{}
	updateOpts.update(&dbUser)

	mockStore.
		EXPECT().
		UpdateDatabaseUser(&dbUser).
		Return(expected, nil).
		Times(1)

	err := updateOpts.Run()

	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}