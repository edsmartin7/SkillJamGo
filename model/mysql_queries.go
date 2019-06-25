package model

const (
   CheckCredentials = `SELECT username, password
                       FROM logins
                       WHERE username=?
                       AND password=?`

   GetClosestZips = `SELECT zipcodes 
                     FROM zipdb
                     WHERE zip=?
                     LIMIT 5`
)
