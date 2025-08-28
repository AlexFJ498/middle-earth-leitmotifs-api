package deleting

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
)

// MovieService is the default implementation of the MovieService interface.
type MovieService struct {
	movieRepository domain.MovieRepository
}

// NewMovieService returns a new MovieService instance.
func NewMovieService(repo domain.MovieRepository) MovieService {
	return MovieService{
		movieRepository: repo,
	}
}

// DeleteMovie deletes a movie by its ID.
func (s *MovieService) DeleteMovie(ctx context.Context, id domain.MovieID) error {
	return s.movieRepository.Delete(ctx, id)
}

// GroupService is the default implementation of the GroupService interface.
type GroupService struct {
	groupRepository domain.GroupRepository
}

// NewGroupService returns a new GroupService instance.
func NewGroupService(repo domain.GroupRepository) GroupService {
	return GroupService{
		groupRepository: repo,
	}
}

// DeleteGroup deletes a group by its ID.
func (s *GroupService) DeleteGroup(ctx context.Context, id domain.GroupID) error {
	return s.groupRepository.Delete(ctx, id)
}

// CategoryService is the default implementation of the CategoryService interface.
type CategoryService struct {
	categoryRepository domain.CategoryRepository
}

// NewCategoryService returns a new CategoryService instance.
func NewCategoryService(repo domain.CategoryRepository) CategoryService {
	return CategoryService{
		categoryRepository: repo,
	}
}

// DeleteCategory deletes a category by its ID.
func (s *CategoryService) DeleteCategory(ctx context.Context, id domain.CategoryID) error {
	return s.categoryRepository.Delete(ctx, id)
}
