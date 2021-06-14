package repository

import (
	"errors"
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
)

type ReviewRepo struct {
	reviews map[int64]*model.Review
	autoID  int64
}

var Reviews ReviewRepo //Khai báo biến toàn cục, global variable

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*model.Review)
	Reviews.InitData("sql:45312")
}

//Pointer receiver ~ method trong Java. Đối tượng chủ thể là *BookRepo
func (r *ReviewRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *ReviewRepo) CreateNewReview(review *model.Review) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	review.Id = nextID
	r.reviews[nextID] = review
	return nextID
}

func (r *ReviewRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewReview(&model.Review{
		BookId:  1,
		Comment: "good",
		Rating:  4,
	})

	r.CreateNewReview(&model.Review{
		BookId:  2,
		Comment: "bad",
		Rating:  2,
	})

	r.CreateNewReview(&model.Review{
		BookId:  2,
		Comment: "bad",
		Rating:  2,
	})

	r.CreateNewReview(&model.Review{
		BookId:  1,
		Comment: "good",
		Rating:  4,
	})
}

func (r *ReviewRepo) GetAllReviews() map[int64]*model.Review {
	return r.reviews
}

func (r *ReviewRepo) FindReviewById(Id int64) (*model.Review, error) {
	if review, ok := r.reviews[Id]; ok {
		return review, nil //tìm được
	} else {
		return nil, errors.New("review not found")
	}
}

func (r *ReviewRepo) DeleteReviewById(Id int64) error {
	if _, ok := r.reviews[Id]; ok {
		delete(r.reviews, Id)
		return nil
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepo) UpdateReview(review *model.Review) error {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return nil //tìm được
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepo) Upsert(review *model.Review) int64 {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return review.Id
	} else {
		return r.CreateNewReview(review)
	}
}

func (r *ReviewRepo) AverageRating() (result map[int64]float32) {
	sumRating := make(map[int64]int)
	numberRating := make(map[int64]int)
	result = make(map[int64]float32)

	for _, value := range r.reviews {
		numberRating[value.BookId]++
		sumRating[value.BookId] += value.Rating
	}
	for key := range numberRating {
		result[key] = float32(sumRating[key]) / float32(numberRating[key])
	}
	return result
}
