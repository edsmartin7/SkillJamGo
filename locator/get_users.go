package locator 

/*
   return all the users near you
*/

import (
   db "skilljamgo/model"
)


//general search by proximity
func ReturnAllUsers(zip int) []string {
   //get closest zips from database
   zipList := db.GetClosestZipcodes()

   //get users from those zips in database
   var closestUsers []string
   for zip := range zipList {
      userList := db.getUsersByZip()
      closestUsers = append(closestUsers, userList)
   }

   return closestUsers
}

//game mode
func ReturnUnseenCloseUsers(zip int) {

}
