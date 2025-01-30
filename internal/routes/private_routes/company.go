package privateRoutes

import (
	"github.com/denyherianto/go-fiber-clean-boilerplate/database"
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/handlers"
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/repositories"
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/usecases/company"
	"github.com/gofiber/fiber/v2"
)

// CompanyRoutes func for describe group of private Companies routes.
func CompanyRoutes(r fiber.Router, db database.Database) {
	companyRepository := repositories.NewCompanyRepository(db)

	companyUseCase := company.NewUseCase(
		companyRepository,
	)

	companyHttpHandler := handlers.NewCompanyHandler(companyUseCase)

	companyRoute := r.Group("/companies")

	companyRoute.Get("/", companyHttpHandler.GetCompanies)
	companyRoute.Get("/:id", companyHttpHandler.GetCompany)
	companyRoute.Post("/", companyHttpHandler.CreateCompany)
	companyRoute.Put("/:id", companyHttpHandler.UpdateCompany)
	companyRoute.Delete("/:id", companyHttpHandler.DeleteCompany)
}
