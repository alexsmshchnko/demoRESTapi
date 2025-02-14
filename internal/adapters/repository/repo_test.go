package repository

import (
	"demorestapi/internal/entity"
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"

	"demorestapi/internal/adapters/mock_postgres"
)

func TestRepo_GetUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existingUser := &entity.User{
		ID:        "123",
		Firstname: "name",
		Lastname:  "name2",
		Email:     "example.com",
		Age:       18,
	}

	unknownId := "000"

	p := mock_postgres.NewMockDataProvider(ctrl)
	p.EXPECT().GetUser(existingUser.ID).Return(existingUser).Times(1)
	p.EXPECT().GetUser(unknownId).Return(&entity.User{}).Times(1)

	r := NewRepo(p)

	type args struct {
		id string
	}
	tests := []struct {
		name        string
		args        args
		wantRes     *entity.User
		cacheBefore bool
		cacheAfter  bool
	}{
		{
			name:        "user exists",
			args:        args{id: existingUser.ID},
			wantRes:     existingUser,
			cacheBefore: false,
			cacheAfter:  true,
		},
		{
			name:        "no user",
			args:        args{id: unknownId},
			wantRes:     &entity.User{},
			cacheBefore: false,
			cacheAfter:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := r.localCache.users[tt.args.id]
			if ok != tt.cacheBefore {
				t.Errorf("Repo.localCache: %v", r.localCache)
			}

			if gotRes := r.GetUser(tt.args.id); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Repo.GetUser() = %v, want %v", gotRes, tt.wantRes)
			}

			_, ok = r.localCache.users[tt.args.id]
			if ok != tt.cacheAfter {
				t.Errorf("Repo.localCache: %v", r.localCache)
			}
		})
	}
}

func TestRepo_AddUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newUser := &entity.User{
		Firstname: "name",
		Lastname:  "name2",
		Email:     "example.com",
		Age:       18,
	}

	commitedUserID := "101"

	p := mock_postgres.NewMockDataProvider(ctrl)
	p.EXPECT().AddUser(newUser).DoAndReturn(func(user *entity.User) (err error) {
		user.ID = commitedUserID
		return nil
	})

	r := NewRepo(p)

	type args struct {
		u *entity.User
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		cacheAfter bool
	}{
		{
			name:       "demo1",
			args:       args{u: newUser},
			wantErr:    false,
			cacheAfter: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.AddUser(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("Repo.AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if _, ok := r.localCache.users[commitedUserID]; ok != tt.cacheAfter {
				t.Errorf("Repo.localCache: %v", r.localCache)
			}
		})
	}
}

func TestRepo_UpdateUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	updUser := &entity.User{
		ID:        "101",
		Firstname: "name",
		Lastname:  "name2",
		Email:     "example.com",
		Age:       20,
	}

	p := mock_postgres.NewMockDataProvider(ctrl)
	p.EXPECT().UpdateUser(updUser).Return(nil).Times(1)

	r := NewRepo(p)

	type args struct {
		u *entity.User
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		cacheAfter bool
	}{
		{
			name:       "demo1",
			args:       args{u: updUser},
			wantErr:    false,
			cacheAfter: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.UpdateUser(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("Repo.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, ok := r.localCache.users[updUser.ID]; ok != tt.cacheAfter {
				t.Errorf("Repo.localCache: %v", r.localCache)
			}
		})
	}
}
