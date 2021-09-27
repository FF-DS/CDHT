package controllers

import (
	"context"
	"monitoring-server/core"
	"monitoring-server/services"
	"monitoring-server/util"
	"strconv"

	"github.com/gin-gonic/gin"

	// "log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "time"
	"net/http"
	// "strconv"
)

/*
This function will return all the configuration profiles created and stored in the collection
*/
func (config ConfigurationController) GetConfigurationProfiles(c *gin.Context){

	configurationCollection := services.ConnectDB(CONFIGURATION_COLLECTION_NAME)

	type RequestBody struct{
		Limit int64 `bson:"limit" json:"limit"`
	}
	
	var request RequestBody

	findOptions := options.Find()

	if err := c.ShouldBindJSON(&request) ; err == nil{
		findOptions.SetLimit(request.Limit)
		filter := bson.M{}

		if cursor , err := configurationCollection.Find(context.TODO() , filter , findOptions ) ; err == nil{

			defer cursor.Close(context.TODO())
			var configurationProfiles []core.ConfigurationProfile
		
			for cursor.Next(context.TODO()) {
				var configurationProfile core.ConfigurationProfile
				if err := cursor.Decode(&configurationProfile) ; err == nil {
					configurationProfiles = append(configurationProfiles, configurationProfile)
				}else{
					message := "an error occured while trying to decode configuration profile"
					util.GetError(err , message , c)
				}
			}
			
			c.JSON(http.StatusOK, gin.H{"message": "All Configuration Profiles retrived successfully", "data": configurationProfiles} )
		}else {
			message := "an error occured while trying to fetch configuration profiles"
			util.GetError(err , message ,  c)
		}
	
	}else {
		message := "an error occured while trying to bind request object"
        util.GetError(err , message , c)
	}
}

/* 
This function will return the currently used configuration profile
*/
func (config ConfigurationController) GetCurrentConfigurationProfile(c *gin.Context){
	configurationCollection := services.ConnectDB(CONFIGURATION_COLLECTION_NAME)

  	filter := bson.D{primitive.E{Key : "config_status", Value: "ACTIVE"}}

  	var activeConfigurationProfile core.ConfigurationProfile

	if err := configurationCollection.FindOne(context.TODO() , filter).Decode(&activeConfigurationProfile) ; err == nil{
		c.JSON(http.StatusOK, gin.H{"message": "Active Configuration Profile retrived successfully", "data": activeConfigurationProfile} )
	}else {
		message := "an error occured while trying to retrive the active profile"
		util.GetError(err , message , c)
	}
}

/* 
This function will set a configuration profile selected by the user and notifies the dispatch server
to propagate the changes accordingly
*/
func (config ConfigurationController) SetCurrentConfigurationProfile(c *gin.Context){

    configurationCollection := services.ConnectDB(CONFIGURATION_COLLECTION_NAME)

	type RequestBody struct{
		NewActiveConfigurationId primitive.ObjectID `json:"new_active_configuration_id" bson:"new_active_configuration_id"`
	}

	var request RequestBody

	findOptions := options.FindOneAndUpdate().SetUpsert(true)

	var activeConfigurationProfile core.ConfigurationProfile
	var activeConfigurationProfileToSet core.ConfigurationProfile

	if err := c.ShouldBindJSON(&request) ; err == nil {
		filter := bson.D{primitive.E{Key : "_id" , Value: request.NewActiveConfigurationId}}

		if err := configurationCollection.FindOneAndUpdate(
			context.TODO(),
			bson.D{primitive.E{Key:"config_status" , Value:"ACTIVE"}},
			bson.D{primitive.E{Key:"$set",Value: bson.D{primitive.E{Key:"config_status", Value: "INACTIVE"}}}} ,
			findOptions,
		).Decode(&activeConfigurationProfile); err == nil{
			if err := configurationCollection.FindOneAndUpdate(
				context.TODO(),
				filter,
				bson.D{primitive.E{Key:"$set",Value: bson.D{primitive.E{Key:"config_status", Value: "ACTIVE"}}}} ,
				findOptions,
			).Decode(&activeConfigurationProfileToSet); err == nil{
				c.JSON(http.StatusOK,
					gin.H{"message": "Active configuration profile changed successfuly" ,
					 "data" : gin.H{"old"  :activeConfigurationProfile , "new" : activeConfigurationProfileToSet}})
			}else{
				message := "an error occured while trying to activate new configuration profile"
				util.GetError(err , message , c)
			}	

		}else{
			message := "an error occured while trying to deactivate the currently active configuration profile"
			util.GetError(err , message , c)
		}

	}else{
		message := "an error occured while trying to bind request object"
        util.GetError(err , message , c)
	}
   
}


/* 
This function will set the nodespace balancing of the system and notify the command dispatch server to 
to propagate the changes accordingly
*/
func (config ConfigurationController) SetNodeSpaceBalancing(c *gin.Context){
    
	configurationCollection := services.ConnectDB(CONFIGURATION_COLLECTION_NAME)

	type RequestBody struct{
		NewJumpSpace int `json:"new_jump_space" bson:"new_jump_space"`
	}

	var request RequestBody

	var activeConfigurationProfile core.ConfigurationProfile

	findOptions := options.FindOneAndUpdate().SetUpsert(true)

	if err := c.ShouldBindJSON(&request) ; err == nil{
		if err := configurationCollection.FindOneAndUpdate(
			context.TODO(),
			bson.D{primitive.E{Key : "config_status" ,Value: "ACTIVE"}},
			bson.D{primitive.E{Key : "$set" ,Value: bson.D{primitive.E{Key:"jump_spacing" , Value :request.NewJumpSpace}}}},
			findOptions,
		).Decode(&activeConfigurationProfile); err == nil{
			c.JSON(http.StatusOK,
				gin.H{"message": "Current active configuration profile's Jump Space Balancing changed successfuly" ,
				 "data" : activeConfigurationProfile})
		}else{
			message := "an error occured while trying to change the jump space of the current configuration profile"
			util.GetError(err, message ,  c)
		}
	}else {
		message := "an error occured while trying to bind request object "
        util.GetError(err , message , c)
	}
}

