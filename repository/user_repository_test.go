package repository

// import (
// 	"testing"

// 	"github.com/leeeeeoy/go_study/dto"
// 	"github.com/leeeeeoy/go_study/ent"
// 	"github.com/leeeeeoy/go_study/ent/enttest"
// 	"github.com/leeeeeoy/go_study/ent/migrate"
// 	_ "github.com/mattn/go-sqlite3"
// )

// func getClient(t *testing.T) *ent.Client {
// 	opts := []enttest.Option{
// 		enttest.WithOptions(ent.Log(t.Log)),
// 		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
// 	}
// 	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)

// 	return client
// }

// func TestCreateUser(t *testing.T) {
// 	client := getClient(t)
// 	defer client.Close()

// 	ur := NewUserRepository(client)

// 	tests := []struct {
// 		name     string
// 		input    *dto.UserRequest
// 		expected *dto.UserResponse
// 	}{
// 		{
// 			"에러, 너무 긴 이름",
// 			&dto.UserRequest{Name: "너무 긴 이름이어서 들어가지 않아야 됩니다.", Email: "Email", Password: "Password"},
// 			&dto.UserResponse{Name: "너무 긴 이름이어서 들어가지 않아야 됩니다.", Email: "Email"},
// 		},
// 		{
// 			"에러, 너무 긴 이메일",
// 			&dto.UserRequest{Name: "Name", Email: "너무 긴 이메일입니다. 따라서 생성할 때 들어가지 않아야 됩니다.", Password: "Password"},
// 			&dto.UserResponse{Name: "Name", Email: "너무 긴 이메일입니다. 따라서 생성할 때 들어가지 않아야 됩니다."},
// 		},
// 		{
// 			"에러, 너무 짧은 비밀번호",
// 			&dto.UserRequest{Name: "Name", Email: "Email", Password: "1234"},
// 			&dto.UserResponse{Name: "Name", Email: "Email"},
// 		},
// 		{
// 			"성공",
// 			&dto.UserRequest{Name: "Name", Email: "Email", Password: "Password"},
// 			&dto.UserResponse{Name: "Name", Email: "Email"},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			output, err := ur.CreateUser(tt.input)

// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			if tt.expected.Name != output.Name {
// 				t.Fatal("이름이 같지 않아요")
// 			}

// 			if tt.expected.Email != output.Email {
// 				t.Fatal("이메일이 같지 않아요")
// 			}

// 		})
// 	}

// }

// func TestGetUserByEmail(t *testing.T) {
// 	client := getClient(t)
// 	defer client.Close()

// 	ur := NewUserRepository(client)
// 	ur.CreateUser(&dto.UserRequest{Name: "Name", Email: "Email", Password: "Password"})

// 	tests := []struct {
// 		name     string
// 		input    string
// 		expected *dto.UserResponse
// 	}{
// 		{
// 			"실패, 조회 실패",
// 			"에베베베",
// 			&dto.UserResponse{Name: "Name", Email: "Email"},
// 		},

// 		{
// 			"성공",
// 			"Email",
// 			&dto.UserResponse{Name: "Name", Email: "Email"},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			output, err := ur.GetUserByEmail(tt.input)

// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			if tt.expected.Name != output.Name {
// 				t.Fatal("이름이 같지 않아요")
// 			}

// 			if tt.expected.Email != output.Email {
// 				t.Fatal("이메일이 같지 않아요")
// 			}

// 		})
// 	}
// }
