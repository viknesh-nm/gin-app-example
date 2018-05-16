package main

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// JobDetails holds the job fields
type JobDetails struct {
	ID, JobStatusID                                   int
	Title, Exp, ShortDescription, Description, Desire string
}

var jobList = []JobDetails{}

// showNewJobPage execute the add new job page (GET)
func showNewJobPage(c *gin.Context) {
	render(c, gin.H{
		"Title": "Add New Job", "UserName": userName,
	}, "add-job.html")
}

// addNewJob posts the job details
// Post Method
func addNewJob(c *gin.Context) {
	if _, err := insertJobDetails(c.PostForm("job_title"), c.PostForm("job_experience"), c.PostForm("job_required_skills"), c.PostForm("job_description"), c.PostForm("job_desire")); err == nil {
		c.Redirect(307, "/admin/job_openings")
	} else {
		c.AbortWithStatus(400)
	}
}

// insertJobDetails fetches the values from the entered form and stores in a DB
// Returns the Job detailed values
func insertJobDetails(jobTitle, jobExperience, jobRequiredSkills, jobDescription, jobDesire string) (*JobDetails, error) {
	_, err := db.Exec("INSERT INTO jobOpenings(title, exp, shortDescription, description, jobStatusID, desire, createdAt, updatedAt) VALUES(?,?,?,?,?,?,?,?)", jobTitle, jobExperience, jobRequiredSkills, jobDescription, 1, jobDesire, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	jobDetails := JobDetails{ID: len(jobList) + 1, Title: jobTitle, Exp: jobExperience, ShortDescription: jobRequiredSkills, Description: jobDescription, JobStatusID: 1, Desire: jobDesire}
	return &jobDetails, nil
}

// listAllJobs lists all the jobs from the DB
func listAllJobs() []JobDetails {
	var (
		lists = []JobDetails{}
		list  JobDetails
		id    int

		jobTitle, jobExperience, jobRequiredSkills, jobDescription, jobDesire string
	)
	rows, err := db.Query("select openingsId, title, exp, shortDescription, description, desire from jobOpenings")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &jobTitle, &jobExperience, &jobRequiredSkills, &jobDescription, &jobDesire)
		if err != nil {
			log.Fatal(err)
		}

		list = JobDetails{ID: id, Title: jobTitle, Exp: jobExperience, ShortDescription: jobRequiredSkills, Description: jobDescription, JobStatusID: 1, Desire: jobDesire}
		lists = append(lists, list)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lists
}

// showEditPage shows the edit job page fetched from the particular ID
// GET method
func showEditPage(c *gin.Context) {
	if openingsID, err := strconv.Atoi(c.Query("openingsID")); err == nil {
		if jobDet, err := getOpeningJobByID(openingsID); err == nil {
			render(c, gin.H{
				"title":   jobDet.Title,
				"payload": jobDet}, "add-job.html")
		} else {
			c.AbortWithError(404, err)
		}
	} else {
		c.AbortWithStatus(404)
	}
}

// getOpeningJobByID fetches the job details based on the ID supplied
func getOpeningJobByID(id int) (*JobDetails, error) {
	var jobTitle, jobExperience, jobRequiredSkills, jobDescription, jobDesire string
	rows, err := db.Query("select openingsId, title, exp, shortDescription, description, desire from jobOpenings where openingsId = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &jobTitle, &jobExperience, &jobRequiredSkills, &jobDescription, &jobDesire)
		if err != nil {
			log.Fatal(err)
		}
		return &JobDetails{ID: id, Title: jobTitle, Exp: jobExperience, ShortDescription: jobRequiredSkills, Description: jobDescription, JobStatusID: 1, Desire: jobDesire}, nil
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return nil, errors.New("Openings not found")
}

// editPage performs the updation of the edited fields
// POST Method
func editPage(c *gin.Context) {
	if openingsID, err := strconv.Atoi(c.Query("openingsID")); err == nil {
		if _, err := updatePage(openingsID, c.PostForm("job_title"), c.PostForm("job_experience"), c.PostForm("job_required_skills"), c.PostForm("job_description"), c.PostForm("job_desire")); err == nil {
			c.Redirect(301, "/admin/job_openings")
		} else {
			c.AbortWithStatus(400)
		}
	} else {
		c.AbortWithStatus(404)
	}
}

// updatePage updates the given values to the DB and returns it back
func updatePage(openingsID int, jobTitle, jobExperience, jobRequiredSkills, jobDescription, jobDesire string) (*JobDetails, error) {
	_, err := db.Exec("UPDATE jobOpenings SET title = ?, exp = ?, shortDescription = ?, description = ?, desire = ?, updatedAt = ? WHERE openingsId = ?",
		jobTitle, jobExperience, jobRequiredSkills, jobDescription, jobDesire, time.Now(), openingsID)
	if err != nil {
		return nil, err
	}
	return &JobDetails{ID: openingsID, Title: jobTitle, Exp: jobExperience, ShortDescription: jobRequiredSkills, Description: jobDescription, JobStatusID: 1, Desire: jobDesire}, nil
}

// deleteJobList deletes the job details with the particular selected id
func deleteJobList(c *gin.Context) {
	if jobID, err := strconv.Atoi(c.Param("id")); err == nil {
		_, err := db.Exec("DELETE FROM jobOpenings where openingsId= ?", jobID)
		if err == nil {
			c.Redirect(301, "/admin/job_openings")
		}
	} else {
		c.AbortWithStatus(404)
	}
}