func (config ConfigurationController) SetReplicationLevel(c *gin.Context){
    
	configurationCollection := services.ConnectDB(CONFIGURATION_COLLECTION_NAME)

	type RequestBody struct{
		NewReplicationLevel int `json:"new_replication_level" bson:"new_replication_level"`
	}

	var request RequestBody

	var activeConfigurationProfile core.ConfigurationProfile

	findOptions := options.FindOneAndUpdate().SetUpsert(true)

	if err := c.ShouldBindJSON(&request) ; err == nil{
		if err := configurationCollection.FindOneAndUpdate(
			context.TODO(),
			bson.D{primitive.E{Key : "config_status" ,Value: "ACTIVE"}},
			bson.D{primitive.E{Key : "$set" ,Value: bson.D{primitive.E{Key:"replication_count" , Value :request.NewReplicationLevel}}}},
			findOptions,
		).Decode(&activeConfigurationProfile); err == nil{
			c.JSON(http.StatusOK,
				gin.H{"message": "Current active configuration profile's replication level changed successfuly" ,
				 "data" : activeConfigurationProfile})
		}else{
			message := "an error occured while trying to change the replication level of the current configuration profile"
			util.GetError(err, message ,  c)
		}
	}else {
		message := "an error occured while trying to bind request object "
        util.GetError(err , message , c)
	}
}

func (config ConfigurationController) GetConfigurationReportEntries(c *gin.Context){
    /* 
    this function will return everything as it is called with out any filters

    but appropriate filter must be applied as all log types are not relevant for the front end
    */
    logCollection := services.ConnectDB(LOG_COLLECTION_NAME)

	type RequestBody struct{
		Limit string `bson:"limit" json:"limit"`
	}
	
	var request RequestBody

	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key:"created_date",Value:  -1}})

	if err := c.ShouldBindJSON(&request) ; err == nil{
		limit , _ := strconv.ParseInt(request.Limit, 10, 64)
		findOptions.SetLimit(limit)

		
		if cursor , err := logCollection.Find(context.TODO() , bson.M{"type" : "TYPE_CONFIGURATION_SERVICE"} , findOptions) ; err == nil{

			defer cursor.Close(context.TODO())
			var logs []core.LogEntry
		
		
			for cursor.Next(context.TODO()) {
				var logEntry core.LogEntry
				if err := cursor.Decode(&logEntry) ; err == nil{

					logs = append(logs, logEntry)
				}else{
					message := "an error occured while trying to decode report entry"
					util.GetError(err , message , c)
				}
			}

			c.JSON(http.StatusOK, gin.H{"message": "Reports retrived succesfully", "data": logs} )

		}else{
			message := "an error occured while trying to fetch report entries"
			util.GetError(err , message , c)
		}
	
	}else{
		message := "an error occured while trying to bind request object"
        util.GetError(err , message , c)
	}
}

/* 
This function will store an newly created configuration profile by the user
*/
func (config ConfigurationController) AddConfigurationProfile(c *gin.Context){
    configurationCollection := services.ConnectDB(CONFIGURATION_COLLECTION_NAME)

	var configurationProfile core.ConfigurationProfile

	if err := c.ShouldBindJSON(&configurationProfile) ; err == nil{
	
        if result, err := configurationCollection.InsertOne(context.TODO(), configurationProfile) ; err == nil{
			c.JSON(http.StatusOK, gin.H{"message": "Configuration Profile added successfully" , "data" : result})
		}else{
            
			message := "an error occured while trying to insert the new configuration profile" 
			util.GetError(err, message , c)
        }
    } else {
        message := "an error occured while trying to bind request object "
        util.GetError(err , message , c)
    }
}

/* 
This function will delete a configuration profile unless its a currently used one
if the currently used one is being requested to be deleted notify the user to set another
profile
*/
func (config ConfigurationController) DeleteConfigurationProfile(c *gin.Context){

	configurationCollection := services.ConnectDB(CONFIGURATION_COLLECTION_NAME)	

	type RequestBody struct{
		ConfigurationId primitive.ObjectID `json:"configuration_id" bson:"configuration_id"`
	}

	var request RequestBody
	var deletedConfigProfile core.ConfigurationProfile

	deleteOptions := options.FindOneAndDelete().
		SetProjection(bson.D{})

	if err := c.ShouldBindJSON(&request) ; err == nil {

		if err := configurationCollection.FindOneAndDelete(
			context.TODO(),
			bson.D{primitive.E{Key :"_id" , Value : request.ConfigurationId}}, 
			deleteOptions,
		).Decode(&deletedConfigProfile) ; err == nil {
			c.JSON(http.StatusOK, gin.H{"message": "Configuration Profile deleted successfully" , "data" : deletedConfigProfile})
		}else{
			message := "an error occured while trying to delete the specified configuration profile"
        util.GetError(err , message , c)
		}
	}else{
		message := "an error occured while trying to bind request object "
        util.GetError(err , message , c)
	}    
}

/* 
This function will clear the configuration profiles database , thread lightly!!!!!!!!!!!!
*/
func (config ConfigurationController) ClearConfigurationProfilesCollection(c *gin.Context){
    
}
