package privateRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inditekno/office-backend/database"
	"github.com/inditekno/office-backend/internal/handlers"
	"github.com/inditekno/office-backend/internal/repositories"
	"github.com/inditekno/office-backend/internal/usecases/company"
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
