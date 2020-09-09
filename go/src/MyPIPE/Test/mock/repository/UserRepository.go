package test

import (
	"MyPIPE/domain/model"
	"time"
)

type UserRepositoryMock struct{}

func (u *UserRepositoryMock)GetAll() []model.User{
	user1 := model.NewUser()
	user2 := model.NewUser()
	user3 := model.NewUser()

	timeNow := time.Now()

	user1.ID = 1
	user1.Name = "テスト太郎"
	user1.Password = "eifsjzefsajafla28739812"
	user1.Email = "taro@example.jp"
	user1.Birthday = time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	user1.Token = ""
	user1.CreatedAt = timeNow
	user1.UpdatedAt = timeNow
	movie1 := GetMovie(1,user1)
	movie2 := GetMovie(2,user1)
	movie3 := GetMovie(3,user1)
	user1.Movies = []model.Movie{*movie1, *movie2, *movie3}

	user2.ID = 2
	user2.Name = "テスト次郎"
	user2.Password = "jfoeurnnfhfa87923"
	user2.Email = "taro@example.jp"
	user2.Birthday = time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local)
	user2.Token = ""
	user2.CreatedAt = timeNow
	user2.UpdatedAt = timeNow
	movie4 := GetMovie(4,user2)
	movie5 := GetMovie(5,user2)
	movie6 := GetMovie(6,user2)
	user2.Movies = []model.Movie{*movie4, *movie5, *movie6}

	user3.ID = 3
	user3.Name = "テスト三郎"
	user3.Password = "feiwrqodfhakjlsd27983"
	user3.Email = "taro@example.jp"
	user3.Birthday = time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	user3.Token = ""
	user3.CreatedAt = timeNow
	user3.UpdatedAt = timeNow
	movie7 := GetMovie(7,user3)
	movie8 := GetMovie(8,user3)
	movie9 := GetMovie(9,user3)
	user1.Movies = []model.Movie{*movie7, *movie8, *movie9}

	return []model.User{*user1,*user2,*user3}
}

func (u *UserRepositoryMock)FindById(id uint64) *model.User{
	user := model.NewUser()
	timeNow := time.Now()
	user.ID = id
	user.Name = "Test太郎"
	user.Password = "fjeoisafj9384299042"
	user.Email = "TARO@example.jp"
	user.Birthday = time.Date(1995, 1, 1, 0, 0, 0, 0, time.Local)
	user.Token = ""
	user.CreatedAt = timeNow
	user.UpdatedAt = timeNow
	movie1 := GetMovie(1,user)
	movie2 := GetMovie(2,user)
	movie3 := GetMovie(3,user)
	user.Movies = []model.Movie{*movie1, *movie2, *movie3}

	return user
}

func (u *UserRepositoryMock)FindByToken(token string) *model.User{
	user := model.NewUser()
	timeNow := time.Now()
	user.ID = 1
	user.Name = "Test太郎"
	user.Password = "fjeoisafj9384299042"
	user.Email = "TARO@example.jp"
	user.Birthday = time.Date(1995, 1, 1, 0, 0, 0, 0, time.Local)
	user.Token = token
	user.CreatedAt = timeNow
	user.UpdatedAt = timeNow
	movie1 := GetMovie(1,user)
	movie2 := GetMovie(2,user)
	movie3 := GetMovie(3,user)
	user.Movies = []model.Movie{*movie1, *movie2, *movie3}

	return user
}

func (u *UserRepositoryMock)FindByEmail(email string) *model.User{
	user := model.NewUser()
	timeNow := time.Now()
	user.ID = 1
	user.Name = "Test太郎"
	user.Password = "fjeoisafj9384299042"
	user.Email = email
	user.Birthday = time.Date(1995, 1, 1, 0, 0, 0, 0, time.Local)
	user.Token = ""
	user.CreatedAt = timeNow
	user.UpdatedAt = timeNow
	movie1 := GetMovie(1,user)
	movie2 := GetMovie(2,user)
	movie3 := GetMovie(3,user)
	user.Movies = []model.Movie{*movie1, *movie2, *movie3}

	return user
}

func (u *UserRepositoryMock)SetUser(user *model.User){

}

func GetMovie(movieId uint64,user *model.User) *model.Movie{
	movie1 := model.NewMovie()

	movie1.ID = movieId
	movie1.StoreName = "my_movie1.mp4"
	movie1.DisplayName = "歌ってみた"
	movie1.UserID = user.ID
	movie1.User = *user
	movie1.CreatedAt = time.Now()
	movie1.UpdatedAt = time.Now()

	return movie1

}