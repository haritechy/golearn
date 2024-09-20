package controller

import (
	"employeeregister/database"
	"employeeregister/models"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type input struct {
	Age         int `json:"age"`
	VoterId     int `json:"voterid"`
	CandidateId int `json:"candidateid"`
}

func CandidatePost(c *gin.Context) {

	var Candidate models.Candidate

	var ValidateName = regexp.MustCompile(`^[a-zA-Z\s]+$`)

	if err := c.BindJSON(&Candidate); err != nil {

		logger.Error("invalid requste", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	StordeLower := strings.ToLower(Candidate.Name)
	Candidate.Name = StordeLower

	if !ValidateName.MatchString(Candidate.Name) {

		logger.Errorf("Please Avoid special character")
		c.JSON(http.StatusBadRequest, "Please Avoid special charcter")
		return
	}

	database.DB.Create(&Candidate)
	c.JSON(http.StatusOK, Candidate)
	logger.WithFields(logrus.Fields{
		"data": Candidate,
	}).Info("Candidate  Create  successfully")
}

func VotingController(c *gin.Context) {

	var input input

	if err := c.BindJSON(&input); err != nil {

		logger.Errorf("invalid input%v ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if input.Age < 18 {

		logger.Error("Your not not eligiple  for voting ")
		c.JSON(http.StatusForbidden, gin.H{"mesage": "Your not eligiple "})

		return

	}
	if input.VoterId < 10 {

		logger.Error("your voter id is wrong")
		c.JSON(http.StatusForbidden, gin.H{"message": "voter id is wronng pease enter crt voter id"})
		return
	}

	var candidte models.Candidate
	var Votes models.Votes
	if err := database.DB.Where("voter_id =?", &input.VoterId).First(&Votes).Error; err == nil {
		logger.Errorf("Alredy  vooted for this voter %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "this voter alredy voted"})
		return

	}

	if err := database.DB.First(&candidte, input.CandidateId).Error; err != nil {

		logger.Errorf("invalid candidate %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	candidte.VoteCount++
	database.DB.Save(&candidte)
	vote := models.Votes{VoterId: input.VoterId, CandidateId: input.CandidateId}
	database.DB.Create(&vote)
	c.JSON(http.StatusOK, gin.H{"Message": fmt.Sprintf("you have succefuly voted for  candidate  %v", candidte.Name)})

}
func VotingResult(c *gin.Context) {

	var winnerofElection []models.Candidate
	if err := database.DB.Order("vote_count DESC").Find(&winnerofElection).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find winner"})
		return
	}
	logger.WithFields(logrus.Fields{
		"data": winnerofElection[0],
	}).Info("Congratulation Party")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("election result winner is  %v  votes", winnerofElection[0]),

		"Places": winnerofElection,
	})

}

func DeleteCandidate(c *gin.Context) {

	var Candidate models.Candidate
	param := c.Param("id")

	if err := database.DB.First(&Candidate, param).Error; err != nil {

		logger.Errorf("id not found", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Delete(Candidate)
	c.JSON(http.StatusOK, gin.H{"Message": "delete candidate suceful"})
}

func TotalResult(c *gin.Context) {

	var TotalResult []models.Candidate

	if err := database.DB.Find(&TotalResult); err == nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, &TotalResult)
}

func SearchCandidate(c *gin.Context) {
	var candidates []models.Candidate
	searchWord := c.Query("word")

	if strings.ToUpper(searchWord) != "" {

		if err := database.DB.Where("name LIKE ?", "%"+searchWord+"%").Find(&candidates).Error; err != nil {

			logger.Errorf("Error while searching for candidates: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while searching for candidates", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, candidates)
	} else {

		logger.Errorf("Please search with a valid query format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid search query"})
	}
}
