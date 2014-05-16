$stars = angular.module('stars',['ngRoute']);

$stars.config(function($routeProvider, $locationProvider){
	$locationProvider.html5Mode(true);
	$routeProvider
		.when('/dashboard',{templateUrl:'/dashboard/page'})
		.when('/employees',{templateUrl:'/employees/page',controller:'EmployeesCtrl'})
		.when('/employee/new',{
			templateUrl:'/employee/page',
			controller:'EmployeeEditCtrl',
			resolve:{
				login: function(){
					return "@";
				}
			}})
		.when('/employee/:login',{
			templateUrl:'/employee/page',
			controller:'EmployeeEditCtrl', 
			resolve:{
				login: function($route){
					return $route.current.params.login
				}
			}
		})
		.when('/employee/profile',{templateUrl:'/employee/profile',controller:'EmployeeProfileCtrl'})
		.when('/assays',{templateUrl:'/assays/page',controller:'AssaysCtrl'})
		
		.otherwise({redirectTo:'/'});
})