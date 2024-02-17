# Go Server
A simple test server in go language

## Request :

Create an HTTP Server with some endpoint :
- `/event` add information into your database of choice. The information added can be of type `impression` `click` `call`
- `/read` read information from your database, you must compute another infos such as `daily_impression`, `hourly_impression`, `click_rate`, `impression_rate`
