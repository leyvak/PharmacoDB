package main

import "gopkg.in/gin-gonic/gin.v1"

// GetTissues handles GET requests for /tissues endpoint.
func GetTissues(c *gin.Context) {
	getDataTypes(c, "List of all tissues in pharmacodb",
		"select tissue_id, tissue_name from tissues;")
}

// GetTissueStats handles GET requests for /tissues/stats endpoint.
func GetTissueStats(c *gin.Context) {
	getDataTypeStats(c, "Number of tissues tested in each dataset", "select dataset_id, tissues from dataset_statistics;")
}

// GetTissueIDs handles GET requests for /tissues/ids endpoint.
func GetTissueIDs(c *gin.Context) {
	getDataTypeIDs(c, "List of all tissue IDs in pharmacodb", "select tissue_id from tissues;")
}

// GetTissueNames handles GET requests for /tissues/names endpoint.
func GetTissueNames(c *gin.Context) {
	getDataTypeNames(c, "List of all tissue names in pharmacodb", "select tissue_name from tissues;")
}
