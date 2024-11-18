package tourOperatorUseCase

import (
	"github.com/Bigthugboy/TourWithUs/internals/domain/model/tourModel"
	model "github.com/Bigthugboy/TourWithUs/internals/domain/model/tourOpModel"
)

type TourOperatorUseCase interface {
	RegisterTourOperator(operator model.TourOperator) (*model.CreateTourOperatorResponse, error)
	Login(req model.LoginRequest) (*model.LoginRes, error)

	CreateTour(tour tourModel.TourDto) (*tourModel.CreateTourResponse, error)           // Create a new tour
	UpdateTour(tourID string, tour tourModel.UpdateTourDto) (*tourModel.TourDto, error) // Update tour details
	DeleteTour(tourID string) (string, error)                                           // Delete a specific tour
	ViewTourDetails(tourID string) (*tourModel.TourDto, error)                          // View details of a specific tour
	ListTours(operatorID string) ([]tourModel.TourDto, error)                           // List all tours by a tour operator

	//ViewBookings(tourID string) ([]model.Booking, error) // View all bookings for a specific tour
	ConfirmBooking(bookingID string) error // Confirm a customer's booking
	CancelBooking(bookingID string) error  // Cancel a customer's booking

	//GenerateTourReport(tourID string) (*model.TourReport, error) // Generate a report for a specific tour
	ManageAvailability(tourID string, availability bool) error // Enable or disable availability of a tour

	UpdateProfile(operator model.TourOperator) (*model.CreateTourOperatorResponse, error) // Update the tour operator's profile
	ChangePassword(req model.ChangePasswordRequest) error                                 // Change the tour operator's password
	Logout(operatorID string) error                                                       // Log out the tour operator
}
