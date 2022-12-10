package controller

const (
	// API represents the group of API.
	API = "/api"
	// APIMeals represents the group of  meals API.
	APIMeals = API + "/Meals"
	// APIMealsID represents the API to get meals data using id.
	APIMealsID = APIMeals + "/:id"
	// APIFoods represents the group of food  API.
	APIFoods = API + "/food"
)

const (
	// APIUser represents the group of auth management API.
	APIUser = API + "/auth"
	// APIUserLoginStatus represents the API to get the status of logged in User.
	APIUserLoginStatus = APIUser + "/loginStatus"
	// APIUserLoginUser represents the API to get the logged in User.
	APIUserLoginUser = APIUser + "/loginUser"
	// APIUserLogin represents the API to login by session authentication.
	APIUserLogin = APIUser + "/login"
	// APIUserLogout represents the API to logout.
	APIUserLogout = APIUser + "/logout"
)
