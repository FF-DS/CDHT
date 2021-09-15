package controllers

import (
	. "monitoring-server/services"
	"monitoring-server/core"
    "github.com/gin-gonic/gin"
	"context"
    "log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
    // "time"
	"net/http"
	"strconv"
)


// type NodesController struct{}

type ConfigurationController struct{}
func (config ConfigurationController) Config(c *gin.Context){
    var success  = map[string]string{}
    success["message"] = "this is from configuration"
    c.JSON(200 , success)
}

/*
This function will return all the configuration profiles created and stored in the collection
  */
func (config ConfigurationController) GetConfigurationProfiles(c *gin.Context){
	configurationCollection := ConnectDB("configurations")

	cursor , err := configurationCollection.Find(context.TODO() , bson.M{} )

	if err != nil{
		GetError(err , c)
		return
	}

	defer cursor.Close(context.TODO())
    var configurationProfiles []core.ConfigurationProfile


    for cursor.Next(context.TODO()) {
		var configurationProfile core.ConfigurationProfile
		err := cursor.Decode(&configurationProfile)

		if err != nil {
			log.Fatal(err)
		}

		configurationProfiles = append(configurationProfiles, configurationProfile)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
 
    c.JSON(http.StatusOK, gin.H{"message": "All Configuration Profiles retrived successfully", "data": configurationProfiles} )
}

/* 
This function will return the currently used configuration profile
*/
func (config ConfigurationController) GetCurrentConfigurationProfile(c *gin.Context){
	configurationCollection := ConnectDB("configurations")

  	filter := bson.D{{"config_status", "ACTIVE"}}

  	var activeConfigurationProfile core.ConfigurationProfile

  	err := configurationCollection.FindOne(context.TODO() , filter).Decode(&activeConfigurationProfile)

	if err != nil{
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Active Configuration Profile retrived successfully", "data": activeConfigurationProfile} )
    
}

/* 
This function will set a configuration profile selected by the user and notifies the dispatch server
to propagate the changes accordingly
*/
func (config ConfigurationController) SetCurrentConfigurationProfile(c *gin.Context){
    configurationCollection := ConnectDB("configurations")

	opts := options.FindOneAndUpdate().SetUpsert(true)

	activeConfigId := c.Query("cuurent_active_config_id")
	newConfigId := c.Query("new_active_config_id")

	var activeConfigurationProfile core.ConfigurationProfile
	var activeConfigurationProfileToSet core.ConfigurationProfile

	updateCurrentToInactive := bson.D{{"$set", bson.D{{"config_status", "INACTIVE"}}}}
	updateNewToActive := bson.D{{"$set", bson.D{{"config_status", "ACTIVE"}}}}

	currentActiveFilter := bson.D{{"_id", activeConfigId}}
	newActiveFilter := bson.D{{"_id", newConfigId}}

	if err := configurationCollection.FindOneAndUpdate(
		context.TODO(),
		currentActiveFilter,
		updateCurrentToInactive,
		opts,
	).Decode(&activeConfigurationProfile); err == nil{

		if err := configurationCollection.FindOneAndUpdate(
			context.TODO(),
			newActiveFilter,
			updateNewToActive,
			opts,
		).Decode(&activeConfigurationProfileToSet); err == nil{

			c.JSON(http.StatusOK,
				 gin.H{"message": "Current active configuration profile changed successfuly" ,
				  "data" : gin.H{"old"  :activeConfigurationProfile , "new" : activeConfigurationProfileToSet}})

		}else{
			GetError(err, c)
		}

	}else{
		GetError(err, c)
	}
   
}


/* 
This function will set the nodespace balancing of the system and notify the command dispatch server to 
to propagate the changes accordingly
*/
func (config ConfigurationController) SetNodeSpaceBalancing(c *gin.Context){
    
	configurationCollection := ConnectDB("configurations")

	opts := options.FindOneAndUpdate().SetUpsert(true)

	activeConfigId , _ := primitive.ObjectIDFromHex(c.Query("cuurent_active_config_id"))
	newJumpSpace , _ := strconv.Atoi(c.Query("new_jump_space"))

	var activeConfigurationProfile core.ConfigurationProfile
	updateJumpSpace := bson.D{{"$set", bson.D{{"jump_space_balancing", newJumpSpace}}}}
	currentActiveFilter := bson.D{{"_id", activeConfigId}}

	if err := configurationCollection.FindOneAndUpdate(
		context.TODO(),
		currentActiveFilter,
		updateJumpSpace,
		opts,
	).Decode(&activeConfigurationProfile); err == nil{
		c.JSON(http.StatusOK,
			gin.H{"message": "Current active configuration profile's Jump S[ace Balancing] changed successfuly" ,
			 "data" : activeConfigurationProfile})

	}else{
		GetError(err, c)
	}

}

/* 
This function will store an newly created configuration profile by the user
*/
func (config ConfigurationController) AddConfigurationProfile(c *gin.Context){
    configurationCollection := ConnectDB("configurations")

	var configurationProfile core.ConfigurationProfile

	if err := c.ShouldBindJSON(&configurationProfile) ; err == nil{
	
        result, err := configurationCollection.InsertOne(context.TODO(), configurationProfile)
    
        if err != nil {
            GetError(err, c)
        }
    
        c.JSON(http.StatusOK, gin.H{"message": "Configuration Profile added successfully" , "data" : result})
    } else {
        c.JSON(401, gin.H{"error": err.Error()})
    }
}

/* 
This function will delete a configuration profile unless its a currently used one
if the currently used one is being requested to be deleted notify the user to set another
profile
*/
func (config ConfigurationController) DeleteConfigurationProfile(c *gin.Context){

	configurationCollection := ConnectDB("configurations")	

	configId := c.Query("config_id")

	
	opts := options.FindOneAndDelete().
		SetProjection(bson.D{})

	var deletedConfigProfile core.ConfigurationProfile
	err := configurationCollection.FindOneAndDelete(
		context.TODO(),
		bson.D{{"_id", configId}},
		opts,
	).Decode(&deletedConfigProfile)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			// notify the user the obj has not been founded
			return
		}
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Configuration Profile deleted successfully" , "data" : deletedConfigProfile})


    
}

/* 
This function will clear the configuration profiles database , thread lightly!!!!!!!!!!!!
*/
func (config ConfigurationController) ClearConfigurationProfilesCollection(c *gin.Context){
    
}
