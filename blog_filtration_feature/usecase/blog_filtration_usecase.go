package usecase

import "blog_filtration_feature/domain"

type BlogFiltrationUseCase struct {
	Repository 	  	domain.BlogFiltrationRepository	
}

func NewFiltrationUseCase(filtrationRepository domain.BlogFiltrationRepository) domain.BlogFiltrationUseCase {
	return &BlogFiltrationUseCase {
		Repository: filtrationRepository,
	}
}

func (fu *BlogFiltrationUseCase) FilterBlog(blogRequest *domain.BlogRequest) ([]*domain.BlogResponse, error) {
	filteredBlog, err := fu.Repository.FilterBlog(blogRequest)
	return filteredBlog, err
}