package routes

import (
	"employeeregister/controller"

	"github.com/gin-gonic/gin"
)

func VoterRoutes(r *gin.Engine) {
	VoterRoutes := r.Group("/vote")
	{
		VoterRoutes.POST("/create", controller.CandidatePost)
		VoterRoutes.POST("/voting", controller.VotingController)
		VoterRoutes.GET("/result", controller.VotingResult)
		VoterRoutes.GET("/results", controller.TotalResult)
		VoterRoutes.DELETE("/delete/:id", controller.DeleteCandidate)
		VoterRoutes.GET("/search", controller.SearchCandidate)

	}
}
